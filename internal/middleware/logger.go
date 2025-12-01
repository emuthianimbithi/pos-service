package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs HTTP requests with detailed information
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get status code
		statusCode := c.Writer.Status()

		// Log format: [timestamp] method path query status latency ip
		log.Printf("[%s] %s %s %s %d %v %s",
			start.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			path,
			query,
			statusCode,
			latency,
			c.ClientIP(),
		)
	}
}
