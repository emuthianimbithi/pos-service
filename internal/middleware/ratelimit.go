package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/config"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// RateLimiter stores rate limiters per IP address
type RateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
	rate     rate.Limit
	burst    int
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(cfg *config.Config) *RateLimiter {
	r := rate.Limit(float64(cfg.RateLimitRequests) / float64(cfg.RateLimitWindow))
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
		rate:     r,
		burst:    cfg.RateLimitRequests,
	}
}

// GetLimiter returns a rate limiter for the given key (usually IP address)
func (rl *RateLimiter) GetLimiter(key string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.limiters[key]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.limiters[key] = limiter
	}

	return limiter
}

// CleanupOldLimiters removes inactive limiters periodically
func (rl *RateLimiter) CleanupOldLimiters() {
	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for range ticker.C {
			rl.mu.Lock()
			// In a production system, you'd track last access time
			// For now, we'll just clear all limiters periodically
			rl.limiters = make(map[string]*rate.Limiter)
			rl.mu.Unlock()
		}
	}()
}

// RateLimitMiddleware applies rate limiting per IP address
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := rl.GetLimiter(ip)

		if !limiter.Allow() {
			response.Error(c, http.StatusTooManyRequests, "Rate limit exceeded. Please try again later.")
			c.Abort()
			return
		}

		c.Next()
	}
}

// RateLimitMiddlewareByUser applies rate limiting per authenticated user
func RateLimitMiddlewareByUser(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Try to get user ID from context (set by AuthMiddleware)
		userID, exists := c.Get("user_id")
		if !exists {
			// Fall back to IP-based rate limiting
			ip := c.ClientIP()
			limiter := rl.GetLimiter(ip)
			if !limiter.Allow() {
				response.Error(c, http.StatusTooManyRequests, "Rate limit exceeded. Please try again later.")
				c.Abort()
				return
			}
		} else {
			// Use user ID for rate limiting
			key := userID.(string)
			limiter := rl.GetLimiter(key)
			if !limiter.Allow() {
				response.Error(c, http.StatusTooManyRequests, "Rate limit exceeded. Please try again later.")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
