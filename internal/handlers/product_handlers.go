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
	"github.com/nabilulilalbab/nadia/internal/services"
)

// Build a map[package_code]price from price-list-all.json
func (h *HTTPHandler) fetchPriceMap() (map[string]int, error) {
	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/price-list-all.json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var raw struct {
		Data []struct {
			PackageCode string `json:"package_code"`
			Price       int    `json:"price"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}
	m := make(map[string]int, len(raw.Data))
	for _, it := range raw.Data {
		m[it.PackageCode] = it.Price
	}
	return m, nil
}

// Convert Package to Product using a provided price map and markup
func packageToProductWithPriceMap(pkg models.Package, priceMap map[string]int, markup int, ns *services.NadiaService) models.Product {
	price := priceMap[pkg.PackageCode]
	finalPrice := price + markup
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
		IsAkrab:                 pkg.IsAkrab,
		IsCircle:                pkg.IsCircle,
		SedangGangguan:          pkg.SedangGangguan,
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

	// Fetch packages information
	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/package-list-all.json", nil)
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

	// Convert packages to products with actual price + 1500
	var packages []models.Package
	packagesData, _ := json.Marshal(nadiaResp.Data)
	json.Unmarshal(packagesData, &packages)

	// Build price map once per request
	priceMap, err := h.fetchPriceMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	var products []models.Product
	for _, pkg := range packages {
		products = append(products, packageToProductWithPriceMap(pkg, priceMap, 1500, h.nadiaService))
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

	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/package-list-all.json", nil)
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

	// Build price map once per request
	priceMap, err := h.fetchPriceMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	var filteredProducts []models.Product
	for _, pkg := range packages {
		product := packageToProductWithPriceMap(pkg, priceMap, 1500, h.nadiaService)

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

	// Fetch packages information
	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/package-list-all.json", nil)
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

	// Convert packages to products with actual price + 500
	var packages []models.Package
	packagesData, _ := json.Marshal(nadiaResp.Data)
	json.Unmarshal(packagesData, &packages)

	// Build price map once per request (SAME AS USER ENDPOINT)
	priceMap, err := h.fetchPriceMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	var products []models.Product
	for _, pkg := range packages {
		products = append(products, packageToProductWithPriceMap(pkg, priceMap, 500, h.nadiaService))
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

	resp, err := h.nadiaService.MakeRequest("POST", "/limited/xl/package-list-all.json", nil)
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

	// Build price map once per request (SAME AS USER ENDPOINT)
	priceMap, err := h.fetchPriceMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.APIResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to fetch price data: " + err.Error(),
			Success:    false,
		})
		return
	}

	var filteredProducts []models.Product
	for _, pkg := range packages {
		product := packageToProductWithPriceMap(pkg, priceMap, 500, h.nadiaService)

		// Apply filters
		if searchReq.Query != "" {
			query := strings.ToLower(searchReq.Query)
			if !strings.Contains(strings.ToLower(product.PackageName), query) &&
				!strings.Contains(strings.ToLower(product.PackageDescription), query) {
				continue
			}
		}

		// Filter by price from new endpoint + 500
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
