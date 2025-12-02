package middleware

import (
	"net/http"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TenantMiddleware ensures business and branch context for multi-tenant operations
func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Business ID can come from JWT claims (set by AuthMiddleware)
		// or from request headers for specific operations
		_, exists := c.Get("business_id")
		if !exists {
			// Try to get from header
			if businessIDHeader := c.GetHeader("X-Business-ID"); businessIDHeader != "" {
				if bid, err := uuid.Parse(businessIDHeader); err == nil {
					c.Set("business_id", bid)
				}
			}
		}

		// For branch-scoped operations
		_, branchExists := c.Get("branch_id")
		if !branchExists {
			// Try to get from header
			if branchIDHeader := c.GetHeader("X-Branch-ID"); branchIDHeader != "" {
				if bid, err := uuid.Parse(branchIDHeader); err == nil {
					c.Set("branch_id", bid)
				}
			}
		}

		c.Next()
	}
}

// RequireBusiness ensures business context exists
func RequireBusiness() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("business_id")
		if !exists {
			response.Error(c, http.StatusBadRequest, "Business context required")
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireBranch ensures branch context exists
func RequireBranch() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("branch_id")
		role, _ := c.Get("role")
		if !exists && role != models.RoleAdmin {
			response.Error(c, http.StatusBadRequest, "Branch context required")
			c.Abort()
			return
		}
		c.Next()
	}
}
