package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, h *handlers.ProductHandler) {
	products := rg.Group("/products")
	{
		products.GET("", h.GetProducts)
		products.POST("", h.CreateProduct)
		products.GET("/:id", h.GetProduct)
		products.PUT("/:id", h.UpdateProduct)
		products.DELETE("/:id", h.DeleteProduct)
		products.POST("/:id/variants", h.CreateVariant)
		products.PUT("/:id/variants/:variant_id", h.UpdateVariant)
		products.DELETE("/:id/variants/:variant_id", h.DeleteVariant)
	}

	categories := rg.Group("/categories")
	{
		categories.GET("", h.GetCategories)
		categories.POST("", h.CreateCategory)
	}
}

func RegisterInventoryRoutes(rg *gin.RouterGroup, h *handlers.InventoryHandler) {
	inventory := rg.Group("/inventory")
	{
		inventory.POST("/adjust", h.AdjustStock)
	}
}
