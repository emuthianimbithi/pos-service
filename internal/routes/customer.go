package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoutes(rg *gin.RouterGroup, h *handlers.CustomerHandler) {
	customers := rg.Group("/customers")
	{
		customers.POST("", h.CreateCustomer)
		customers.GET("", h.ListCustomers)
		customers.GET("/:id", h.GetCustomer)
		customers.PUT("/:id", h.UpdateCustomer)
		customers.DELETE("/:id", h.DeleteCustomer)
	}
}
