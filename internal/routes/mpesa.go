package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/emuthianimbithi/pos-service/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterMpesaRoutes(rg *gin.RouterGroup, h *handlers.MpesaHandler) {
	mpesa := rg.Group("/mpesa")
	mpesa.Use(middleware.RequireBranch())
	{
		mpesa.POST("/config", middleware.RequireRole("admin", "manager"), h.CreateConfig)
		mpesa.POST("/stkpush", h.InitiateSTKPush)
		mpesa.GET("/status/:id", h.CheckStatus)
	}
}

func RegisterPublicMpesaRoutes(rg *gin.RouterGroup, h *handlers.MpesaHandler) {
	rg.POST("/mpesa/callback", h.Callback)
}
