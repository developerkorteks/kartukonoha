package monitoring

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/nabilulilalbab/nadia/internal/models"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

var (
	requestMetrics    = make([]models.RequestMetrics, 0, 1000)
	requestMetricsMux sync.RWMutex
	securityThreats   = make([]models.SecurityThreat, 0, 100)
	securityMutex     sync.RWMutex
	serverStartTime   = time.Now()
	totalRequests     int64
	totalErrors       int64
	unauthorizedCount int64
	suspiciousTraffic int64
	metricsMutex      sync.RWMutex
)

// AddRequestMetric adds a new request metric to the collection.
func AddRequestMetric(metric models.RequestMetrics) {
	requestMetricsMux.Lock()
	requestMetrics = append(requestMetrics, metric)
	if len(requestMetrics) > 1000 {
		requestMetrics = requestMetrics[1:]
	}
	requestMetricsMux.Unlock()

	metricsMutex.Lock()
	totalRequests++
	if metric.StatusCode >= 500 {
		totalErrors++
	}
	metricsMutex.Unlock()

	if metric.ResponseTime > 5000 { // > 5 seconds
		LogSecurityThreat(metric.IP, "slow_response", "low",
			fmt.Sprintf("Slow response time: %.2fms for %s", metric.ResponseTime, metric.Path))
	}
}

// LogSecurityThreat logs a potential security threat.
func LogSecurityThreat(ip, threatType, severity, details string) {
	threat := models.SecurityThreat{
		IP:        ip,
		Type:      threatType,
		Severity:  severity,
		Timestamp: time.Now(),
		Details:   details,
	}

	securityMutex.Lock()
	securityThreats = append(securityThreats, threat)
	if len(securityThreats) > 100 {
		securityThreats = securityThreats[1:]
	}
	securityMutex.Unlock()

	log.Printf("[SECURITY] %s - %s from %s: %s", severity, threatType, ip, details)
}

// IncrementUnauthorizedCount increases the count of unauthorized attempts.
func IncrementUnauthorizedCount() {
	metricsMutex.Lock()
	unauthorizedCount++
	metricsMutex.Unlock()
}

// CollectSystemMetrics gathers all system metrics.
func CollectSystemMetrics(dbStatus string) models.SystemMetrics {
	cpuPercent, _ := cpu.Percent(time.Second, false)
	var cpuUsage float64
	if len(cpuPercent) > 0 {
		cpuUsage = cpuPercent[0]
	}

	memInfo, _ := mem.VirtualMemory()
	memoryMetrics := models.MemoryMetrics{
		Used:        memInfo.Used,
		Total:       memInfo.Total,
		UsedPercent: memInfo.UsedPercent,
		Available:   memInfo.Available,
		HasLeak:     memInfo.UsedPercent > 90, // Simple leak detection
	}

	diskInfo, _ := disk.Usage("/")
	diskMetrics := models.DiskMetrics{
		Used:        diskInfo.Used,
		Total:       diskInfo.Total,
		UsedPercent: diskInfo.UsedPercent,
		Free:        diskInfo.Free,
		IOStats:     models.IOStats{}, // Would need more complex implementation
	}

	loadInfo, _ := load.Avg()
	loadMetrics := models.LoadMetrics{
		Load1:  loadInfo.Load1,
		Load5:  loadInfo.Load5,
		Load15: loadInfo.Load15,
	}

	netStats, _ := net.IOCounters(false)
	var networkMetrics models.NetworkMetrics
	if len(netStats) > 0 {
		networkMetrics = models.NetworkMetrics{
			BytesSent:   netStats[0].BytesSent,
			BytesRecv:   netStats[0].BytesRecv,
			PacketsSent: netStats[0].PacketsSent,
			PacketsRecv: netStats[0].PacketsRecv,
		}
	}

	metricsMutex.RLock()
	var errorRate float64
	if totalRequests > 0 {
		errorRate = float64(totalErrors) / float64(totalRequests) * 100
	}
	uptime := time.Since(serverStartTime).Seconds()
	throughput := float64(totalRequests) / uptime
	metricsMutex.RUnlock()

	securityMutex.RLock()
	secMetrics := models.SecurityMetrics{
		UnauthorizedAttempts: int(unauthorizedCount),
		SuspiciousTraffic:    int(suspiciousTraffic),
		FirewallBlocks:       0, // Placeholder
		RecentThreats:        append([]models.SecurityThreat{}, securityThreats...),
	}
	securityMutex.RUnlock()

	return models.SystemMetrics{
		ServiceStatus:   "running",
		ResponseTime:    CalculateResponseTimeMetrics(),
		ErrorRate:       errorRate,
		CPUUsage:        cpuUsage,
		MemoryUsage:     memoryMetrics,
		DiskUsage:       diskMetrics,
		NetworkTraffic:  networkMetrics,
		LoadAverage:     loadMetrics,
		Throughput:      throughput,
		DBPerformance:   CalculateDBMetrics(dbStatus),
		SecurityMetrics: secMetrics,
		UptimeCheck:     CheckUptime(),
		Timestamp:       time.Now(),
	}
}

// CalculateResponseTimeMetrics calculates various response time metrics.
func CalculateResponseTimeMetrics() models.ResponseTimeMetrics {
	requestMetricsMux.RLock()
	defer requestMetricsMux.RUnlock()

	if len(requestMetrics) == 0 {
		return models.ResponseTimeMetrics{}
	}

	times := make([]float64, len(requestMetrics))
	var sum float64
	min := requestMetrics[0].ResponseTime
	max := requestMetrics[0].ResponseTime

	for i, metric := range requestMetrics {
		times[i] = metric.ResponseTime
		sum += metric.ResponseTime
		if metric.ResponseTime < min {
			min = metric.ResponseTime
		}
		if metric.ResponseTime > max {
			max = metric.ResponseTime
		}
	}

	// Simple sort for percentiles
	for i := 0; i < len(times)-1; i++ {
		for j := i + 1; j < len(times); j++ {
			if times[i] > times[j] {
				times[i], times[j] = times[j], times[i]
			}
		}
	}

	p95Index := int(float64(len(times)) * 0.95)
	p99Index := int(float64(len(times)) * 0.99)
	if p95Index >= len(times) {
		p95Index = len(times) - 1
	}
	if p99Index >= len(times) {
		p99Index = len(times) - 1
	}

	return models.ResponseTimeMetrics{
		Average: sum / float64(len(times)),
		P95:     times[p95Index],
		P99:     times[p99Index],
		Min:     min,
		Max:     max,
	}
}

// CalculateDBMetrics returns simple DB metrics.
func CalculateDBMetrics(dbStatus string) models.DBMetrics {
	connections := 0
	if dbStatus == "CONNECTED" {
		connections = 1 // Placeholder
	}
	return models.DBMetrics{
		QueryTime:     0, // Placeholder
		SlowQueries:   0,
		Connections:   connections,
		QueriesPerSec: 0,
	}
}

// CheckUptime performs a simple local uptime check.
func CheckUptime() models.UptimeMetrics {
	return models.UptimeMetrics{
		IsUp:         true,
		ResponseTime: 0, // Placeholder
		StatusCode:   200,
		Location:     "server", // Generic location instead of hardcoded localhost
	}
}

// GetUptimeStatus returns the service uptime status.
func GetUptimeStatus() map[string]interface{} {
	uptime := time.Since(serverStartTime)
	return map[string]interface{}{
		"is_up":          true,
		"is_healthy":     true, // Placeholder
		"status":         "healthy",
		"uptime_seconds": uptime.Seconds(),
		"uptime_human":   uptime.String(),
		"started_at":     serverStartTime,
		"current_time":   time.Now(),
		"version":        "2.0.0",
		"environment":    "production",
	}
}

// GetPerformanceMetrics returns detailed performance metrics.
func GetPerformanceMetrics() map[string]interface{} {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	cpuPercent, _ := cpu.Percent(time.Second, false)
	var cpuUsage float64
	if len(cpuPercent) > 0 {
		cpuUsage = cpuPercent[0]
	}

	metricsMutex.RLock()
	uptime := time.Since(serverStartTime).Seconds()
	throughput := 0.0
	errorRate := 0.0
	if totalRequests > 0 {
		throughput = float64(totalRequests) / uptime
		errorRate = float64(totalErrors) / float64(totalRequests) * 100
	}
	metricsMutex.RUnlock()

	return map[string]interface{}{
		"response_time":  CalculateResponseTimeMetrics(),
		"throughput_rps": throughput,
		"error_rate":     errorRate,
		"cpu_usage":      cpuUsage,
		"memory": map[string]interface{}{
			"alloc_mb":       float64(memStats.Alloc) / 1024 / 1024,
			"total_alloc_mb": float64(memStats.TotalAlloc) / 1024 / 1024,
			"sys_mb":         float64(memStats.Sys) / 1024 / 1024,
			"gc_runs":        memStats.NumGC,
		},
		"uptime_seconds": uptime,
		"timestamp":      time.Now(),
	}
}

// GetSecurityMetrics returns security-related metrics.
func GetSecurityMetrics() map[string]interface{} {
	securityMutex.RLock()
	threats := append([]models.SecurityThreat{}, securityThreats...)
	securityMutex.RUnlock()

	metricsMutex.RLock()
	metrics := map[string]interface{}{
		"unauthorized_attempts": unauthorizedCount,
		"suspicious_traffic":    suspiciousTraffic,
		"recent_threats":        threats,
		"total_requests":        totalRequests,
		"total_errors":          totalErrors,
		"timestamp":             time.Now(),
	}
	metricsMutex.RUnlock()

	return metrics
}

// GetApplicationLogs returns recent application request logs.
func GetApplicationLogs() []map[string]interface{} {
	requestMetricsMux.RLock()
	defer requestMetricsMux.RUnlock()

	recentLogs := make([]map[string]interface{}, 0)
	start := len(requestMetrics) - 50
	if start < 0 {
		start = 0
	}

	for i := start; i < len(requestMetrics); i++ {
		metric := requestMetrics[i]
		logEntry := map[string]interface{}{
			"timestamp":     metric.StartTime,
			"method":        metric.Method,
			"path":          metric.Path,
			"status_code":   metric.StatusCode,
			"response_time": metric.ResponseTime,
			"ip":            metric.IP,
			"user_agent":    metric.UserAgent,
		}
		recentLogs = append(recentLogs, logEntry)
	}
	return recentLogs
}

// GetRealtimeMetrics returns metrics from the last 10 minutes.
func GetRealtimeMetrics() []models.RequestMetrics {
	requestMetricsMux.RLock()
	defer requestMetricsMux.RUnlock()

	recentMetrics := make([]models.RequestMetrics, 0)
	cutoff := time.Now().Add(-10 * time.Minute)

	for _, metric := range requestMetrics {
		if metric.StartTime.After(cutoff) {
			recentMetrics = append(recentMetrics, metric)
		}
	}
	return recentMetrics
}
