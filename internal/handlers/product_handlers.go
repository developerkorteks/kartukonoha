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

// Helper function to convert Package to Product with manipulated price
func packageToProduct(pkg models.Package) models.Product {
	// Add 1500 to original price
	manipulatedPrice := pkg.PackagePrice + 1500

	return models.Product{
		PackageCode:             pkg.PackageCode,
		PackageName:             pkg.PackageName,
		PackageNameShort:        pkg.PackageNameShort,
		PackageDescription:      pkg.PackageDescription,
		PackagePrice:            manipulatedPrice,
		PackagePriceFormatted:   formatRupiah(manipulatedPrice),
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

	// Convert packages to products with manipulated prices
	var packages []models.Package
	packagesData, _ := json.Marshal(nadiaResp.Data)
	json.Unmarshal(packagesData, &packages)

	var products []models.Product
	for _, pkg := range packages {
		products = append(products, packageToProduct(pkg))
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

	// Convert packages to products with manipulated prices
	var packages []models.Package
	packagesData, _ := json.Marshal(nadiaResp.Data)
	json.Unmarshal(packagesData, &packages)

	var filteredProducts []models.Product
	for _, pkg := range packages {
		product := packageToProduct(pkg)

		// Apply filters
		if searchReq.Query != "" {
			query := strings.ToLower(searchReq.Query)
			if !strings.Contains(strings.ToLower(product.PackageName), query) &&
				!strings.Contains(strings.ToLower(product.PackageDescription), query) {
				continue
			}
		}

		// Filter by manipulated price (not original price)
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

// Helper function to convert package to reseller product (with +500 price manipulation)
func convertToResellerProduct(pkg models.Package) models.Product {
	manipulatedPrice := pkg.PackagePrice + 500 // Add 500 rupiah for resellers
	return models.Product{
		PackageCode:             pkg.PackageCode,
		PackageName:             pkg.PackageName,
		PackageNameShort:        pkg.PackageNameShort,
		PackageDescription:      pkg.PackageDescription,
		PackagePrice:            manipulatedPrice,
		PackagePriceFormatted:   formatRupiah(manipulatedPrice),
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
// @Description Retrieve all available products from Nadia API with manipulated prices (+500 rupiah)
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

	// Convert packages to reseller products with manipulated prices
	var products []models.Product
	if packages, ok := nadiaResp.Data.([]interface{}); ok {
		for _, pkg := range packages {
			if pkgMap, ok := pkg.(map[string]interface{}); ok {
				var packageData models.Package
				jsonData, _ := json.Marshal(pkgMap)
				json.Unmarshal(jsonData, &packageData)
				products = append(products, convertToResellerProduct(packageData))
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
// @Description Search and filter products by name, price, payment method, etc. with manipulated prices (+500 rupiah)
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

	// Convert and filter packages
	var filteredProducts []models.Product
	if packages, ok := nadiaResp.Data.([]interface{}); ok {
		for _, pkg := range packages {
			if pkgMap, ok := pkg.(map[string]interface{}); ok {
				var packageData models.Package
				jsonData, _ := json.Marshal(pkgMap)
				json.Unmarshal(jsonData, &packageData)

				product := convertToResellerProduct(packageData)

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
