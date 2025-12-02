package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterShiftRoutes(rg *gin.RouterGroup, h *handlers.ShiftHandler) {
	shifts := rg.Group("/shifts")
	{
		shifts.POST("/open", h.OpenShift)
		shifts.POST("/close", h.CloseShift)
		shifts.GET("", h.GetShifts)
	}
}
