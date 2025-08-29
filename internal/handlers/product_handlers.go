package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/internal/models"
)

// Helper function to format price in Indonesian Rupiah
func formatRupiah(amount int) string {
	if amount == 0 {
		return "Rp. 0,00"
	}

	// Convert to string and add thousand separators
	str := strconv.Itoa(amount)
	n := len(str)
	if n <= 3 {
		return fmt.Sprintf("Rp. %s,00", str)
	}

	// Add dots for thousands
	var result strings.Builder
	for i, digit := range str {
		if i > 0 && (n-i)%3 == 0 {
			result.WriteString(".")
		}
		result.WriteRune(digit)
	}

	return fmt.Sprintf("Rp. %s,00", result.String())
}

// Helper function to fetch price data from the new price endpoint
func (h *HTTPHandler) fetchPriceData() (map[string]models.PriceData, error) {
	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/price-list-all.json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var priceResp models.APIResponse
	json.Unmarshal(body, &priceResp)

	var priceList []models.PriceData
	priceData, _ := json.Marshal(priceResp.Data)
	json.Unmarshal(priceData, &priceList)

	// Create a map for quick lookup by package_code
	priceMap := make(map[string]models.PriceData)
	for _, price := range priceList {
		priceMap[price.PackageCode] = price
	}

	return priceMap, nil
}

// Helper function to convert Package to Product with price from new endpoint + 1500
func packageToProductWithNewPrice(pkg models.Package, priceMap map[string]models.PriceData) models.Product {
	// Default to original price + 1500 if not found in new price data
	finalPrice := pkg.PackagePrice + 1500

	// Check if we have price data from the new endpoint
	if priceData, exists := priceMap[pkg.PackageCode]; exists {
		// Use member_price from new endpoint + 1500
		finalPrice = priceData.MemberPrice + 1500
	}

	return models.Product{
		PackageCode:             pkg.PackageCode,
		PackageName:             pkg.PackageName,
		PackageNameShort:        pkg.PackageNameShort,
		PackageDescription:      pkg.PackageDescription,
		PackagePrice:            finalPrice,
		PackagePriceFormatted:   formatRupiah(finalPrice),
		HaveDailyLimit:          pkg.HaveDailyLimit,
		DailyLimitDetails:       pkg.DailyLimitDetails,
		NoNeedLogin:             pkg.NoNeedLogin,
		CanMultiTrx:             pkg.CanMultiTrx,
		CanScheduledTrx:         pkg.CanScheduledTrx,
		HaveCutOffTime:          pkg.HaveCutOffTime,
		CutOffTime:              pkg.CutOffTime,
		NeedCheckStock:          pkg.NeedCheckStock,
		IsShowPaymentMethod:     pkg.IsShowPaymentMethod,
		AvailablePaymentMethods: pkg.AvailablePaymentMethods,
	}
}

// GetAllProducts godoc
// @Summary Get all available products for users
// @Description Retrieve all available products from Nadia API with manipulated prices (+1500 rupiah)
// @Tags products
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of products returned" default(100)
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/user/products [get]
// @Security ApiKeyAuth || BearerAuth
func (h *HTTPHandler) GetAllProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	limit, _ := strconv.Atoi(limitStr)

	// Fetch price data from new endpoint
	priceMap, err := h.fetchPriceData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch products: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	// Convert packages to products with prices from new endpoint
	var packages []models.Package
	packagesData, _ := json.Marshal(nadiaResp.Data)
	json.Unmarshal(packagesData, &packages)

	var products []models.Product
	for _, pkg := range packages {
		products = append(products, packageToProductWithNewPrice(pkg, priceMap))
	}

	// Apply limit if specified
	if limit > 0 && len(products) > limit {
		products = products[:limit]
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Retrieved %d products", len(products)),
		Success:    true,
		Data:       products,
	})
}

// SearchProducts godoc
// @Summary Search products with filters for users
// @Description Search and filter products by name, price, payment method, etc. with manipulated prices (+1500 rupiah)
// @Tags products
// @Accept json
// @Produce json
// @Param search body models.PackageSearchRequest true "Search filters"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/user/products/search [post]
// @Security ApiKeyAuth || BearerAuth
func (h *HTTPHandler) SearchProducts(c *gin.Context) {
	var searchReq models.PackageSearchRequest
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid search request: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Fetch price data from new endpoint
	priceMap, err := h.fetchPriceData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch products: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	// Convert packages to products with prices from new endpoint
	var packages []models.Package
	packagesData, _ := json.Marshal(nadiaResp.Data)
	json.Unmarshal(packagesData, &packages)

	var filteredProducts []models.Product
	for _, pkg := range packages {
		product := packageToProductWithNewPrice(pkg, priceMap)

		// Apply filters
		if searchReq.Query != "" {
			query := strings.ToLower(searchReq.Query)
			if !strings.Contains(strings.ToLower(product.PackageName), query) &&
				!strings.Contains(strings.ToLower(product.PackageDescription), query) {
				continue
			}
		}

		// Filter by price from new endpoint + 1500
		if searchReq.MaxPrice > 0 && product.PackagePrice > searchReq.MaxPrice {
			continue
		}
		if searchReq.MinPrice > 0 && product.PackagePrice < searchReq.MinPrice {
			continue
		}

		// Filter by payment method
		if searchReq.PaymentMethod != "" {
			hasPaymentMethod := false
			for _, pm := range product.AvailablePaymentMethods {
				if pm.PaymentMethod == searchReq.PaymentMethod {
					hasPaymentMethod = true
					break
				}
			}
			if !hasPaymentMethod {
				continue
			}
		}

		filteredProducts = append(filteredProducts, product)
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Found %d products matching your criteria", len(filteredProducts)),
		Success:    true,
		Data:       filteredProducts,
	})
}

// GetProductStock godoc
// @Summary Get product stock information for users
// @Description Retrieve stock information for all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/user/products/stock [get]
// @Security ApiKeyAuth || BearerAuth
func (h *HTTPHandler) GetProductStock(c *gin.Context) {
	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/check-stock-package-global.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get product stock: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}

// Helper function to convert package to reseller product with price from new endpoint + 500
func convertToResellerProductWithNewPrice(pkg models.Package, priceMap map[string]models.PriceData) models.Product {
	// Default to original price + 500 if not found in new price data
	finalPrice := pkg.PackagePrice + 500

	// Check if we have price data from the new endpoint
	if priceData, exists := priceMap[pkg.PackageCode]; exists {
		// Use reseller_price from new endpoint + 500
		finalPrice = priceData.ResellerPrice + 500
	}

	return models.Product{
		PackageCode:             pkg.PackageCode,
		PackageName:             pkg.PackageName,
		PackageNameShort:        pkg.PackageNameShort,
		PackageDescription:      pkg.PackageDescription,
		PackagePrice:            finalPrice,
		PackagePriceFormatted:   formatRupiah(finalPrice),
		HaveDailyLimit:          pkg.HaveDailyLimit,
		DailyLimitDetails:       pkg.DailyLimitDetails,
		NoNeedLogin:             pkg.NoNeedLogin,
		CanMultiTrx:             pkg.CanMultiTrx,
		CanScheduledTrx:         pkg.CanScheduledTrx,
		HaveCutOffTime:          pkg.HaveCutOffTime,
		CutOffTime:              pkg.CutOffTime,
		NeedCheckStock:          pkg.NeedCheckStock,
		IsShowPaymentMethod:     pkg.IsShowPaymentMethod,
		AvailablePaymentMethods: pkg.AvailablePaymentMethods,
	}
}

// GetAllResellerProducts godoc
// @Summary Get all available products for resellers
// @Description Retrieve all available products from Nadia API with prices from new price endpoint (reseller_price + 500 rupiah)
// @Tags reseller-products
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of products returned" default(100)
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/reseller/products [get]
// @Security ApiKeyAuth || BearerAuth
func (h *HTTPHandler) GetAllResellerProducts(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	limit, _ := strconv.Atoi(limitStr)

	// Fetch price data from new endpoint
	priceMap, err := h.fetchPriceData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch products: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	// Convert packages to reseller products with prices from new endpoint
	var products []models.Product
	if packages, ok := nadiaResp.Data.([]interface{}); ok {
		for _, pkg := range packages {
			if pkgMap, ok := pkg.(map[string]interface{}); ok {
				var packageData models.Package
				jsonData, _ := json.Marshal(pkgMap)
				json.Unmarshal(jsonData, &packageData)
				products = append(products, convertToResellerProductWithNewPrice(packageData, priceMap))
			}
		}
	}

	// Apply limit if specified
	if limit > 0 && len(products) > limit {
		products = products[:limit]
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Retrieved %d products", len(products)),
		Success:    true,
		Data:       products,
	})
}

// SearchResellerProducts godoc
// @Summary Search products with filters for resellers
// @Description Search and filter products by name, price, payment method, etc. with prices from new price endpoint (reseller_price + 500 rupiah)
// @Tags reseller-products
// @Accept json
// @Produce json
// @Param search body models.PackageSearchRequest true "Search filters"
// @Success 200 {object} models.APIResponse
// @Failure 400 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/reseller/products/search [post]
// @Security ApiKeyAuth || BearerAuth
func (h *HTTPHandler) SearchResellerProducts(c *gin.Context) {
	var searchReq models.PackageSearchRequest
	if err := c.ShouldBindJSON(&searchReq); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid search request: " + err.Error(),
			Success:    false,
		})
		return
	}

	// Fetch price data from new endpoint
	priceMap, err := h.fetchPriceData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/package-list-all.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch products: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	// Convert and filter packages with prices from new endpoint
	var filteredProducts []models.Product
	if packages, ok := nadiaResp.Data.([]interface{}); ok {
		for _, pkg := range packages {
			if pkgMap, ok := pkg.(map[string]interface{}); ok {
				var packageData models.Package
				jsonData, _ := json.Marshal(pkgMap)
				json.Unmarshal(jsonData, &packageData)

				product := convertToResellerProductWithNewPrice(packageData, priceMap)

				// Apply filters
				if searchReq.Query != "" {
					query := strings.ToLower(searchReq.Query)
					if !strings.Contains(strings.ToLower(product.PackageName), query) &&
						!strings.Contains(strings.ToLower(product.PackageDescription), query) {
						continue
					}
				}

				if searchReq.MinPrice > 0 && product.PackagePrice < searchReq.MinPrice {
					continue
				}

				if searchReq.MaxPrice > 0 && product.PackagePrice > searchReq.MaxPrice {
					continue
				}

				if searchReq.PaymentMethod != "" {
					hasPaymentMethod := false
					for _, pm := range product.AvailablePaymentMethods {
						if pm.PaymentMethod == searchReq.PaymentMethod {
							hasPaymentMethod = true
							break
						}
					}
					if !hasPaymentMethod {
						continue
					}
				}

				filteredProducts = append(filteredProducts, product)
			}
		}
	}

	c.JSON(http.StatusOK, models.APIResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Found %d products matching your criteria", len(filteredProducts)),
		Success:    true,
		Data:       filteredProducts,
	})
}

// GetResellerProductStock godoc
// @Summary Get product stock information for resellers
// @Description Retrieve stock information for all products
// @Tags reseller-products
// @Accept json
// @Produce json
// @Success 200 {object} models.APIResponse
// @Failure 500 {object} models.APIResponse
// @Router /api/reseller/products/stock [get]
// @Security ApiKeyAuth || BearerAuth
func (h *HTTPHandler) GetResellerProductStock(c *gin.Context) {
	resp, err := h.nadiaService.MakeRequest("GET", "/limited/xl/check-stock-package-global.json", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get product stock: " + err.Error(),
			Success:    false,
		})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var nadiaResp models.APIResponse
	json.Unmarshal(body, &nadiaResp)

	c.JSON(resp.StatusCode, nadiaResp)
}
