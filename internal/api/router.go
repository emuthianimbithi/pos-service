package api

import (
	"github.com/emuthianimbithi/pos-service/internal/container"
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/emuthianimbithi/pos-service/internal/middleware"
	"github.com/emuthianimbithi/pos-service/internal/routes"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router with all routes and middleware
func SetupRouter(c *container.Container) *gin.Engine {
	router := gin.New()

	// Initialize Handlers
	mpesaHandler := handlers.NewMpesaHandler(c.Services.Mpesa)
	notificationHandler := handlers.NewNotificationHandler(c.Services.Alert)
	authHandler := handlers.NewAuthHandler(c.Services.Auth)
	productHandler := handlers.NewProductHandler(c.Services.Product)
	inventoryHandler := handlers.NewInventoryHandler(c.Services.Inventory)
	shiftHandler := handlers.NewShiftHandler(c.Services.Shift)
	saleHandler := handlers.NewSaleHandler(c.Services.Sale)
	customerHandler := handlers.NewCustomerHandler(c.Services.Customer)
	businessHandler := handlers.NewBusinessHandler(c.Services.Business)
	branchHandler := handlers.NewBranchHandler(c.Services.Branch)
	userHandler := handlers.NewUserHandler(c.Services.User)
	reportHandler := handlers.NewReportHandler(c.Services.Report)
	healthHandler := handlers.NewHealthHandler(c.DB)

	// Global Middleware
	router.Use(gin.Recovery())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RateLimitMiddleware(middleware.NewRateLimiter(c.Config)))

	// Audit Middleware
	router.Use(middleware.AuditMiddleware(c.DB))

	// Public Routes
	public := router.Group("/api/v1")
	{
		public.GET("/health", healthHandler.HealthCheck)
		routes.RegisterAuthRoutes(public, authHandler)
	}

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		v1.POST("/mpesa/callback", mpesaHandler.Callback)

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(c.Config))
		protected.Use(middleware.TenantMiddleware())
		{
			routes.RegisterMpesaRoutes(protected, mpesaHandler)
			routes.RegisterNotificationRoutes(protected, notificationHandler)

			// Product & Inventory
			protected.Use(middleware.RequireBranch()) // Ensure branch context for these
			{
				routes.RegisterProductRoutes(protected, productHandler)
				routes.RegisterInventoryRoutes(protected, inventoryHandler)
				routes.RegisterShiftRoutes(protected, shiftHandler)
				routes.RegisterSaleRoutes(protected, saleHandler)
				routes.RegisterCustomerRoutes(protected, customerHandler)
				routes.RegisterBusinessRoutes(protected, businessHandler)
				routes.RegisterBranchRoutes(protected, branchHandler)
				routes.RegisterUserRoutes(protected, userHandler)
				routes.RegisterReportRoutes(protected, reportHandler)
			}
		}
	}

	return router
}
