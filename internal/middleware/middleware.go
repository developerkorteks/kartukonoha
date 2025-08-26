package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/models"
	"github.com/nabilulilalbab/nadia/internal/monitoring"
)

// AuthMiddleware creates a middleware for API key authentication.
// It supports both static API Key (X-API-Key) and Bearer Token (Authorization header).
func AuthMiddleware(apiKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for health check, swagger, and other public endpoints
		if isPublicPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		// 1. Check for static API Key in X-API-Key header
		providedKey := c.GetHeader("X-API-Key")
		if providedKey == apiKey {
			c.Next()
			return
		}

		// 2. Check for Bearer Token in Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			// Check if the format is "Bearer <token>"
			if len(parts) == 2 && parts[0] == "Bearer" {
				token := parts[1]
				// Simple validation for the generated token. In a real-world scenario, this should be a proper JWT validation.
				if strings.HasPrefix(token, "nadia-token-") {
					c.Next()
					return
				}
			}
		}

		// 3. Fallback for Swagger UI: Check if the generated token was mistakenly put in the X-API-Key field.
		// This improves user experience when testing with Swagger.
		if strings.HasPrefix(providedKey, "nadia-token-") {
			c.Next()
			return
		}

		// If all checks fail, deny access.
		monitoring.LogSecurityThreat(c.ClientIP(), "unauthorized_access", "medium",
			fmt.Sprintf("Invalid API key or token attempt from %s", c.ClientIP()))
		monitoring.IncrementUnauthorizedCount()

		c.JSON(401, models.APIResponse{
			StatusCode: 401,
			Message:    "Unauthorized: Invalid API key or token",
			Success:    false,
		})
		c.Abort()
	}
}

// JWTClaims represents the JWT token claims
type JWTClaims struct {
	Username string `json:"username"`
	Exp      int64  `json:"exp"`
}

// JWTMiddleware creates a middleware for JWT token authentication for user endpoints
func JWTMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Authorization header required",
				Success:    false,
			})
			c.Abort()
			return
		}

		// Check Bearer token format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.APIResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid authorization header format",
				Success:    false,
			})
			c.Abort()
			return
		}

		token := parts[1]

		// Simple JWT validation (in production, use proper JWT library)
		if !validateJWT(token, jwtSecret) {
			monitoring.LogSecurityThreat(c.ClientIP(), "invalid_jwt_token", "medium",
				fmt.Sprintf("Invalid JWT token attempt from %s", c.ClientIP()))
			monitoring.IncrementUnauthorizedCount()

			c.JSON(http.StatusUnauthorized, models.APIResponse{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid or expired token",
				Success:    false,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// validateJWT validates a simple JWT token (basic implementation)
func validateJWT(token, secret string) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}

	// Decode header and payload
	_, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	// Verify signature
	expectedSignature := generateJWTSignature(parts[0]+"."+parts[1], secret)
	if parts[2] != expectedSignature {
		return false
	}

	// Check expiration
	var claims JWTClaims
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return false
	}

	if time.Now().Unix() > claims.Exp {
		return false
	}

	return true
}

// generateJWTSignature generates HMAC-SHA256 signature for JWT
func generateJWTSignature(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// HybridAuthMiddleware creates a middleware that accepts both API Key and JWT authentication
func HybridAuthMiddleware(apiKey, jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for health check, swagger, and other public endpoints
		if isPublicPath(c.Request.URL.Path) {
			c.Next()
			return
		}

		// 1. Check for static API Key in X-API-Key header
		providedKey := c.GetHeader("X-API-Key")
		if providedKey == apiKey {
			c.Next()
			return
		}

		// 2. Check for Bearer Token in Authorization header (JWT)
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			// Check if the format is "Bearer <token>"
			if len(parts) == 2 && parts[0] == "Bearer" {
				token := parts[1]

				// First try simple token validation (legacy tokens)
				if strings.HasPrefix(token, "nadia-token-") {
					c.Next()
					return
				}

				// Then try JWT validation
				if validateJWT(token, jwtSecret) {
					c.Next()
					return
				}
			}
		}

		// 3. Fallback for Swagger UI: Check if the generated token was mistakenly put in the X-API-Key field.
		if strings.HasPrefix(providedKey, "nadia-token-") {
			c.Next()
			return
		}

		// If all checks fail, deny access.
		monitoring.LogSecurityThreat(c.ClientIP(), "unauthorized_access", "medium",
			fmt.Sprintf("Invalid API key or token attempt from %s", c.ClientIP()))
		monitoring.IncrementUnauthorizedCount()

		c.JSON(http.StatusUnauthorized, models.APIResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized: Invalid API key or token",
			Success:    false,
		})
		c.Abort()
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

		// Extract source from request body if available
		source := "unknown"
		if c.Request.Method == "POST" {
			// Try to extract source from request body
			if bodyBytes, err := io.ReadAll(c.Request.Body); err == nil {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				var requestData map[string]interface{}
				if json.Unmarshal(bodyBytes, &requestData) == nil {
					if sourceValue, exists := requestData["source"]; exists {
						if sourceStr, ok := sourceValue.(string); ok {
							source = sourceStr
						}
					}
				}
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		log.Printf("[%s] %s %s - %d - %.2fms - %s - Source: %s",
			c.Request.Method, c.Request.URL.Path, c.ClientIP(),
			c.Writer.Status(), responseTime, c.Request.UserAgent(), source)
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
