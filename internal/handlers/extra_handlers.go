package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/models"
)

// GetTransactions godoc
// @Summary Get transactions with filters
// @Description Get transactions with optional filters like status, source, date range, and limit
// @Tags transactions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param limit query int false "Limit number of results" default(100)
// @Param status query string false "Filter by status (SUCCESS, FAILED, PENDING)"
// @Param source query string false "Filter by source (whatsapp_bot, etc)"
// @Param from query string false "Start date (YYYY-MM-DD)"
// @Param to query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} models.APIResponse{data=[]models.TransactionRecord}
// @Failure 401 {object} models.APIResponse
// @Router /api/transactions [get]
func (h *HTTPHandler) GetTransactions(c *gin.Context) {
	// Parse query parameters
	limitStr := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}

	status := c.Query("status")
	source := c.Query("source")
	from := c.Query("from")
	to := c.Query("to")

	// Get transactions from service
	transactions, err := h.transactionService.GetTransactionsWithFilters(limit, status, source, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve transactions",
			Success:    false,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Transactions retrieved successfully",
		Success:    true,
		Data:       transactions,
	})
}

// GetDailyStats godoc
// @Summary Get daily transaction statistics
// @Description Get daily statistics for transactions including counts and amounts
// @Tags analytics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Success 200 {object} models.APIResponse
// @Failure 401 {object} models.APIResponse
// @Router /api/stats/daily [get]
func (h *HTTPHandler) GetDailyStats(c *gin.Context) {
	stats, err := h.transactionService.GetDailyStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve daily statistics",
			Success:    false,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Daily statistics retrieved successfully",
		Success:    true,
		Data:       stats,
	})
}

// ExportTransactions godoc
// @Summary Export transactions to CSV
// @Description Export transactions data to CSV format
// @Tags export
// @Accept json
// @Produce text/csv
// @Security ApiKeyAuth
// @Security BearerAuth
// @Success 200 {file} string "CSV file"
// @Failure 401 {object} models.APIResponse
// @Router /api/export/transactions [get]
func (h *HTTPHandler) ExportTransactions(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve transactions for export",
			Success:    false,
		})
		return
	}

	// Set CSV headers
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=transactions.csv")

	// Write CSV header
	csvData := "ID,Phone,Package,Amount,Status,Source,Created At\n"
	
	// Write transaction data
	for _, tx := range transactions {
		csvData += fmt.Sprintf("%s,%s,%s,%d,%s,%s,%s\n",
			tx.ID, tx.PhoneNumber, tx.PackageName, tx.Amount, tx.Status, tx.Source,
			tx.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	c.String(http.StatusOK, csvData)
}

// ExportInvoices godoc
// @Summary Export invoices to CSV
// @Description Export invoices data to CSV format
// @Tags export
// @Accept json
// @Produce text/csv
// @Security ApiKeyAuth
// @Security BearerAuth
// @Success 200 {file} string "CSV file"
// @Failure 401 {object} models.APIResponse
// @Router /api/export/invoices [get]
func (h *HTTPHandler) ExportInvoices(c *gin.Context) {
	invoices, err := h.transactionService.GetAllInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve invoices for export",
			Success:    false,
		})
		return
	}

	// Set CSV headers
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=invoices.csv")

	// Write CSV header
	csvData := "ID,Phone,Package,Amount,Status,Payment Method,Created At\n"
	
	// Write invoice data
	for _, inv := range invoices {
		csvData += fmt.Sprintf("%s,%s,%s,%d,%s,%s,%s\n",
			inv.ID, inv.PhoneNumber, inv.PackageName, inv.Amount, inv.Status, inv.PaymentMethod,
			inv.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	c.String(http.StatusOK, csvData)
}

// CheckActivePackages godoc
// @Summary Check active packages on XL card
// @Description Check all active packages/quotas on an XL card using OTP verification
// @Tags card
// @Accept json
// @Produce json
// @Param request body models.SimpleCheckStatusRequest true "Check active packages request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/card/packages [post]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) CheckActivePackages(c *gin.Context) {
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

	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/package-active-list.json", payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to check active packages: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetBalance godoc
// @Summary Get account balance
// @Description Retrieve current account balance
// @Tags wallet
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/balance [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) GetBalance(c *gin.Context) {
	resp, err := h.nadiaService.MakeRequest("GET", "/wallet/balance.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to get balance: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// CheckTransaction godoc
// @Summary Check transaction status
// @Description Check the status of a transaction by transaction ID
// @Tags transaction
// @Accept json
// @Produce json
// @Param request body models.SimpleTransactionCheckRequest true "Transaction check request"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/transaction/check [post]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) CheckTransaction(c *gin.Context) {
	var req models.SimpleTransactionCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{StatusCode: http.StatusBadRequest, Message: "Invalid request: " + err.Error(), Success: false})
		return
	}

	payload := map[string]string{"trx_id": req.TransactionID}
	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/check-transaction.json", payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to check transaction: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetPaymentMethods godoc
// @Summary Get available payment methods
// @Description Retrieve all available payment methods
// @Tags payment
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/payment-methods [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) GetPaymentMethods(c *gin.Context) {
	resp, err := h.nadiaService.MakeRequest("GET", "/wallet/payment-methods.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to get payment methods: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetPackageStock godoc
// @Summary Get package stock information
// @Description Retrieve stock information for all packages
// @Tags packages
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/packages/stock [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) GetPackageStock(c *gin.Context) {
	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/check-stock-package-global.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to get package stock: " + err.Error(), Success: false})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// GetInvoices godoc
// @Summary Get invoice list
// @Description Retrieve list of invoices with pagination
// @Tags invoice
// @Accept json
// @Produce json
// @Param search query string false "Search term"
// @Param limit query int false "Limit per page" default(20)
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/invoices [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) GetInvoices(c *gin.Context) {
	// This is a simplified handler. A full implementation would exist in NadiaService.
	c.JSON(http.StatusOK, models.APIResponse{StatusCode: http.StatusOK, Message: "Invoice endpoint not fully implemented in this refactor.", Success: true})
}

// GetInvoiceDetail godoc
// @Summary Get invoice detail
// @Description Get detailed information about a specific invoice
// @Tags invoice
// @Accept json
// @Produce json
// @Param id path string true "Invoice ID"
// @Success 200 {object} models.APIResponse
// @Failure 404 {object} models.APIResponse
// @Router /api/invoices/{id} [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) GetInvoiceDetail(c *gin.Context) {
	c.JSON(http.StatusNotFound, models.APIResponse{StatusCode: http.StatusNotFound, Message: "Invoice endpoint not fully implemented in this refactor.", Success: false})
}

// GetInvoiceStatsHandler godoc
// @Summary Get invoice statistics
// @Description Get comprehensive invoice statistics including payment rates and revenue
// @Tags invoice
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Router /api/invoice/stats [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *HTTPHandler) GetInvoiceStatsHandler(c *gin.Context) {
	// Get invoice stats from transaction service (using transactions as invoices)
	stats, err := h.transactionService.GetInvoiceStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to retrieve invoice statistics",
			Success:    false,
		})
		return
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    "Invoice statistics retrieved successfully",
		Success:    true,
		Data: map[string]interface{}{
			"stats": stats,
		},
	})
}