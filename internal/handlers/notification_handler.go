package handlers

import (
	"net/http"

	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NotificationHandler struct {
	service *services.AlertService
}

func NewNotificationHandler(service *services.AlertService) *NotificationHandler {
	return &NotificationHandler{
		service: service,
	}
}

// GetAlerts returns active alerts for the logged-in user
func (h *NotificationHandler) GetAlerts(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}
	role, err := utils.GetUserRole(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Role required")
		return
	}

	// Branch is optional (admin might not have one)
	var branchIDPtr *uuid.UUID
	branchID, err := utils.GetBranchID(c)
	if err == nil {
		branchIDPtr = &branchID
	}

	alerts, err := h.service.GetMyAlerts(userID, role, businessID, branchIDPtr)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch alerts")
		return
	}

	response.Success(c, http.StatusOK, "Alerts retrieved", alerts)
}

// MarkRead marks an alert as read
func (h *NotificationHandler) MarkRead(c *gin.Context) {
	id, err := utils.ParseUUIDParam(c, "id")
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.MarkAsRead(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update alert")
		return
	}

	response.Success(c, http.StatusOK, "Alert marked as read", nil)
}

// GetNotifications returns recent notifications for the logged-in user
func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	notifs, err := h.service.GetMyNotifications(userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch notifications")
		return
	}

	response.Success(c, http.StatusOK, "Notifications retrieved", notifs)
}
