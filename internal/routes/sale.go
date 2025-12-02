package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterSaleRoutes(rg *gin.RouterGroup, h *handlers.SaleHandler) {
	sales := rg.Group("/sales")
	{
		sales.POST("", h.CreateSale)
		sales.GET("", h.GetSales)
		sales.GET("/:id", h.GetSaleByID)
		sales.POST("/:id/refund", h.RefundSale)
	}
}
