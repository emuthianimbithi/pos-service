package repository

import (
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

// CreateAlert creates a new alert
func (r *NotificationRepository) CreateAlert(alert *models.Alert) error {
	return r.db.Create(alert).Error
}

// GetActiveAlertsForUser fetches unread, non-expired alerts for a specific user based on ID and Role
func (r *NotificationRepository) GetActiveAlertsForUser(userID uuid.UUID, role string, businessID uuid.UUID, branchID *uuid.UUID) ([]models.Alert, error) {
	var alerts []models.Alert

	query := r.db.Where("business_id = ?", businessID).
		Where("is_read = ?", false).
		Where("expires_at > ?", time.Now())

	// Filter by Branch (Global alerts OR specific branch alerts)
	if branchID != nil {
		query = query.Where("branch_id IS NULL OR branch_id = ?", branchID)
	} else {
		query = query.Where("branch_id IS NULL")
	}

	// Filter by Target (Specific User OR Role OR Public)
	// Logic: (TargetUser == UserID) OR (TargetRole == Role) OR (TargetUser IS NULL AND TargetRole IS NULL)
	query = query.Where(
		r.db.Where("target_user_id = ?", userID).
			Or("target_role = ?", role).
			Or("target_user_id IS NULL AND target_role = ''"),
	)

	err := query.Order("created_at desc").Find(&alerts).Error
	return alerts, err
}

// MarkAlertAsRead marks an alert as read
func (r *NotificationRepository) MarkAlertAsRead(alertID uuid.UUID) error {
	now := time.Now()
	return r.db.Model(&models.Alert{}).Where("id = ?", alertID).Updates(map[string]interface{}{
		"is_read": true,
		"read_at": &now,
	}).Error
}

// CreateNotification creates a transient notification
func (r *NotificationRepository) CreateNotification(notif *models.Notification) error {
	return r.db.Create(notif).Error
}

// GetRecentNotifications fetches recent notifications for a user
func (r *NotificationRepository) GetRecentNotifications(userID uuid.UUID, limit int) ([]models.Notification, error) {
	var notifs []models.Notification
	err := r.db.Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Find(&notifs).Error
	return notifs, err
}
