package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, h *handlers.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}
}
