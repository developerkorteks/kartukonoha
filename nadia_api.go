// Package main provides Nadia Package Purchase API
// @title Nadia Package Purchase API
// @version 2.0
// @description API untuk pembelian paket XL dengan flow yang disederhanakan dan user-friendly
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/nabilulilalbab/nadia/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Configuration constants
const (
	ssoApiBaseURL = "https://sso.putri-veronica.my.id/api"
	nadiaApiURL   = "https://putri-veronica.my.id/api/user"
	ssoStaticKey  = "18e0cf008b474614a0b20f9c170cd33f"
	username      = "kortkeks"
	password      = "nabilalbab78"

	// Auth constants
	adminAPIKey = "nadia-admin-2024-secure-key"
	jwtSecret   = "nadia-jwt-secret-key-2024"
)

// API Response structure
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

// OTP session management
type OTPSession struct {
	PhoneNumber string    `json:"phone_number"`
	AuthID      string    `json:"auth_id"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

var (
	otpSessions  = make(map[string]*OTPSession)
	sessionMutex sync.RWMutex

	// Database connection
	db *sql.DB

	// Transaction tracking (keep in-memory for fast access, sync with DB)
	transactions     = make(map[string]*TransactionRecord)
	transactionMutex = sync.RWMutex{}
)

// Token manager (reuse from existing code)
type TokenManager struct {
	token     string
	expiresAt time.Time
	mutex     sync.RWMutex
}

func (tm *TokenManager) getValidToken() (string, error) {
	tm.mutex.RLock()
	if tm.token != "" && time.Now().Before(tm.expiresAt) {
		token := tm.token
		tm.mutex.RUnlock()
		return token, nil
	}
	tm.mutex.RUnlock()

	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	if tm.token != "" && time.Now().Before(tm.expiresAt) {
		return tm.token, nil
	}

	newToken, err := tm.refreshToken()
	if err != nil {
		return "", err
	}

	tm.token = newToken
	tm.expiresAt = time.Now().Add(8 * time.Hour)
	return newToken, nil
}

func (tm *TokenManager) refreshToken() (string, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", fmt.Errorf("failed to create cookie jar: %v", err)
	}
	httpClient := &http.Client{Jar: jar}

	sessionToken, err := getSessionToken(httpClient)
	if err != nil {
		return "", fmt.Errorf("failed to get session token: %v", err)
	}

	finalJWT, err := performApiLogin(httpClient, sessionToken)
	if err != nil {
		return "", fmt.Errorf("failed to perform login: %v", err)
	}

	return finalJWT, nil
}

func getSessionToken(client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", ssoApiBaseURL+"/oauth/request_token", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("access-key", ssoStaticKey)
	req.Header.Set("Referer", "https://sso.putri-veronica.my.id/")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed with status %d: %s", resp.StatusCode, string(body))
	}

	var sessionResp struct {
		Success     bool   `json:"success"`
		AccessToken string `json:"access_token"`
		Message     string `json:"message"`
	}

	if err := json.Unmarshal(body, &sessionResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !sessionResp.Success || sessionResp.AccessToken == "" {
		return "", fmt.Errorf("invalid session token: %s", sessionResp.Message)
	}

	return sessionResp.AccessToken, nil
}

func performApiLogin(client *http.Client, sessionToken string) (string, error) {
	payload := map[string]string{
		"username": username,
		"password": password,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", ssoApiBaseURL+"/user/login", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access-token", sessionToken)
	req.Header.Set("Referer", "https://sso.putri-veronica.my.id/")
	req.Header.Set("Origin", "https://sso.putri-veronica.my.id")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("login failed with status %d: %s", resp.StatusCode, string(body))
	}

	var finalResp struct {
		Success bool   `json:"success"`
		Token   string `json:"token"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(body, &finalResp); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	if !finalResp.Success || finalResp.Token == "" {
		return "", fmt.Errorf("invalid JWT token: %s", finalResp.Message)
	}

	return finalResp.Token, nil
}

func makeNadiaRequest(method, endpoint string, payload interface{}) (*http.Response, error) {
	return makeNadiaRequestWithRetry(method, endpoint, payload, 0)
}

func makeNadiaRequestWithRetry(method, endpoint string, payload interface{}, retryCount int) (*http.Response, error) {
	token, err := tokenManager.getValidToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get valid token: %v", err)
	}

	var reqBody io.Reader
	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %v", err)
		}
		reqBody = bytes.NewBuffer(payloadBytes)
	}

	req, err := http.NewRequest(method, nadiaApiURL+endpoint, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-token", token)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Referer", "https://putri-veronica.my.id/limited/tembak-paket-xl")
	req.Header.Set("Origin", "https://putri-veronica.my.id")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)

	// Handle unauthorized response with JWT refresh
	if err == nil && resp.StatusCode == 401 && retryCount < 2 {
		resp.Body.Close()
		log.Printf("Received 401 Unauthorized, attempting JWT refresh (retry %d/2)", retryCount+1)

		// Force token refresh
		tokenManager.mutex.Lock()
		tokenManager.token = ""
		tokenManager.expiresAt = time.Time{}
		tokenManager.mutex.Unlock()

		// Retry the request
		return makeNadiaRequestWithRetry(method, endpoint, payload, retryCount+1)
	}

	return resp, err
}

// Utility functions
func normalizePhoneNumber(phone string) string {
	// Remove spaces and special characters
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Handle different formats
	if strings.HasPrefix(phone, "+62") {
		phone = phone[3:]
	} else if strings.HasPrefix(phone, "62") {
		phone = phone[2:]
	}

	// Remove leading zero
	if strings.HasPrefix(phone, "0") {
		phone = phone[1:]
	}

	return phone
}

func formatPhoneNumber(phone string, withCountryCode bool) string {
	normalized := normalizePhoneNumber(phone)
	if withCountryCode {
		return "62" + normalized
	}
	return normalized
}

// Global token manager
var tokenManager *TokenManager

func init() {
	tokenManager = &TokenManager{}
}

// Transaction tracking functions
func generateTransactionID() string {
	return fmt.Sprintf("TXN_%d", time.Now().UnixNano())
}

func recordTransaction(phoneNumber, packageCode, packageName, paymentMethod, source string) *TransactionRecord {
	transactionMutex.Lock()
	defer transactionMutex.Unlock()

	record := &TransactionRecord{
		ID:            generateTransactionID(),
		PhoneNumber:   phoneNumber,
		PackageCode:   packageCode,
		PackageName:   packageName,
		PaymentMethod: paymentMethod,
		Source:        source,
		Status:        "PENDING",
		CreatedAt:     time.Now(),
	}

	transactions[record.ID] = record

	// Save to database
	if err := saveTransactionToDB(record); err != nil {
		log.Printf("Failed to save transaction to database: %v", err)
	}

	return record
}

func updateTransactionStatus(id, status, trxID string, amount, processingFee int, errorMsg string) {
	transactionMutex.Lock()
	defer transactionMutex.Unlock()

	if record, exists := transactions[id]; exists {
		record.Status = status
		record.TrxID = trxID
		record.Amount = amount
		record.ProcessingFee = processingFee
		record.ErrorMessage = errorMsg
		now := time.Now()
		record.CompletedAt = &now

		// Save to database
		if err := saveTransactionToDB(record); err != nil {
			log.Printf("Failed to update transaction in database: %v", err)
		}
	}
}

func getSystemStats() SystemStats {
	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	var total, successful, failed, revenue int

	for _, tx := range transactions {
		total++
		if tx.Status == "SUCCESS" {
			successful++
			revenue += tx.Amount
		} else if tx.Status == "FAILED" {
			failed++
		}
	}

	successRate := 0.0
	if total > 0 {
		successRate = float64(successful) / float64(total) * 100
	}

	return SystemStats{
		TotalTransactions:      total,
		SuccessfulTransactions: successful,
		FailedTransactions:     failed,
		TotalRevenue:           revenue,
		SuccessRate:            successRate,
		LastUpdated:            time.Now(),
	}
}

func getSourceStats() []SourceStats {
	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	sourceMap := make(map[string]*SourceStats)

	for _, tx := range transactions {
		source := tx.Source
		if source == "" {
			source = "unknown"
		}

		if _, exists := sourceMap[source]; !exists {
			sourceMap[source] = &SourceStats{Source: source}
		}

		sourceMap[source].Count++
		if tx.Status == "SUCCESS" {
			sourceMap[source].Revenue += tx.Amount
		}
	}

	var result []SourceStats
	for _, stats := range sourceMap {
		if stats.Count > 0 {
			// Calculate success rate for this source
			successCount := 0
			for _, tx := range transactions {
				txSource := tx.Source
				if txSource == "" {
					txSource = "unknown"
				}
				if txSource == stats.Source && tx.Status == "SUCCESS" {
					successCount++
				}
			}
			stats.SuccessRate = float64(successCount) / float64(stats.Count) * 100
		}
		result = append(result, *stats)
	}

	return result
}

func getRecentTransactions(limit int) []TransactionRecord {
	// Get transactions from database (already sorted by created_at DESC)
	txList, err := getAllTransactionsFromDB(limit)
	if err != nil {
		log.Printf("Failed to get transactions from DB: %v", err)
		// Fallback to in-memory data
		transactionMutex.RLock()
		defer transactionMutex.RUnlock()

		var result []TransactionRecord
		for _, tx := range transactions {
			result = append(result, *tx)
		}

		// Sort by created time (newest first)
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].CreatedAt.Before(result[j].CreatedAt) {
					result[i], result[j] = result[j], result[i]
				}
			}
		}

		if len(result) > limit {
			result = result[:limit]
		}

		return result
	}

	// Convert to []TransactionRecord
	var result []TransactionRecord
	for _, tx := range txList {
		result = append(result, *tx)
	}

	return result
}

// API Handlers

// GetAllPackages godoc
// @Summary Get all available packages
// @Description Retrieve all available packages from Nadia API without filtering
// @Tags packages
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of packages returned" default(100)
// @Success 200 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/packages [get]
func getAllPackages(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	limit, _ := strconv.Atoi(limitStr)

	resp, err := makeNadiaRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to fetch packages: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp struct {
		StatusCode int       `json:"statusCode"`
		Message    string    `json:"message"`
		Success    bool      `json:"success"`
		Data       []Package `json:"data"`
	}

	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Apply limit if specified
	packages := nadiaResp.Data
	if limit > 0 && len(packages) > limit {
		packages = packages[:limit]
	}

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    fmt.Sprintf("Retrieved %d packages", len(packages)),
		Success:    true,
		Data:       packages,
	})
}

// SearchPackages godoc
// @Summary Search packages with filters
// @Description Search and filter packages by name, price, payment method, etc.
// @Tags packages
// @Accept json
// @Produce json
// @Param search body PackageSearchRequest true "Search filters"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/packages/search [post]
func searchPackages(c *gin.Context) {
	var searchReq PackageSearchRequest
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid search request: " + err.Error(),
			Success:    false,
		})
		return
	}

	resp, err := makeNadiaRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to fetch packages: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp struct {
		StatusCode int       `json:"statusCode"`
		Message    string    `json:"message"`
		Success    bool      `json:"success"`
		Data       []Package `json:"data"`
	}

	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Apply filters
	var filteredPackages []Package
	for _, pkg := range nadiaResp.Data {
		// Text search
		if searchReq.Query != "" {
			query := strings.ToLower(searchReq.Query)
			if !strings.Contains(strings.ToLower(pkg.PackageName), query) &&
				!strings.Contains(strings.ToLower(pkg.PackageDescription), query) {
				continue
			}
		}

		// Payment method filter
		if searchReq.PaymentMethod != "" {
			hasPaymentMethod := false
			for _, pm := range pkg.AvailablePaymentMethods {
				if pm.PaymentMethod == searchReq.PaymentMethod {
					hasPaymentMethod = true
					break
				}
			}
			if !hasPaymentMethod {
				continue
			}
		}

		// Price filters
		if searchReq.MinPrice > 0 && pkg.PackagePrice < searchReq.MinPrice {
			continue
		}
		if searchReq.MaxPrice > 0 && pkg.PackagePrice > searchReq.MaxPrice {
			continue
		}

		filteredPackages = append(filteredPackages, pkg)
	}

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    fmt.Sprintf("Found %d packages matching your criteria", len(filteredPackages)),
		Success:    true,
		Data:       filteredPackages,
	})
}

// RequestOTP godoc
// @Summary Request OTP for phone number
// @Description Request OTP code to be sent to the specified phone number. The OTP will be stored in session for 5 minutes.
// @Tags otp
// @Accept json
// @Produce json
// @Param request body SimpleOTPRequest true "OTP request"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/otp/request [post]
func requestOTP(c *gin.Context) {
	var req SimpleOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid request: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Use phone number as-is for OTP request, since normalization might cause issues
	phoneForOTP := req.PhoneNumber

	// Only normalize for session storage
	normalizedPhone := normalizePhoneNumber(req.PhoneNumber)

	payload := map[string]string{
		"phone": phoneForOTP,
	}

	resp, err := makeNadiaRequest("POST", "/limited/xl/request-otp.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to request OTP: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// First try to parse as object
	var nadiaResp struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Success    bool        `json:"success"`
		Data       interface{} `json:"data"`
	}

	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Extract auth_id and can_resend_in from data
	var authID string
	var canResendIn int

	if nadiaResp.Success && nadiaResp.Data != nil {
		// Try to parse data as object first
		if dataMap, ok := nadiaResp.Data.(map[string]interface{}); ok {
			if authIDVal, exists := dataMap["auth_id"]; exists {
				if authIDStr, ok := authIDVal.(string); ok {
					authID = authIDStr
				}
			}
			if canResendVal, exists := dataMap["can_resend_in"]; exists {
				if canResendFloat, ok := canResendVal.(float64); ok {
					canResendIn = int(canResendFloat)
				}
			}
		} else if dataArray, ok := nadiaResp.Data.([]interface{}); ok && len(dataArray) > 0 {
			// If data is array, try to get first element
			if firstItem, ok := dataArray[0].(map[string]interface{}); ok {
				if authIDVal, exists := firstItem["auth_id"]; exists {
					if authIDStr, ok := authIDVal.(string); ok {
						authID = authIDStr
					}
				}
				if canResendVal, exists := firstItem["can_resend_in"]; exists {
					if canResendFloat, ok := canResendVal.(float64); ok {
						canResendIn = int(canResendFloat)
					}
				}
			}
		}

		// Store OTP session if we have auth_id
		if authID != "" {
			sessionMutex.Lock()
			otpSessions[normalizedPhone] = &OTPSession{
				PhoneNumber: normalizedPhone,
				AuthID:      authID,
				CreatedAt:   time.Now(),
				ExpiresAt:   time.Now().Add(5 * time.Minute),
			}
			sessionMutex.Unlock()
		}
	}

	c.JSON(resp.StatusCode, APIResponse{
		StatusCode: nadiaResp.StatusCode,
		Message:    nadiaResp.Message,
		Success:    nadiaResp.Success,
		Data: map[string]interface{}{
			"phone_number":  req.PhoneNumber,
			"can_resend_in": canResendIn,
			"expires_in":    300, // 5 minutes
		},
	})
}

// VerifyOTP godoc
// @Summary Verify OTP and get access token
// @Description Verify OTP code and get access token for purchasing packages
// @Tags otp
// @Accept json
// @Produce json
// @Param request body SimpleVerifyOTPRequest true "Verify OTP request"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/otp/verify [post]
func verifyOTP(c *gin.Context) {
	var req SimpleVerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid request: " + err.Error(),
			Success:    false,
		})
		return
	}

	normalizedPhone := normalizePhoneNumber(req.PhoneNumber)

	// Get OTP session
	sessionMutex.RLock()
	session, exists := otpSessions[normalizedPhone]
	sessionMutex.RUnlock()

	if !exists {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "OTP session not found. Please request OTP first.",
			Success:    false,
		})
		return
	}

	if time.Now().After(session.ExpiresAt) {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "OTP session expired. Please request a new OTP.",
			Success:    false,
		})
		return
	}

	// Use phone number as-is for verify request
	phoneForVerify := req.PhoneNumber

	payload := map[string]string{
		"phone":   phoneForVerify,
		"auth_id": session.AuthID,
		"otp":     req.OTPCode,
	}

	resp, err := makeNadiaRequest("POST", "/limited/xl/request-login.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to verify OTP: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Parse response
	var nadiaResp struct {
		StatusCode int         `json:"statusCode"`
		Message    string      `json:"message"`
		Success    bool        `json:"success"`
		Data       interface{} `json:"data"`
	}

	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Extract access_token
	var accessToken string
	if nadiaResp.Success && nadiaResp.Data != nil {
		if dataMap, ok := nadiaResp.Data.(map[string]interface{}); ok {
			if tokenVal, exists := dataMap["access_token"]; exists {
				if tokenStr, ok := tokenVal.(string); ok {
					accessToken = tokenStr
				}
			}
		}
	}

	c.JSON(resp.StatusCode, APIResponse{
		StatusCode: nadiaResp.StatusCode,
		Message:    nadiaResp.Message,
		Success:    nadiaResp.Success,
		Data: map[string]interface{}{
			"access_token": accessToken,
			"phone_number": req.PhoneNumber,
		},
	})
}

// PurchasePackage godoc
// @Summary Purchase a package with access token
// @Description Purchase a package using phone number, package code, payment method, and access token
// @Tags purchase
// @Accept json
// @Produce json
// @Param request body SimplePurchaseRequest true "Purchase request"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/purchase [post]
func purchasePackage(c *gin.Context) {
	var req SimplePurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid request: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Use phone number with country code format for purchase
	phoneWithCountryCode := formatPhoneNumber(req.PhoneNumber, true)
	normalizedPhone := normalizePhoneNumber(req.PhoneNumber)

	// Record transaction start
	source := req.Source
	if source == "" {
		source = "api_direct"
	}

	txRecord := recordTransaction(req.PhoneNumber, req.PackageCode, "", req.PaymentMethod, source)

	payload := map[string]interface{}{
		"package_code":   req.PackageCode,
		"phone":          phoneWithCountryCode,
		"access_token":   req.AccessToken,
		"payment_method": req.PaymentMethod,
	}

	resp, err := makeNadiaRequest("POST", "/limited/xl/beli-paket-otp.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to purchase package: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		updateTransactionStatus(txRecord.ID, "FAILED", "", 0, 0, "Failed to parse response: "+err.Error())
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Update transaction status and extract details
	var trxID string
	var amount, processingFee int
	var packageName string
	var errorMsg string

	if nadiaResp.Success {
		// Extract transaction details from response data
		if dataMap, ok := nadiaResp.Data.(map[string]interface{}); ok {
			if trxIDVal, exists := dataMap["trx_id"]; exists {
				if trxIDStr, ok := trxIDVal.(string); ok {
					trxID = trxIDStr
				}
			}
			if packageNameVal, exists := dataMap["package_name"]; exists {
				if packageNameStr, ok := packageNameVal.(string); ok {
					packageName = packageNameStr
				}
			}
			if processingFeeVal, exists := dataMap["package_processing_fee"]; exists {
				if processingFeeFloat, ok := processingFeeVal.(float64); ok {
					processingFee = int(processingFeeFloat)
				}
			}
		}

		// Get package price from package data
		amount = getPackagePrice(req.PackageCode)

		updateTransactionStatus(txRecord.ID, "SUCCESS", trxID, amount, processingFee, "")

		// Clean up OTP session after successful purchase
		sessionMutex.Lock()
		delete(otpSessions, normalizedPhone)
		sessionMutex.Unlock()
	} else {
		errorMsg = nadiaResp.Message
		updateTransactionStatus(txRecord.ID, "FAILED", "", 0, 0, errorMsg)
	}

	// Update package name in transaction record if we got it
	if packageName != "" {
		transactionMutex.Lock()
		if record, exists := transactions[txRecord.ID]; exists {
			record.PackageName = packageName
		}
		transactionMutex.Unlock()
	}

	c.JSON(resp.StatusCode, nadiaResp)
}

// CheckCardStatus godoc
// @Summary Check XL card status and balance
// @Description Check the status, balance, and active period of an XL card using OTP verification
// @Tags card
// @Accept json
// @Produce json
// @Param request body SimpleCheckStatusRequest true "Check status request"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/card/status [post]
func checkCardStatus(c *gin.Context) {
	var req SimpleCheckStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid request: " + err.Error(),
			Success:    false,
		})
		return
	}

	normalizedPhone := normalizePhoneNumber(req.PhoneNumber)

	// Get OTP session
	sessionMutex.RLock()
	session, exists := otpSessions[normalizedPhone]
	sessionMutex.RUnlock()

	if !exists {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "OTP session not found. Please request OTP first.",
			Success:    false,
		})
		return
	}

	if time.Now().After(session.ExpiresAt) {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "OTP session expired. Please request a new OTP.",
			Success:    false,
		})
		return
	}

	// Create access token from OTP code and auth_id
	accessToken := req.OTPCode + ":" + session.AuthID

	payload := map[string]string{
		"access_token": accessToken,
	}

	resp, err := makeNadiaRequest("POST", "/limited/xl/status-kartu.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to check card status: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(resp.StatusCode, nadiaResp)
}

// CheckActivePackages godoc
// @Summary Check active packages on XL card
// @Description Check all active packages/quotas on an XL card using OTP verification
// @Tags card
// @Accept json
// @Produce json
// @Param request body SimpleCheckStatusRequest true "Check active packages request"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/card/packages [post]
func checkActivePackages(c *gin.Context) {
	var req SimpleCheckStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid request: " + err.Error(),
			Success:    false,
		})
		return
	}

	normalizedPhone := normalizePhoneNumber(req.PhoneNumber)

	// Get OTP session
	sessionMutex.RLock()
	session, exists := otpSessions[normalizedPhone]
	sessionMutex.RUnlock()

	if !exists {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "OTP session not found. Please request OTP first.",
			Success:    false,
		})
		return
	}

	if time.Now().After(session.ExpiresAt) {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "OTP session expired. Please request a new OTP.",
			Success:    false,
		})
		return
	}

	// Create access token from OTP code and auth_id
	accessToken := req.OTPCode + ":" + session.AuthID

	payload := map[string]string{
		"access_token": accessToken,
	}

	resp, err := makeNadiaRequest("POST", "/limited/xl/package-active-list.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to check active packages: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetBalance godoc
// @Summary Get account balance
// @Description Retrieve current account balance
// @Tags wallet
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/balance [get]
func getBalance(c *gin.Context) {
	resp, err := makeNadiaRequest("GET", "/wallet/balance.json", nil)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to get balance: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(200, nadiaResp)
}

// CheckTransaction godoc
// @Summary Check transaction status
// @Description Check the status of a transaction by transaction ID
// @Tags transaction
// @Accept json
// @Produce json
// @Param request body SimpleTransactionCheckRequest true "Transaction check request"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/transaction/check [post]
func checkTransaction(c *gin.Context) {
	var req SimpleTransactionCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, APIResponse{
			StatusCode: 400,
			Message:    "Invalid request: " + err.Error(),
			Success:    false,
		})
		return
	}

	payload := map[string]string{
		"transaction_id": req.TransactionID,
	}

	resp, err := makeNadiaRequest("POST", "/limited/xl/check-transaction.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to check transaction: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetPaymentMethods godoc
// @Summary Get available payment methods
// @Description Retrieve all available payment methods
// @Tags payment
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/payment-methods [get]
func getPaymentMethods(c *gin.Context) {
	resp, err := makeNadiaRequest("GET", "/wallet/payment-methods.json", nil)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to get payment methods: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(200, nadiaResp)
}

// GetPackageStock godoc
// @Summary Get package stock information
// @Description Retrieve stock information for all packages
// @Tags packages
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/packages/stock [get]
func getPackageStock(c *gin.Context) {
	resp, err := makeNadiaRequest("GET", "/limited/xl/check-stock-package-global.json", nil)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to get package stock: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(200, nadiaResp)
}

// GetInvoices godoc
// @Summary Get invoice list
// @Description Retrieve list of invoices with pagination
// @Tags invoice
// @Accept json
// @Produce json
// @Param search query string false "Search term"
// @Param limit query int false "Limit per page" default(20)
// @Success 200 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/invoices [get]
func getInvoices(c *gin.Context) {
	search := c.DefaultQuery("search", "")
	limitStr := c.DefaultQuery("limit", "20")
	limit, _ := strconv.Atoi(limitStr)

	payload := map[string]interface{}{
		"search":          search,
		"last_created_at": 0,
		"limit":           limit,
	}

	resp, err := makeNadiaRequest("POST", "/invoice/list.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to get invoices: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var nadiaResp APIResponse
	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response: " + err.Error(),
			Success:    false,
		})
		return
	}

	c.JSON(200, nadiaResp)
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if the API is running properly
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/health [get]
func healthCheck(c *gin.Context) {
	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "API is running successfully",
		Success:    true,
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"version":   "2.0.0",
			"features": []string{
				"Simplified OTP flow",
				"Automatic phone number formatting",
				"Session management",
				"Package search and filtering",
				"Complete XL card management",
			},
		},
	})
}

// Helper function to get package price by package code
func getPackagePrice(packageCode string) int {
	// Make request to get all packages
	resp, err := makeNadiaRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		log.Printf("Failed to fetch packages for price lookup: %v", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read packages response: %v", err)
		return 0
	}

	var nadiaResp struct {
		StatusCode int       `json:"statusCode"`
		Message    string    `json:"message"`
		Success    bool      `json:"success"`
		Data       []Package `json:"data"`
	}

	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		log.Printf("Failed to parse packages response: %v", err)
		return 0
	}

	if !nadiaResp.Success {
		log.Printf("Failed to get packages: %s", nadiaResp.Message)
		return 0
	}

	// Find package by code
	for _, pkg := range nadiaResp.Data {
		if pkg.PackageCode == packageCode {
			return pkg.PackagePrice
		}
	}

	log.Printf("Package not found: %s", packageCode)
	return 0
}

// Database functions
func initDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./nadia_transactions.db")
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	// Create transactions table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS transactions (
		id TEXT PRIMARY KEY,
		phone_number TEXT NOT NULL,
		package_code TEXT NOT NULL,
		package_name TEXT,
		payment_method TEXT NOT NULL,
		source TEXT,
		status TEXT NOT NULL,
		amount INTEGER DEFAULT 0,
		processing_fee INTEGER DEFAULT 0,
		trx_id TEXT,
		created_at DATETIME NOT NULL,
		completed_at DATETIME,
		error_message TEXT
	);
	
	CREATE INDEX IF NOT EXISTS idx_phone_number ON transactions(phone_number);
	CREATE INDEX IF NOT EXISTS idx_status ON transactions(status);
	CREATE INDEX IF NOT EXISTS idx_source ON transactions(source);
	CREATE INDEX IF NOT EXISTS idx_created_at ON transactions(created_at);
	`

	if _, err := db.Exec(createTableSQL); err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("Database initialized successfully")
	return nil
}

func saveTransactionToDB(tx *TransactionRecord) error {
	query := `
	INSERT OR REPLACE INTO transactions 
	(id, phone_number, package_code, package_name, payment_method, source, status, 
	 amount, processing_fee, trx_id, created_at, completed_at, error_message)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	var completedAt *time.Time
	if tx.CompletedAt != nil {
		completedAt = tx.CompletedAt
	}

	_, err := db.Exec(query,
		tx.ID, tx.PhoneNumber, tx.PackageCode, tx.PackageName, tx.PaymentMethod,
		tx.Source, tx.Status, tx.Amount, tx.ProcessingFee, tx.TrxID,
		tx.CreatedAt, completedAt, tx.ErrorMessage)

	if err != nil {
		log.Printf("Failed to save transaction to DB: %v", err)
		return err
	}

	return nil
}

func loadTransactionsFromDB() error {
	query := `
	SELECT id, phone_number, package_code, package_name, payment_method, source, 
	       status, amount, processing_fee, trx_id, created_at, completed_at, error_message
	FROM transactions
	ORDER BY created_at DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("failed to query transactions: %v", err)
	}
	defer rows.Close()

	transactionMutex.Lock()
	defer transactionMutex.Unlock()

	count := 0
	for rows.Next() {
		tx := &TransactionRecord{}
		var completedAt sql.NullTime

		err := rows.Scan(
			&tx.ID, &tx.PhoneNumber, &tx.PackageCode, &tx.PackageName,
			&tx.PaymentMethod, &tx.Source, &tx.Status, &tx.Amount,
			&tx.ProcessingFee, &tx.TrxID, &tx.CreatedAt, &completedAt,
			&tx.ErrorMessage)

		if err != nil {
			log.Printf("Failed to scan transaction: %v", err)
			continue
		}

		if completedAt.Valid {
			tx.CompletedAt = &completedAt.Time
		}

		transactions[tx.ID] = tx
		count++
	}

	log.Printf("Loaded %d transactions from database", count)
	return nil
}

func getAllTransactionsFromDB(limit int) ([]*TransactionRecord, error) {
	query := `
	SELECT id, phone_number, package_code, package_name, payment_method, source, 
	       status, amount, processing_fee, trx_id, created_at, completed_at, error_message
	FROM transactions
	ORDER BY created_at DESC
	`

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %v", err)
	}
	defer rows.Close()

	var result []*TransactionRecord
	for rows.Next() {
		tx := &TransactionRecord{}
		var completedAt sql.NullTime

		err := rows.Scan(
			&tx.ID, &tx.PhoneNumber, &tx.PackageCode, &tx.PackageName,
			&tx.PaymentMethod, &tx.Source, &tx.Status, &tx.Amount,
			&tx.ProcessingFee, &tx.TrxID, &tx.CreatedAt, &completedAt,
			&tx.ErrorMessage)

		if err != nil {
			log.Printf("Failed to scan transaction: %v", err)
			continue
		}

		if completedAt.Valid {
			tx.CompletedAt = &completedAt.Time
		}

		result = append(result, tx)
	}

	return result, nil
}

func main() {
	// Initialize database
	if err := initDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Load existing transactions from database
	if err := loadTransactionsFromDB(); err != nil {
		log.Printf("Warning: Failed to load transactions from database: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Add CORS middleware
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API Routes
	api := r.Group("/api")
	{
		// System
		api.GET("/health", healthCheck)

		// Packages
		api.GET("/packages", getAllPackages)
		api.POST("/packages/search", searchPackages)
		api.GET("/packages/stock", getPackageStock)

		// OTP
		api.POST("/otp/request", requestOTP)
		api.POST("/otp/verify", verifyOTP)

		// Purchase
		api.POST("/purchase", purchasePackage)

		// Card Management
		api.POST("/card/status", checkCardStatus)
		api.POST("/card/packages", checkActivePackages)

		// Wallet
		api.GET("/balance", getBalance)
		api.GET("/payment-methods", getPaymentMethods)

		// Transaction
		api.POST("/transaction/check", checkTransaction)

		// Invoice
		api.GET("/invoices", getInvoices)
		api.GET("/invoices/:id", getInvoiceDetail)
		api.GET("/invoice/stats", getInvoiceStatsHandler)

		// Dashboard & Monitoring
		api.GET("/dashboard", getDashboard)
		api.GET("/transactions", getTransactions)
		api.GET("/transactions/:id", getTransactionDetail)
		api.GET("/stats", getStats)
		api.GET("/stats/hourly", getHourlyStats)
		api.GET("/stats/daily", getDailyStats)
		api.GET("/export/transactions", exportTransactions)
		api.GET("/export/invoices", exportInvoices)
	}

	// Serve Dashboard
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/dashboard")
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("dashboard_enhanced.html")
	})

	r.GET("/dashboard/old", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("dashboard.html")
	})

	r.GET("/api-flow", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("api_flow_example.html")
	})

	// Serve Swagger documentation
	r.GET("/swagger", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Starting Nadia API Server v2.0 on :8080")
	log.Println("📊 Dashboard available at: http://localhost:8080/dashboard")
	log.Println("📚 Swagger Documentation available at: http://localhost:8080/swagger/index.html")
	log.Println("")
	log.Println("🚀 New Features:")
	log.Println("  ✅ Enhanced Dashboard with comprehensive monitoring & analytics")
	log.Println("  ✅ Invoice global monitoring (not by source)")
	log.Println("  ✅ Transaction tracking with source identification")
	log.Println("  ✅ Invoice statistics and payment rate tracking")
	log.Println("  ✅ CSV export for both transactions and invoices")
	log.Println("  ✅ Advanced filtering and search capabilities")
	log.Println("  ✅ Real-time charts and analytics")
	log.Println("  ✅ Automatic JWT refresh on unauthorized errors")
	log.Println("  ✅ Simplified OTP flow - no manual auth_id management")
	log.Println("  ✅ Automatic phone number formatting")
	log.Println("  ✅ Session management with auto-cleanup")
	log.Println("  ✅ Package search and filtering")
	log.Println("  ✅ Complete XL card management")
	log.Println("  ✅ User-friendly request/response format")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Dashboard & Monitoring Handlers

// GetDashboard godoc
// @Summary Get dashboard data
// @Description Get comprehensive dashboard data including stats, recent transactions, and balance
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 500 {object} APIResponse
// @Router /api/dashboard [get]
func getDashboard(c *gin.Context) {
	// Get balance
	balanceResp, err := makeNadiaRequest("GET", "/wallet/balance.json", nil)
	var balance interface{}
	if err == nil {
		defer balanceResp.Body.Close()
		body, _ := io.ReadAll(balanceResp.Body)
		json.Unmarshal(body, &balance)
	}

	// Determine system status
	systemStatus := "OPERATIONAL"
	if err != nil {
		systemStatus = "DEGRADED"
	}

	dashboardData := DashboardData{
		Stats:              getSystemStats(),
		InvoiceStats:       getInvoiceStats(),
		RecentTransactions: getRecentTransactions(10),
		RecentInvoices:     getRecentInvoices(10),
		SourceBreakdown:    getSourceStats(),
		Balance:            balance,
		SystemStatus:       systemStatus,
	}

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "Dashboard data retrieved successfully",
		Success:    true,
		Data:       dashboardData,
	})
}

// GetTransactions godoc
// @Summary Get transaction history
// @Description Get paginated transaction history with optional filtering
// @Tags dashboard
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of transactions" default(50)
// @Param offset query int false "Offset for pagination" default(0)
// @Param status query string false "Filter by status (SUCCESS, FAILED, PENDING)"
// @Param source query string false "Filter by source (telegram_bot, whatsapp_bot, etc.)"
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/transactions [get]
func getTransactions(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")
	statusFilter := c.Query("status")
	sourceFilter := c.Query("source")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	var filtered []TransactionRecord
	for _, tx := range transactions {
		// Apply filters
		if statusFilter != "" && tx.Status != statusFilter {
			continue
		}
		if sourceFilter != "" && tx.Source != sourceFilter {
			continue
		}
		filtered = append(filtered, *tx)
	}

	// Sort by created time (newest first)
	for i := 0; i < len(filtered)-1; i++ {
		for j := i + 1; j < len(filtered); j++ {
			if filtered[i].CreatedAt.Before(filtered[j].CreatedAt) {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}

	// Apply pagination
	total := len(filtered)
	if offset >= total {
		filtered = []TransactionRecord{}
	} else {
		end := offset + limit
		if end > total {
			end = total
		}
		filtered = filtered[offset:end]
	}

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "Transactions retrieved successfully",
		Success:    true,
		Data: map[string]interface{}{
			"transactions": filtered,
			"total":        total,
			"limit":        limit,
			"offset":       offset,
		},
	})
}

// GetTransactionDetail godoc
// @Summary Get transaction detail
// @Description Get detailed information about a specific transaction
// @Tags dashboard
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/transactions/{id} [get]
func getTransactionDetail(c *gin.Context) {
	id := c.Param("id")

	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	if tx, exists := transactions[id]; exists {
		c.JSON(200, APIResponse{
			StatusCode: 200,
			Message:    "Transaction detail retrieved successfully",
			Success:    true,
			Data:       tx,
		})
	} else {
		c.JSON(404, APIResponse{
			StatusCode: 404,
			Message:    "Transaction not found",
			Success:    false,
		})
	}
}

// GetStats godoc
// @Summary Get system statistics
// @Description Get comprehensive system statistics including transaction counts, revenue, and success rates
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/stats [get]
func getStats(c *gin.Context) {
	stats := getSystemStats()
	sourceStats := getSourceStats()

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "Statistics retrieved successfully",
		Success:    true,
		Data: map[string]interface{}{
			"system_stats": stats,
			"source_stats": sourceStats,
		},
	})
}

// GetHourlyStats godoc
// @Summary Get hourly statistics
// @Description Get transaction statistics grouped by hour for the last 24 hours
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/stats/hourly [get]
func getHourlyStats(c *gin.Context) {
	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	now := time.Now()
	hourlyStats := make(map[int]map[string]int)

	// Initialize 24 hours
	for i := 0; i < 24; i++ {
		hourlyStats[i] = map[string]int{
			"total":      0,
			"successful": 0,
			"failed":     0,
			"revenue":    0,
		}
	}

	// Process transactions from last 24 hours
	for _, tx := range transactions {
		if now.Sub(tx.CreatedAt) <= 24*time.Hour {
			hour := tx.CreatedAt.Hour()
			hourlyStats[hour]["total"]++

			if tx.Status == "SUCCESS" {
				hourlyStats[hour]["successful"]++
				hourlyStats[hour]["revenue"] += tx.Amount
			} else if tx.Status == "FAILED" {
				hourlyStats[hour]["failed"]++
			}
		}
	}

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "Hourly statistics retrieved successfully",
		Success:    true,
		Data:       hourlyStats,
	})
}

// GetDailyStats godoc
// @Summary Get daily statistics
// @Description Get transaction statistics grouped by day for the last 30 days
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/stats/daily [get]
func getDailyStats(c *gin.Context) {
	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	now := time.Now()
	dailyStats := make(map[string]map[string]int)

	// Initialize last 30 days
	for i := 0; i < 30; i++ {
		date := now.AddDate(0, 0, -i).Format("2006-01-02")
		dailyStats[date] = map[string]int{
			"total":      0,
			"successful": 0,
			"failed":     0,
			"revenue":    0,
		}
	}

	// Process transactions from last 30 days
	for _, tx := range transactions {
		if now.Sub(tx.CreatedAt) <= 30*24*time.Hour {
			date := tx.CreatedAt.Format("2006-01-02")
			if stats, exists := dailyStats[date]; exists {
				stats["total"]++

				if tx.Status == "SUCCESS" {
					stats["successful"]++
					stats["revenue"] += tx.Amount
				} else if tx.Status == "FAILED" {
					stats["failed"]++
				}
			}
		}
	}

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "Daily statistics retrieved successfully",
		Success:    true,
		Data:       dailyStats,
	})
}

// ExportTransactions godoc
// @Summary Export transactions to CSV
// @Description Export transaction data to CSV format with optional filtering
// @Tags dashboard
// @Accept json
// @Produce text/csv
// @Param status query string false "Filter by status (SUCCESS, FAILED, PENDING)"
// @Param source query string false "Filter by source (telegram_bot, whatsapp_bot, etc.)"
// @Param from query string false "Start date (YYYY-MM-DD)"
// @Param to query string false "End date (YYYY-MM-DD)"
// @Success 200 {string} string "CSV data"
// @Router /api/export/transactions [get]
func exportTransactions(c *gin.Context) {
	statusFilter := c.Query("status")
	sourceFilter := c.Query("source")
	fromDate := c.Query("from")
	toDate := c.Query("to")

	transactionMutex.RLock()
	defer transactionMutex.RUnlock()

	var filtered []TransactionRecord
	for _, tx := range transactions {
		// Apply filters
		if statusFilter != "" && tx.Status != statusFilter {
			continue
		}
		if sourceFilter != "" && tx.Source != sourceFilter {
			continue
		}

		// Date filters
		if fromDate != "" {
			if from, err := time.Parse("2006-01-02", fromDate); err == nil {
				if tx.CreatedAt.Before(from) {
					continue
				}
			}
		}
		if toDate != "" {
			if to, err := time.Parse("2006-01-02", toDate); err == nil {
				if tx.CreatedAt.After(to.Add(24 * time.Hour)) {
					continue
				}
			}
		}

		filtered = append(filtered, *tx)
	}

	// Sort by created time (newest first)
	for i := 0; i < len(filtered)-1; i++ {
		for j := i + 1; j < len(filtered); j++ {
			if filtered[i].CreatedAt.Before(filtered[j].CreatedAt) {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}

	// Generate CSV
	var csvData strings.Builder
	csvData.WriteString("ID,Phone Number,Package Code,Package Name,Payment Method,Source,Status,Amount,Processing Fee,TrxID,Created At,Completed At,Error Message\n")

	for _, tx := range filtered {
		completedAt := ""
		if tx.CompletedAt != nil {
			completedAt = tx.CompletedAt.Format("2006-01-02 15:04:05")
		}

		csvData.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%d,%d,%s,%s,%s,%s\n",
			tx.ID,
			tx.PhoneNumber,
			tx.PackageCode,
			tx.PackageName,
			tx.PaymentMethod,
			tx.Source,
			tx.Status,
			tx.Amount,
			tx.ProcessingFee,
			tx.TrxID,
			tx.CreatedAt.Format("2006-01-02 15:04:05"),
			completedAt,
			tx.ErrorMessage,
		))
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=transactions_%s.csv", time.Now().Format("20060102_150405")))
	c.String(200, csvData.String())
}

// Invoice helper functions
func getInvoiceStats() InvoiceStats {
	// Get invoice data from API
	payload := map[string]interface{}{
		"search":          "",
		"last_created_at": 0,
		"limit":           1000, // Get more data for stats
	}

	resp, err := makeNadiaRequest("POST", "/invoice/list.json", payload)
	if err != nil {
		return InvoiceStats{
			LastUpdated: time.Now(),
		}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return InvoiceStats{
			LastUpdated: time.Now(),
		}
	}

	var invoiceResp struct {
		StatusCode int             `json:"statusCode"`
		Success    bool            `json:"success"`
		Data       []InvoiceRecord `json:"data"`
	}

	if err := json.Unmarshal(body, &invoiceResp); err != nil || !invoiceResp.Success {
		return InvoiceStats{
			LastUpdated: time.Now(),
		}
	}

	// Calculate stats
	stats := InvoiceStats{
		TotalInvoices: len(invoiceResp.Data),
		LastUpdated:   time.Now(),
	}

	totalRevenue := 0
	for _, invoice := range invoiceResp.Data {
		switch invoice.InvoiceStatusID {
		case "paid":
			stats.PaidInvoices++
			totalRevenue += invoice.Amount
		case "unpaid":
			stats.UnpaidInvoices++
		case "pending":
			stats.PendingInvoices++
		}
	}

	stats.TotalRevenue = totalRevenue
	if stats.TotalInvoices > 0 {
		stats.PaymentRate = float64(stats.PaidInvoices) / float64(stats.TotalInvoices) * 100
	}

	return stats
}

func getRecentInvoices(limit int) []InvoiceRecord {
	payload := map[string]interface{}{
		"search":          "",
		"last_created_at": 0,
		"limit":           limit,
	}

	resp, err := makeNadiaRequest("POST", "/invoice/list.json", payload)
	if err != nil {
		return []InvoiceRecord{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []InvoiceRecord{}
	}

	var invoiceResp struct {
		StatusCode int             `json:"statusCode"`
		Success    bool            `json:"success"`
		Data       []InvoiceRecord `json:"data"`
	}

	if err := json.Unmarshal(body, &invoiceResp); err != nil || !invoiceResp.Success {
		return []InvoiceRecord{}
	}

	return invoiceResp.Data
}

// GetInvoiceStats godoc
// @Summary Get invoice statistics
// @Description Get comprehensive invoice statistics including payment rates and revenue
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/invoice/stats [get]
func getInvoiceStatsHandler(c *gin.Context) {
	stats := getInvoiceStats()
	recentInvoices := getRecentInvoices(20)

	c.JSON(200, APIResponse{
		StatusCode: 200,
		Message:    "Invoice statistics retrieved successfully",
		Success:    true,
		Data: map[string]interface{}{
			"stats":           stats,
			"recent_invoices": recentInvoices,
		},
	})
}

// GetInvoiceDetail godoc
// @Summary Get invoice detail
// @Description Get detailed information about a specific invoice
// @Tags invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/invoices/{id} [get]
func getInvoiceDetail(c *gin.Context) {
	invoiceID := c.Param("id")

	// Get all invoices and find the specific one
	payload := map[string]interface{}{
		"search":          invoiceID,
		"last_created_at": 0,
		"limit":           100,
	}

	resp, err := makeNadiaRequest("POST", "/invoice/list.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to get invoice: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var invoiceResp struct {
		StatusCode int             `json:"statusCode"`
		Success    bool            `json:"success"`
		Data       []InvoiceRecord `json:"data"`
	}

	if err := json.Unmarshal(body, &invoiceResp); err != nil || !invoiceResp.Success {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response",
			Success:    false,
		})
		return
	}

	// Find the specific invoice
	for _, invoice := range invoiceResp.Data {
		if invoice.ID == invoiceID || invoice.InvoiceID == invoiceID {
			c.JSON(200, APIResponse{
				StatusCode: 200,
				Message:    "Invoice detail retrieved successfully",
				Success:    true,
				Data:       invoice,
			})
			return
		}
	}

	c.JSON(404, APIResponse{
		StatusCode: 404,
		Message:    "Invoice not found",
		Success:    false,
	})
}

// ExportInvoices godoc
// @Summary Export invoices to CSV
// @Description Export invoice data to CSV format with optional filtering
// @Tags invoice
// @Accept json
// @Produce text/csv
// @Param status query string false "Filter by status (paid, unpaid, pending)"
// @Param search query string false "Search term"
// @Param from query string false "Start date (YYYY-MM-DD)"
// @Param to query string false "End date (YYYY-MM-DD)"
// @Success 200 {string} string "CSV data"
// @Router /api/export/invoices [get]
func exportInvoices(c *gin.Context) {
	statusFilter := c.Query("status")
	searchFilter := c.Query("search")
	fromDate := c.Query("from")
	toDate := c.Query("to")

	// Get invoice data
	payload := map[string]interface{}{
		"search":          searchFilter,
		"last_created_at": 0,
		"limit":           1000, // Get more data for export
	}

	resp, err := makeNadiaRequest("POST", "/invoice/list.json", payload)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to get invoices: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to read response: " + err.Error(),
			Success:    false,
		})
		return
	}

	var invoiceResp struct {
		StatusCode int             `json:"statusCode"`
		Success    bool            `json:"success"`
		Data       []InvoiceRecord `json:"data"`
	}

	if err := json.Unmarshal(body, &invoiceResp); err != nil || !invoiceResp.Success {
		c.JSON(500, APIResponse{
			StatusCode: 500,
			Message:    "Failed to parse response",
			Success:    false,
		})
		return
	}

	// Filter invoices
	var filtered []InvoiceRecord
	for _, invoice := range invoiceResp.Data {
		// Status filter
		if statusFilter != "" && invoice.InvoiceStatusID != statusFilter {
			continue
		}

		// Date filters
		if fromDate != "" {
			createdTime := time.Unix(0, int64(parseTimestamp(invoice.CreatedAt))*int64(time.Millisecond))
			if from, err := time.Parse("2006-01-02", fromDate); err == nil {
				if createdTime.Before(from) {
					continue
				}
			}
		}
		if toDate != "" {
			createdTime := time.Unix(0, int64(parseTimestamp(invoice.CreatedAt))*int64(time.Millisecond))
			if to, err := time.Parse("2006-01-02", toDate); err == nil {
				if createdTime.After(to.Add(24 * time.Hour)) {
					continue
				}
			}
		}

		filtered = append(filtered, invoice)
	}

	// Generate CSV
	var csvData strings.Builder
	csvData.WriteString("Invoice ID,Title,Customer Name,Email,Phone,Amount,Status,Created At,Due Date,Paid At,Payment Method\n")

	for _, invoice := range filtered {
		createdAt := time.Unix(0, int64(parseTimestamp(invoice.CreatedAt))*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
		dueDate := time.Unix(0, int64(parseTimestamp(invoice.DueDate))*int64(time.Millisecond)).Format("2006-01-02 15:04:05")

		paidAt := ""
		if invoice.PaidAt != nil {
			if paidAtStr, ok := invoice.PaidAt.(string); ok && paidAtStr != "" {
				paidAt = time.Unix(0, int64(parseTimestamp(paidAtStr))*int64(time.Millisecond)).Format("2006-01-02 15:04:05")
			}
		}

		paymentMethod := ""
		if len(invoice.Payments) > 0 {
			paymentMethod = invoice.Payments[0].PaymentMethodID
		}

		csvData.WriteString(fmt.Sprintf("%s,%s,%s,%s,%s,%d,%s,%s,%s,%s,%s\n",
			invoice.InvoiceID,
			invoice.Title,
			invoice.Fullname,
			invoice.Email,
			invoice.Phone,
			invoice.Amount,
			invoice.InvoiceStatusID,
			createdAt,
			dueDate,
			paidAt,
			paymentMethod,
		))
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=invoices_%s.csv", time.Now().Format("20060102_150405")))
	c.String(200, csvData.String())
}

// Helper function to parse timestamp
func parseTimestamp(timestampStr string) int64 {
	if timestamp, err := strconv.ParseInt(timestampStr, 10, 64); err == nil {
		return timestamp
	}
	return 0
}
