package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBusinessRoutes(rg *gin.RouterGroup, h *handlers.BusinessHandler) {
	businesses := rg.Group("/businesses")
	{
		businesses.POST("", h.CreateBusiness)
		businesses.GET("", h.ListBusinesses)
		businesses.GET("/:id", h.GetBusiness)
		businesses.PUT("/:id", h.UpdateBusiness)
		businesses.DELETE("/:id", h.DeleteBusiness)
	}
}
