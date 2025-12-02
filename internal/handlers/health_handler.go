package handlers

import (
	"net/http"
	"time"

	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Services  map[string]string `json:"services"`
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	services := make(map[string]string)
	status := "healthy"
	statusCode := http.StatusOK

	// Check Database
	sqlDB, err := h.db.DB()
	if err != nil {
		services["database"] = "error: " + err.Error()
		status = "unhealthy"
		statusCode = http.StatusServiceUnavailable
	} else {
		if err := sqlDB.Ping(); err != nil {
			services["database"] = "down: " + err.Error()
			status = "unhealthy"
			statusCode = http.StatusServiceUnavailable
		} else {
			services["database"] = "up"
		}
	}

	// Add other checks here if needed (e.g., Redis, external APIs)

	response.Success(c, statusCode, "Health check completed", HealthResponse{
		Status:    status,
		Timestamp: time.Now(),
		Services:  services,
	})
}
