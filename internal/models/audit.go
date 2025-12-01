package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuditLog tracks all API requests and actions performed in the system
type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	// Request Information
	Method     string `gorm:"type:varchar(10);not null" json:"method"`
	Path       string `gorm:"type:varchar(255);not null" json:"path"`
	StatusCode int    `gorm:"not null" json:"status_code"`
	IPAddress  string `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent  string `gorm:"type:text" json:"user_agent"`

	// User Context
	UserID     *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"`
	BusinessID *uuid.UUID `gorm:"type:uuid;index" json:"business_id,omitempty"`
	BranchID   *uuid.UUID `gorm:"type:uuid;index" json:"branch_id,omitempty"`

	// Action Details
	Action      string `gorm:"type:varchar(100)" json:"action"`
	Resource    string `gorm:"type:varchar(100)" json:"resource"`
	ResourceID  string `gorm:"type:varchar(100)" json:"resource_id,omitempty"`
	Description string `gorm:"type:text" json:"description,omitempty"`

	// Request/Response Data (optional, for sensitive operations)
	RequestBody  string `gorm:"type:text" json:"request_body,omitempty"`
	ResponseBody string `gorm:"type:text" json:"response_body,omitempty"`

	// Performance
	Duration int64 `json:"duration_ms"` // in milliseconds

	// Additional metadata
	Metadata string `gorm:"type:jsonb" json:"metadata,omitempty"`
}

// BeforeCreate hook to generate UUID
func (a *AuditLog) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}

// TableName specifies the table name for AuditLog
func (*AuditLog) TableName() string {
	return "audit_logs"
}
