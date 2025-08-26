package models

import "time"

// APIResponse structure
type APIResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
}

// Package structure
type Package struct {
	PackageCode             string          `json:"package_code"`
	PackageName             string          `json:"package_name"`
	PackageNameShort        string          `json:"package_name_alias_short"`
	PackageDescription      string          `json:"package_description"`
	PackagePrice            int             `json:"package_harga_int"`
	PackagePriceFormatted   string          `json:"package_harga"`
	HaveDailyLimit          bool            `json:"have_daily_limit"`
	DailyLimitDetails       DailyLimit      `json:"daily_limit_details"`
	NoNeedLogin             bool            `json:"no_need_login"`
	CanMultiTrx             bool            `json:"can_multi_trx"`
	CanScheduledTrx         bool            `json:"can_scheduled_trx"`
	HaveCutOffTime          bool            `json:"have_cut_off_time"`
	CutOffTime              CutOffTime      `json:"cut_off_time"`
	NeedCheckStock          bool            `json:"need_check_stock"`
	IsShowPaymentMethod     bool            `json:"is_show_payment_method"`
	AvailablePaymentMethods []PaymentMethod `json:"available_payment_methods"`
}

type DailyLimit struct {
	MaxDailyTransactionLimit     int `json:"max_daily_transaction_limit"`
	CurrentDailyTransactionCount int `json:"current_daily_transaction_count"`
}

type CutOffTime struct {
	ProhibitedHourStarttime string `json:"prohibited_hour_starttime"`
	ProhibitedHourEndtime   string `json:"prohibited_hour_endtime"`
}

type PaymentMethod struct {
	Order                    int    `json:"order"`
	PaymentMethod            string `json:"payment_method"`
	PaymentMethodDisplayName string `json:"payment_method_display_name"`
	Description              string `json:"desc"`
}

// Product structure for user endpoints (with manipulated price)
type Product struct {
	PackageCode             string          `json:"package_code"`
	PackageName             string          `json:"package_name"`
	PackageNameShort        string          `json:"package_name_alias_short"`
	PackageDescription      string          `json:"package_description"`
	PackagePrice            int             `json:"package_harga_int"` // Original price + 1500
	PackagePriceFormatted   string          `json:"package_harga"`     // Formatted manipulated price
	HaveDailyLimit          bool            `json:"have_daily_limit"`
	DailyLimitDetails       DailyLimit      `json:"daily_limit_details"`
	NoNeedLogin             bool            `json:"no_need_login"`
	CanMultiTrx             bool            `json:"can_multi_trx"`
	CanScheduledTrx         bool            `json:"can_scheduled_trx"`
	HaveCutOffTime          bool            `json:"have_cut_off_time"`
	CutOffTime              CutOffTime      `json:"cut_off_time"`
	NeedCheckStock          bool            `json:"need_check_stock"`
	IsShowPaymentMethod     bool            `json:"is_show_payment_method"`
	AvailablePaymentMethods []PaymentMethod `json:"available_payment_methods"`
}

// Monitoring structures
type SystemMetrics struct {
	ServiceStatus   string              `json:"service_status"`
	ResponseTime    ResponseTimeMetrics `json:"response_time"`
	ErrorRate       float64             `json:"error_rate"`
	CPUUsage        float64             `json:"cpu_usage"`
	MemoryUsage     MemoryMetrics       `json:"memory_usage"`
	DiskUsage       DiskMetrics         `json:"disk_usage"`
	NetworkTraffic  NetworkMetrics      `json:"network_traffic"`
	LoadAverage     LoadMetrics         `json:"load_average"`
	Throughput      float64             `json:"throughput"`
	DBPerformance   DBMetrics           `json:"db_performance"`
	SecurityMetrics SecurityMetrics     `json:"security_metrics"`
	UptimeCheck     UptimeMetrics       `json:"uptime_check"`
	Timestamp       time.Time           `json:"timestamp"`
}

type ResponseTimeMetrics struct {
	Average float64 `json:"average_ms"`
	P95     float64 `json:"p95_ms"`
	P99     float64 `json:"p99_ms"`
	Min     float64 `json:"min_ms"`
	Max     float64 `json:"max_ms"`
}

type MemoryMetrics struct {
	Used        uint64  `json:"used_bytes"`
	Total       uint64  `json:"total_bytes"`
	UsedPercent float64 `json:"used_percent"`
	Available   uint64  `json:"available_bytes"`
	HasLeak     bool    `json:"has_memory_leak"`
}

type DiskMetrics struct {
	Used        uint64  `json:"used_bytes"`
	Total       uint64  `json:"total_bytes"`
	UsedPercent float64 `json:"used_percent"`
	Free        uint64  `json:"free_bytes"`
	IOStats     IOStats `json:"io_stats"`
}

type IOStats struct {
	ReadBytes  uint64 `json:"read_bytes"`
	WriteBytes uint64 `json:"write_bytes"`
	ReadOps    uint64 `json:"read_ops"`
	WriteOps   uint64 `json:"write_ops"`
}

type NetworkMetrics struct {
	BytesSent   uint64  `json:"bytes_sent"`
	BytesRecv   uint64  `json:"bytes_recv"`
	PacketsSent uint64  `json:"packets_sent"`
	PacketsRecv uint64  `json:"packets_recv"`
	Latency     float64 `json:"latency_ms"`
	Bandwidth   float64 `json:"bandwidth_mbps"`
}

type LoadMetrics struct {
	Load1  float64 `json:"load_1min"`
	Load5  float64 `json:"load_5min"`
	Load15 float64 `json:"load_15min"`
}

type DBMetrics struct {
	QueryTime     float64 `json:"avg_query_time_ms"`
	SlowQueries   int     `json:"slow_queries"`
	Connections   int     `json:"active_connections"`
	QueriesPerSec float64 `json:"queries_per_sec"`
}

type SecurityMetrics struct {
	UnauthorizedAttempts int              `json:"unauthorized_attempts"`
	SuspiciousTraffic    int              `json:"suspicious_traffic"`
	FirewallBlocks       int              `json:"firewall_blocks"`
	RecentThreats        []SecurityThreat `json:"recent_threats"`
}

type SecurityThreat struct {
	IP        string    `json:"ip"`
	Type      string    `json:"type"`
	Severity  string    `json:"severity"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}

type UptimeMetrics struct {
	IsUp         bool    `json:"is_up"`
	ResponseTime float64 `json:"response_time_ms"`
	StatusCode   int     `json:"status_code"`
	Location     string  `json:"location"`
}

// Request tracking for metrics
type RequestMetrics struct {
	StartTime    time.Time
	EndTime      time.Time
	StatusCode   int
	ResponseTime float64
	Path         string
	Method       string
	IP           string
	UserAgent    string
}

// Auth structures
type AuthRequest struct {
	APIKey string `json:"api_key" binding:"required"`
}

type AuthResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Simple request structures for easier use
type SimpleOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" example:"087786388052"`
}

type SimpleVerifyOTPRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" example:"087786388052"`
	OTPCode     string `json:"otp_code" binding:"required" example:"123456"`
}

type SimplePurchaseRequest struct {
	PhoneNumber   string `json:"phone_number" binding:"required" example:"087786388052"`
	PackageCode   string `json:"package_code" binding:"required" example:"XL_MASTIF_30D_P_V1"`
	PaymentMethod string `json:"payment_method" binding:"required" example:"BALANCE"`
	AccessToken   string `json:"access_token" binding:"required" example:"1097690:ec321a89-6593-4b01-9a9e-383744301283"`
	Source        string `json:"source,omitempty" example:"telegram_bot"`
}

// Monitoring and Tracking Structures
type TransactionRecord struct {
	ID            string     `json:"id"`
	PhoneNumber   string     `json:"phone_number"`
	PackageCode   string     `json:"package_code"`
	PackageName   string     `json:"package_name"`
	PaymentMethod string     `json:"payment_method"`
	Source        string     `json:"source"`
	Status        string     `json:"status"`
	Amount        int        `json:"amount"`
	ProcessingFee int        `json:"processing_fee"`
	TrxID         string     `json:"trx_id"`
	CreatedAt     time.Time  `json:"created_at"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
	ErrorMessage  string     `json:"error_message,omitempty"`
}

type SystemStats struct {
	TotalTransactions      int       `json:"total_transactions"`
	SuccessfulTransactions int       `json:"successful_transactions"`
	FailedTransactions     int       `json:"failed_transactions"`
	TotalRevenue           int       `json:"total_revenue"`
	SuccessRate            float64   `json:"success_rate"`
	LastUpdated            time.Time `json:"last_updated"`
}

type SourceStats struct {
	Source      string  `json:"source"`
	Count       int     `json:"count"`
	Revenue     int     `json:"revenue"`
	SuccessRate float64 `json:"success_rate"`
}

// Invoice structures
type InvoiceRecord struct {
	ID              string                 `json:"id"`
	InvoiceID       string                 `json:"invoice_id"`
	Title           string                 `json:"title"`
	Amount          int                    `json:"amount"`
	InvoiceStatusID string                 `json:"invoice_status_id"`
	CreatedAt       string                 `json:"created_at"`
	CreatedAtMonth  string                 `json:"created_at_month"`
	CreatedAtYear   int                    `json:"created_at_year"`
	ModifiedAt      interface{}            `json:"modified_at"`
	PaidAt          interface{}            `json:"paid_at"`
	DueDate         string                 `json:"due_date"`
	Username        string                 `json:"username"`
	Fullname        string                 `json:"fullname"`
	Email           string                 `json:"email"`
	Phone           string                 `json:"phone"`
	Metadata        map[string]interface{} `json:"metadata"`
	Payments        []PaymentRecord        `json:"payments"`
}

type PaymentRecord struct {
	ID               string      `json:"id"`
	Amount           int         `json:"amount"`
	PaymentStatusID  string      `json:"payment_status_id"`
	PaymentMethodID  string      `json:"payment_method_id"`
	ReferenceType    string      `json:"reference_type"`
	ReferenceID      string      `json:"reference_id"`
	ExternalID       string      `json:"external_id"`
	PublisherOrderID interface{} `json:"publisher_order_id"`
	PaymentLink      string      `json:"payment_link"`
	CreatedAt        string      `json:"created_at"`
	CompletedAt      interface{} `json:"completed_at"`
	ExpiredAt        string      `json:"expired_at"`
}

type InvoiceStats struct {
	TotalInvoices   int       `json:"total_invoices"`
	PaidInvoices    int       `json:"paid_invoices"`
	UnpaidInvoices  int       `json:"unpaid_invoices"`
	PendingInvoices int       `json:"pending_invoices"`
	TotalRevenue    int       `json:"total_revenue"`
	PaymentRate     float64   `json:"payment_rate"`
	LastUpdated     time.Time `json:"last_updated"`
}

type DashboardData struct {
	Stats              SystemStats         `json:"stats"`
	InvoiceStats       InvoiceStats        `json:"invoice_stats"`
	RecentTransactions []TransactionRecord `json:"recent_transactions"`
	RecentInvoices     []InvoiceRecord     `json:"recent_invoices"`
	SourceBreakdown    []SourceStats       `json:"source_breakdown"`
	Balance            interface{}         `json:"balance"`
	SystemStatus       string              `json:"system_status"`
	MonitoringData     interface{}         `json:"monitoring_data"`
}

type SimpleCheckStatusRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" example:"087786388052"`
	OTPCode     string `json:"otp_code" binding:"required" example:"123456"`
}

type SimpleTransactionCheckRequest struct {
	TransactionID string `json:"transaction_id" binding:"required" example:"a1e3b046-3dda-4eb5-93c2-f0eb01a1a453"`
}

type PackageSearchRequest struct {
	Query         string `json:"query,omitempty" example:"masa aktif"`
	PaymentMethod string `json:"payment_method,omitempty" example:"BALANCE"`
	MaxPrice      int    `json:"max_price,omitempty" example:"10000"`
	MinPrice      int    `json:"min_price,omitempty" example:"1000"`
}

// OTPSession management
type OTPSession struct {
	PhoneNumber string    `json:"phone_number"`
	AuthID      string    `json:"auth_id"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}
