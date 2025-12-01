package handlers

import (
	"github.com/emuthianimbithi/pos-service/internal/config"
	"github.com/emuthianimbithi/pos-service/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter initializes the Gin router with all routes and middleware
func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.New()

	// Initialize rate limiter
	rateLimiter := middleware.NewRateLimiter(cfg)
	rateLimiter.CleanupOldLimiters()

	// Global middleware
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RateLimitMiddleware(rateLimiter))
	router.Use(middleware.AuditMiddleware(db))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"service": "pos-service",
		})
	})

	// Initialize handlers
	mpesaHandler := NewMpesaHandler(db)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		// Callback must be public for Safaricom to reach it
		v1.POST("/mpesa/callback", mpesaHandler.Callback)

		// Protected routes (authentication required)
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		protected.Use(middleware.TenantMiddleware())
		{
			// M-Pesa routes (Branch scoped)
			mpesa := protected.Group("/mpesa")
			mpesa.Use(middleware.RequireBranch())
			{
				mpesa.POST("/config", middleware.RequireRole("admin", "manager"), mpesaHandler.CreateConfig)
				mpesa.POST("/stkpush", mpesaHandler.InitiateSTKPush)
				mpesa.GET("/status/:id", mpesaHandler.CheckStatus)
			}
			// Business routes (Admin only)
			// businesses := protected.Group("/businesses")
			// businesses.Use(middleware.RequireRole("admin"))
			// {
			// 	// TODO: Add business routes
			// 	// businesses.GET("", businessHandler.List)
			// 	// businesses.POST("", businessHandler.Create)
			// 	// businesses.GET("/:id", businessHandler.Get)
			// 	// businesses.PUT("/:id", businessHandler.Update)
			// 	// businesses.DELETE("/:id", businessHandler.Delete)
			// }

			// Branch routes (Business scoped)
			// branches := protected.Group("/branches")
			// branches.Use(middleware.RequireBusiness())
			// {
			// 	// TODO: Add branch routes
			// 	// branches.GET("", branchHandler.List)
			// 	// branches.POST("", branchHandler.Create)
			// 	// branches.GET("/:id", branchHandler.Get)
			// 	// branches.PUT("/:id", branchHandler.Update)
			// 	// branches.DELETE("/:id", branchHandler.Delete)
			// }

			// Product routes (Branch scoped)
			// products := protected.Group("/products")
			// products.Use(middleware.RequireBranch())
			// {
			// 	// TODO: Add product routes
			// 	// products.GET("", productHandler.List)
			// 	// products.POST("", productHandler.Create)
			// 	// products.GET("/:id", productHandler.Get)
			// 	// products.PUT("/:id", productHandler.Update)
			// 	// products.DELETE("/:id", productHandler.Delete)
			// }

			// Sales routes (Branch scoped)
			// sales := protected.Group("/sales")
			// sales.Use(middleware.RequireBranch())
			// {
			// 	// TODO: Add sales routes
			// 	// sales.GET("", saleHandler.List)
			// 	// sales.POST("", saleHandler.Create)
			// 	// sales.GET("/:id", saleHandler.Get)
			// }

			// Customer routes (Branch scoped)
			// customers := protected.Group("/customers")
			// customers.Use(middleware.RequireBranch())
			// {
			// 	// TODO: Add customer routes
			// 	// customers.GET("", customerHandler.List)
			// 	// customers.POST("", customerHandler.Create)
			// 	// customers.GET("/:id", customerHandler.Get)
			// 	// customers.PUT("/:id", customerHandler.Update)
			// 	// customers.DELETE("/:id", customerHandler.Delete)
			// }
		}
	}

	return router
}
