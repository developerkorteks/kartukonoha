package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/models"
)

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

	payload := map[string]string{"transaction_id": req.TransactionID}
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
	c.JSON(http.StatusOK, models.APIResponse{StatusCode: http.StatusOK, Message: "Invoice endpoint not fully implemented in this refactor.", Success: true})
}