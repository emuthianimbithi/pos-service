package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterReportRoutes(rg *gin.RouterGroup, h *handlers.ReportHandler) {
	reports := rg.Group("/reports")
	{
		reports.GET("/sales", h.GetSalesReport)
		reports.GET("/inventory", h.GetInventoryReport)
		reports.GET("/low-stock", h.GetLowStockReport)
		reports.GET("/customers", h.GetCustomerReport)
	}
}
