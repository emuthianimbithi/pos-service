package services

import (
	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

type AlertService struct {
	repo *repository.NotificationRepository
}

func NewAlertService(repo *repository.NotificationRepository) *AlertService {
	return &AlertService{repo: repo}
}

// CreateUserAlert sends an alert to a specific user
func (s *AlertService) CreateUserAlert(businessID uuid.UUID, userID uuid.UUID, title, message, priority string) error {
	alert := &models.Alert{
		TenantModel:  models.TenantModel{BusinessID: businessID},
		Title:        title,
		Message:      message,
		Type:         "user_alert",
		Priority:     models.AlertPriority(priority),
		TargetUserID: &userID,
	}
	return s.repo.CreateAlert(alert)
}

// CreateRoleAlert sends an alert to all users with a specific role (optionally scoped to a branch)
func (s *AlertService) CreateRoleAlert(businessID uuid.UUID, branchID *uuid.UUID, role, title, message, priority string) error {
	alert := &models.Alert{
		TenantModel: models.TenantModel{BusinessID: businessID},
		BranchID:    branchID,
		Title:       title,
		Message:     message,
		Type:        "role_alert",
		Priority:    models.AlertPriority(priority),
		TargetRole:  role,
	}
	return s.repo.CreateAlert(alert)
}

// GetMyAlerts retrieves active alerts for the current user context
func (s *AlertService) GetMyAlerts(userID uuid.UUID, role string, businessID uuid.UUID, branchID *uuid.UUID) ([]models.Alert, error) {
	return s.repo.GetActiveAlertsForUser(userID, role, businessID, branchID)
}

// MarkAsRead marks an alert as read
func (s *AlertService) MarkAsRead(alertID uuid.UUID) error {
	return s.repo.MarkAlertAsRead(alertID)
}

// GetMyNotifications retrieves recent notifications for the user
func (s *AlertService) GetMyNotifications(userID uuid.UUID) ([]models.Notification, error) {
	return s.repo.GetRecentNotifications(userID, 20) // Limit to last 20
}
