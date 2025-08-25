package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/models"
	"github.com/nabilulilalbab/nadia/internal/services"
	"github.com/nabilulilalbab/nadia/internal/utils"
)

// HTTPHandler holds all dependencies for the handlers.
type HTTPHandler struct {
	nadiaService       *services.NadiaService
	transactionService *services.TransactionService
}

// NewHTTPHandler creates a new HTTPHandler.
func NewHTTPHandler(ns *services.NadiaService, ts *services.TransactionService) *HTTPHandler {
	return &HTTPHandler{
		nadiaService:       ns,
		transactionService: ts,
	}
}

// HealthCheck godoc
// @Summary Health check
// @Description Check if the API is running properly
// @Tags system
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Router /api/health [get]
func (h *HTTPHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "API is running successfully",
		Success:    true,
		Data: map[string]interface{}{
			"timestamp": time.Now().Format(time.RFC3339),
			"version":   "2.0.0",
		},
	})
}

// AuthLogin godoc
// @Summary Authenticate with API key
// @Description Login with API key to get access token
// @Tags authentication
// @Accept json
// @Produce json
// @Param request body models.AuthRequest true "Authentication request"
// @Success 200 {object} models.APIResponse{data=models.AuthResponse}
// @Failure 400 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/auth/login [post]
func (h *HTTPHandler) AuthLogin(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request format",
			Success:    false,
		})
		return
	}

	// In a real app, you'd compare against a configured value
	if req.APIKey != "nadia-admin-2024-secure-key" {
		c.JSON(http.StatusUnauthorized, models.APIResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid API key",
			Success:    false,
		})
		return
	}

	token := fmt.Sprintf("nadia-token-%d", time.Now().Unix())
	expiresAt := time.Now().Add(24 * time.Hour)

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Authentication successful",
		Success:    true,
		Data: models.AuthResponse{
			Token:     token,
			ExpiresAt: expiresAt,
		},
	})
}

// GetAllPackages godoc
// @Summary Get all available packages
// @Description Retrieve all available packages from Nadia API without filtering
// @Tags packages
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of packages returned" default(100)
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/packages [get]
func (h *HTTPHandler) GetAllPackages(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	limit, _ := strconv.Atoi(limitStr)

	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to fetch packages: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp struct {
		Data []models.Package `json:"data"`
	}
	json.Unmarshal(body, &nadiaResp)

	packages := nadiaResp.Data
	if limit > 0 && len(packages) > limit {
		packages = packages[:limit]
	}

	c.JSON(http.StatusOK, models.APIResponse{StatusCode: http.StatusOK, Message: fmt.Sprintf("Retrieved %d packages", len(packages)), Success: true, Data: packages})
}

// SearchPackages godoc
// @Summary Search packages with filters
// @Description Search and filter packages by name, price, payment method, etc.
// @Tags packages
// @Accept json
// @Produce json
// @Param search body models.PackageSearchRequest true "Search filters"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/packages/search [post]
func (h *HTTPHandler) SearchPackages(c *gin.Context) {
	var searchReq models.PackageSearchRequest
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: "Invalid search request: " + err.Error(), Success: false})
		return
	}

	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to fetch packages: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp struct {
		Data []models.Package `json:"data"`
	}
	json.Unmarshal(body, &nadiaResp)

	var filteredPackages []models.Package
	for _, pkg := range nadiaResp.Data {
		if searchReq.Query != "" {
			query := strings.ToLower(searchReq.Query)
			if !strings.Contains(strings.ToLower(pkg.PackageName), query) && !strings.Contains(strings.ToLower(pkg.PackageDescription), query) {
				continue
			}
		}
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
		if searchReq.MinPrice > 0 && pkg.PackagePrice < searchReq.MinPrice {
			continue
		}
		if searchReq.MaxPrice > 0 && pkg.PackagePrice > searchReq.MaxPrice {
			continue
		}
		filteredPackages = append(filteredPackages, pkg)
	}

	c.JSON(http.StatusOK, models.APIResponse{StatusCode: http.StatusOK, Message: fmt.Sprintf("Found %d packages matching your criteria", len(filteredPackages)), Success: true, Data: filteredPackages})
}

// RequestOTP godoc
// @Summary Request OTP for phone number
// @Description Request OTP code to be sent to the specified phone number.
// @Tags otp
// @Accept json
// @Produce json
// @Param request body models.SimpleOTPRequest true "OTP request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/otp/request [post]
func (h *HTTPHandler) RequestOTP(c *gin.Context) {
	var req models.SimpleOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: "Invalid request: " + err.Error(), Success: false})
		return
	}

	payload := map[string]string{"phone": req.PhoneNumber}
	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/request-otp.json", payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to request OTP: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	if nadiaResp.Success {
		if dataMap, ok := nadiaResp.Data.(map[string]interface{}); ok {
			if authID, ok := dataMap["auth_id"].(string); ok {
				h.transactionService.CreateOTPSession(req.PhoneNumber, authID)
			}
		}
	}

	c.JSON(resp.StatusCode, nadiaResp)
}

// VerifyOTP godoc
// @Summary Verify OTP and get access token
// @Description Verify OTP code and get access token for purchasing packages
// @Tags otp
// @Accept json
// @Produce json
// @Param request body models.SimpleVerifyOTPRequest true "Verify OTP request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/otp/verify [post]
func (h *HTTPHandler) VerifyOTP(c *gin.Context) {
	var req models.SimpleVerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: "Invalid request: " + err.Error(), Success: false})
		return
	}

	session, err := h.transactionService.GetOTPSession(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: err.Error(), Success: false})
		return
	}

	payload := map[string]string{"phone": req.PhoneNumber, "auth_id": session.AuthID, "otp": req.OTPCode}
	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/request-login.json", payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to verify OTP: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// PurchasePackage godoc
// @Summary Purchase a package with access token
// @Description Purchase a package using phone number, package code, payment method, and access token
// @Tags purchase
// @Accept json
// @Produce json
// @Param request body models.SimplePurchaseRequest true "Purchase request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/purchase [post]
func (h *HTTPHandler) PurchasePackage(c *gin.Context) {
	var req models.SimplePurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: "Invalid request: " + err.Error(), Success: false})
		return
	}

	source := req.Source
	if source == "" {
		source = "api_direct"
	}
	txRecord := h.transactionService.RecordTransaction(req.PhoneNumber, req.PackageCode, "", req.PaymentMethod, source)

	payload := map[string]interface{}{
		"package_code":   req.PackageCode,
		"phone":          utils.FormatPhoneNumber(req.PhoneNumber, true),
		"access_token":   req.AccessToken,
		"payment_method": req.PaymentMethod,
	}

	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/beli-paket-otp.json", payload)
	if err != nil {
		h.transactionService.UpdateTransactionStatus(txRecord.ID, "FAILED", "", 0, 0, err.Error())
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to purchase package: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	if nadiaResp.Success {
		amount := h.nadiaService.GetPackagePrice(req.PackageCode)
		var trxID, packageName string
		var processingFee int
		if dataMap, ok := nadiaResp.Data.(map[string]interface{}); ok {
			trxID, _ = dataMap["trx_id"].(string)
			packageName, _ = dataMap["package_name"].(string)
			if fee, ok := dataMap["package_processing_fee"].(float64); ok {
				processingFee = int(fee)
			}
		}
		h.transactionService.UpdateTransactionStatus(txRecord.ID, "SUCCESS", trxID, amount, processingFee, "")
		if packageName != "" {
			h.transactionService.UpdateTransactionPackageName(txRecord.ID, packageName)
		}
		h.transactionService.DeleteOTPSession(req.PhoneNumber)
	} else {
		h.transactionService.UpdateTransactionStatus(txRecord.ID, "FAILED", "", 0, 0, nadiaResp.Message)
	}

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetDashboardData godoc
// @Summary Get comprehensive dashboard data
// @Description Get dashboard data including stats, recent transactions, balance, and real-time monitoring metrics
// @Tags dashboard
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse{data=models.DashboardData}
// @Failure 500 {object} models.APIResponse
// @Router /api/dashboard [get]
func (h *HTTPHandler) GetDashboardData(c *gin.Context) {
	// This is a simplified version. A full implementation would fetch balance, invoice stats etc.
	dashboardData := models.DashboardData{
		Stats:              h.transactionService.GetSystemStats(),
		RecentTransactions: h.transactionService.GetRecentTransactions(10),
		SourceBreakdown:    h.transactionService.GetSourceStats(),
		SystemStatus:       "OPERATIONAL", // Simplified
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Dashboard data retrieved successfully",
		Success:    true,
		Data:       dashboardData,
	})
}

// GetTransactionDetail godoc
// @Summary Get transaction detail
// @Description Get detailed information about a specific transaction
// @Tags dashboard
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /api/transactions/{id} [get]
func (h *HTTPHandler) GetTransactionDetail(c *gin.Context) {
	id := c.Param("id")
	tx, exists := h.transactionService.GetTransactionDetail(id)
	if !exists {
		c.JSON(http.StatusNotFound, models.APIResponse{StatusCode: http.StatusNotFound, Message: "Transaction not found", Success: false})
		return
	}
	c.JSON(http.StatusOK, models.APIResponse{StatusCode: http.StatusOK, Message: "Transaction detail retrieved successfully", Success: true, Data: tx})
}

// CheckCardStatus godoc
// @Summary Check XL card status and balance
// @Description Check the status, balance, and active period of an XL card using OTP verification
// @Tags card
// @Accept json
// @Produce json
// @Param request body models.SimpleCheckStatusRequest true "Check status request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/card/status [post]
func (h *HTTPHandler) CheckCardStatus(c *gin.Context) {
	var req models.SimpleCheckStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: "Invalid request: " + err.Error(), Success: false})
		return
	}

	session, err := h.transactionService.GetOTPSession(req.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: err.Error(), Success: false})
		return
	}

	accessToken := req.OTPCode + ":" + session.AuthID
	payload := map[string]string{"access_token": accessToken}

	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/status-kartu.json", payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to check card status: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}
