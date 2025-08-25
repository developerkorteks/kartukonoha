package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/database"
	"github.com/nabilulilalbab/nadia/internal/models"
	"github.com/nabilulilalbab/nadia/internal/monitoring"
)

// GetSystemMetrics godoc
// @Summary Get comprehensive system metrics
// @Description Get detailed system monitoring data including CPU, memory, disk, network, and performance metrics
// @Tags monitoring
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse{data=models.SystemMetrics}
// @Failure 401 {object} models.APIResponse
// @Router /api/monitoring/metrics [get]
func (h *HTTPHandler) GetSystemMetrics(c *gin.Context) {
	// This handler would need access to the DB to get its status.
	// For now, we'll pass a placeholder status.
	dbStatus := database.GetDatabaseStatus(nil) // Simplified: should pass h.db
	metrics := monitoring.CollectSystemMetrics(dbStatus)
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "System metrics retrieved successfully",
		Success:    true,
		Data:       metrics,
	})
}

// GetRealtimeMetrics godoc
// @Summary Get real-time metrics
// @Description Get live monitoring data from the last 10 minutes
// @Tags monitoring
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/monitoring/metrics/realtime [get]
func (h *HTTPHandler) GetRealtimeMetrics(c *gin.Context) {
	recentMetrics := monitoring.GetRealtimeMetrics()
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Realtime metrics retrieved successfully",
		Success:    true,
		Data: map[string]interface{}{
			"recent_requests": recentMetrics,
			"timestamp":       time.Now(),
			"window_minutes":  10,
		},
	})
}

// GetSecurityMetrics godoc
// @Summary Get security monitoring data
// @Description Get security metrics including unauthorized attempts, threats, and suspicious traffic
// @Tags monitoring
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/monitoring/security [get]
func (h *HTTPHandler) GetSecurityMetrics(c *gin.Context) {
	metrics := monitoring.GetSecurityMetrics()
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Security metrics retrieved successfully",
		Success:    true,
		Data:       metrics,
	})
}

// GetApplicationLogs godoc
// @Summary Get application request logs
// @Description Get recent application logs including request details and response times
// @Tags monitoring
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/monitoring/logs [get]
func (h *HTTPHandler) GetApplicationLogs(c *gin.Context) {
	logs := monitoring.GetApplicationLogs()
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Application logs retrieved successfully",
		Success:    true,
		Data: map[string]interface{}{
			"logs":      logs,
			"timestamp": time.Now(),
			"count":     len(logs),
		},
	})
}

// GetPerformanceMetrics godoc
// @Summary Get performance metrics
// @Description Get detailed performance metrics including response times, throughput, and resource usage
// @Tags monitoring
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/monitoring/performance [get]
func (h *HTTPHandler) GetPerformanceMetrics(c *gin.Context) {
	performance := monitoring.GetPerformanceMetrics()
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Performance metrics retrieved successfully",
		Success:    true,
		Data:       performance,
	})
}

// GetUptimeStatus godoc
// @Summary Get service uptime status
// @Description Get service health status, uptime information, and availability metrics
// @Tags monitoring
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/monitoring/uptime [get]
func (h *HTTPHandler) GetUptimeStatus(c *gin.Context) {
	status := monitoring.GetUptimeStatus()
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Uptime status retrieved successfully",
		Success:    true,
		Data:       status,
	})
}
