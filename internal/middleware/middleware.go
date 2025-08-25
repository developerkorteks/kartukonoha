package middleware

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/models"
	"github.com/nabilulilalbab/nadia/internal/monitoring"
)

// AuthMiddleware creates a middleware for API key authentication.
func AuthMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for health check, swagger, and other public endpoints
		if isPublicPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		providedKey := c.GetHeader("X-API-Key")
		if providedKey == "" {
			providedKey = c.Query("api_key")
		}

		if providedKey != apiKey {
			monitoring.LogSecurityThreat(c.ClientIP(), "unauthorized_access", "medium",
				fmt.Sprintf("Invalid API key attempt from %s", c.ClientIP()))
			monitoring.IncrementUnauthorizedCount()

			c.JSON(401, models.APIResponse{
				StatusCode: 401,
				Message:    "Unauthorized: Invalid API key",
				Success:    false,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// MonitoringMiddleware records metrics for each request.
func MonitoringMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		// Don't log metrics for swagger assets
		if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
			return
		}

		endTime := time.Now()
		responseTime := float64(endTime.Sub(startTime).Nanoseconds()) / 1e6 // in ms

		metric := models.RequestMetrics{
			StartTime:    startTime,
			EndTime:      endTime,
			StatusCode:   c.Writer.Status(),
			ResponseTime: responseTime,
			Path:         c.Request.URL.Path,
			Method:       c.Request.Method,
			IP:           c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
		}
		monitoring.AddRequestMetric(metric)

		log.Printf("[%s] %s %s - %d - %.2fms - %s",
			c.Request.Method, c.Request.URL.Path, c.ClientIP(),
			c.Writer.Status(), responseTime, c.Request.UserAgent())
	}
}

// CORSMiddleware sets the CORS headers for the API.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func isPublicPath(path string) bool {
	publicPaths := []string{
		"/api/health",
		"/api/auth/login",
		"/",
		"/dashboard",
		"/dashboard/old",
		"/api-flow",
	}

	for _, p := range publicPaths {
		if p == path {
			return true
		}
	}

	if strings.HasPrefix(path, "/swagger/") {
		return true
	}

	return false
}
