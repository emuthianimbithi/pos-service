package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	maxLogSize = 1024 * 1024 // 1MB limit for logging bodies
)

var sensitiveFields = []string{
	"password",
	"password_confirmation",
	"token",
	"access_token",
	"refresh_token",
	"secret",
	"api_key",
	"credit_card",
	"cvv",
}

// bodyLogWriter is a wrapper to capture response body
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	// Only buffer if we haven't exceeded the limit
	if w.body.Len() < maxLogSize {
		remaining := maxLogSize - w.body.Len()
		if len(b) > remaining {
			w.body.Write(b[:remaining])
			w.body.WriteString("... (truncated)")
		} else {
			w.body.Write(b)
		}
	}
	return w.ResponseWriter.Write(b)
}

// scrubData removes sensitive fields from JSON data
func scrubData(data []byte) string {
	if len(data) == 0 {
		return ""
	}

	// If data is too large, don't try to parse it, just return truncated string
	if len(data) > maxLogSize {
		return string(data[:maxLogSize]) + "... (truncated)"
	}

	// Try to unmarshal into a map
	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		// Not a JSON object, maybe an array or plain string
		// For now, just return as string
		return string(data)
	}

	// Recursively scrub sensitive fields
	scrubMap(obj)

	// Marshal back to string
	scrubbed, err := json.Marshal(obj)
	if err != nil {
		return string(data)
	}

	return string(scrubbed)
}

func scrubMap(m map[string]interface{}) {
	for k, v := range m {
		// Check if key is sensitive
		isSensitive := false
		for _, field := range sensitiveFields {
			if strings.Contains(strings.ToLower(k), field) {
				isSensitive = true
				break
			}
		}

		if isSensitive {
			m[k] = "[REDACTED]"
			continue
		}

		// Recurse if value is a map
		if nestedMap, ok := v.(map[string]interface{}); ok {
			scrubMap(nestedMap)
		} else if nestedArray, ok := v.([]interface{}); ok {
			// Check array items if they are maps
			for _, item := range nestedArray {
				if itemMap, ok := item.(map[string]interface{}); ok {
					scrubMap(itemMap)
				}
			}
		}
	}
}

// AuditMiddleware logs all requests to the audit_logs table
func AuditMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 1. Capture Request Body
		var requestBody []byte
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			if len(bodyBytes) <= maxLogSize {
				requestBody = bodyBytes
			} else {
				requestBody = []byte("Request body too large to log")
			}
		}

		// 2. Capture Response Body
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		duration := time.Since(start).Milliseconds()

		// Scrub sensitive data
		scrubbedRequest := scrubData(requestBody)
		scrubbedResponse := scrubData(blw.body.Bytes())

		if scrubbedRequest == "" {
			scrubbedRequest = "{}"
		}
		if scrubbedResponse == "" {
			scrubbedResponse = "{}"
		}

		// Ensure metadata is always valid JSON
		metadata := datatypes.JSON([]byte(`{}`))

		auditLog := models.AuditLog{
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
			StatusCode:   c.Writer.Status(),
			IPAddress:    c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
			Duration:     duration,
			RequestBody:  scrubbedRequest,
			ResponseBody: scrubbedResponse,
			Metadata:     string(metadata), // <-- FIXED
		}

		// Extract user context
		if userID, exists := c.Get("user_id"); exists {
			if uid, ok := userID.(uuid.UUID); ok {
				auditLog.UserID = &uid
			}
		}

		if businessID, exists := c.Get("business_id"); exists {
			if bid, ok := businessID.(uuid.UUID); ok {
				auditLog.BusinessID = &bid
			}
		}

		if branchID, exists := c.Get("branch_id"); exists {
			if bid, ok := branchID.(uuid.UUID); ok {
				auditLog.BranchID = &bid
			}
		}

		// Extract action details
		if action, exists := c.Get("audit_action"); exists {
			if a, ok := action.(string); ok {
				auditLog.Action = a
			}
		}

		if resource, exists := c.Get("audit_resource"); exists {
			if r, ok := resource.(string); ok {
				auditLog.Resource = r
			}
		}

		if resourceID, exists := c.Get("audit_resource_id"); exists {
			if rid, ok := resourceID.(string); ok {
				auditLog.ResourceID = rid
			}
		}

		// Save immediately (Cloud Run safe)
		if err := db.Session(&gorm.Session{NewDB: true}).Create(&auditLog).Error; err != nil {
			// log.Printf("Failed to create audit log: %v", err)
		}
	}
}

// SetAuditAction is a helper to set audit action in context
func SetAuditAction(c *gin.Context, action, resource, resourceID string) {
	c.Set("audit_action", action)
	c.Set("audit_resource", resource)
	if resourceID != "" {
		c.Set("audit_resource_id", resourceID)
	}
}
