package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterNotificationRoutes(rg *gin.RouterGroup, h *handlers.NotificationHandler) {
	notifs := rg.Group("/notifications")
	{
		notifs.GET("/alerts", h.GetAlerts)
		notifs.PUT("/alerts/:id/read", h.MarkRead)
		notifs.GET("/recent", h.GetNotifications)
	}
}
