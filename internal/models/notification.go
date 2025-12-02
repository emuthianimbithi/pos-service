package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AlertPriority defines the urgency of an alert
type AlertPriority string

const (
	PriorityLow      AlertPriority = "low"
	PriorityMedium   AlertPriority = "medium"
	PriorityHigh     AlertPriority = "high"
	PriorityCritical AlertPriority = "critical"
)

// Alert represents a persistent, actionable notification
type Alert struct {
	BaseModel
	TenantModel
	BranchID *uuid.UUID `json:"branch_id,omitempty" gorm:"type:uuid"` // Optional: specific to a branch

	// Content
	Title      string        `json:"title" gorm:"not null"`
	Message    string        `json:"message" gorm:"not null"`
	Type       string        `json:"type" gorm:"not null"` // e.g., "stock_low", "security", "system"
	Priority   AlertPriority `json:"priority" gorm:"default:'medium'"`
	ActionLink string        `json:"action_link,omitempty"` // URL to resolve the issue

	// Targeting
	TargetUserID *uuid.UUID `json:"target_user_id,omitempty" gorm:"type:uuid;index"` // Specific user
	TargetRole   string     `json:"target_role,omitempty" gorm:"index"`              // e.g., "manager", "admin"

	// Status
	IsRead    bool       `json:"is_read" gorm:"default:false"`
	ReadAt    *time.Time `json:"read_at,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// BeforeCreate hook to set default expiration if not provided
func (a *Alert) BeforeCreate(tx *gorm.DB) error {
	if a.ExpiresAt == nil {
		// Default expiration: 30 days
		expiry := time.Now().AddDate(0, 0, 30)
		a.ExpiresAt = &expiry
	}
	return a.BaseModel.BeforeCreate(tx)
}

// Notification represents a transient, real-time update (toast)
// These might not always be stored in DB, but good to have a record
type Notification struct {
	BaseModel
	TenantModel

	UserID  uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Title   string    `json:"title" gorm:"not null"`
	Message string    `json:"message" gorm:"not null"`
	Type    string    `json:"type" gorm:"default:'info'"` // success, error, info, warning
	IsRead  bool      `json:"is_read" gorm:"default:false"`
}
