package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler) {
	users := rg.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("", h.ListUsers)
		users.GET("/:id", h.GetUserProfile)
		users.PUT("/:id", h.UpdateUserProfile)
		users.PUT("/:id/password", h.ChangePassword)
		users.DELETE("/:id", h.DeactivateUser)
	}
}
