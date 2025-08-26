package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nabilulilalbab/nadia/docs" // swagger docs
	"github.com/nabilulilalbab/nadia/internal/config"
	"github.com/nabilulilalbab/nadia/internal/database"
	"github.com/nabilulilalbab/nadia/internal/handlers"
	"github.com/nabilulilalbab/nadia/internal/middleware"
	"github.com/nabilulilalbab/nadia/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Nadia Package Purchase API
// @version 2.0
// @description API untuk pembelian paket XL dengan flow yang disederhanakan dan user-friendly
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host
// @BasePath /
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @description API Key for authentication. Use: nadia-admin-2024-secure-key

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and the token from /api/auth/login.

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Set Gin mode based on environment
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Configure Swagger host dynamically
	setupSwaggerHost(cfg)

	// Initialize database
	db, err := database.InitDatabase(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize services
	tokenManager := services.NewTokenManager(cfg)
	nadiaService := services.NewNadiaService(cfg, tokenManager)
	transactionService, err := services.NewTransactionService(db)
	if err != nil {
		log.Fatalf("Failed to initialize transaction service: %v", err)
	}

	// Initialize handlers
	httpHandler := handlers.NewHTTPHandler(nadiaService, transactionService)

	// Initialize Gin router
	r := gin.Default()

	// Setup middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.MonitoringMiddleware())

	// API Routes
	api := r.Group("/api")
	{
		// System
		api.GET("/health", httpHandler.HealthCheck)

		// Authentication (public)
		api.POST("/auth/login", httpHandler.AuthLogin)

		// Monitoring endpoints (protected)
		monitoring := api.Group("/monitoring")
		monitoring.Use(middleware.AuthMiddleware(cfg.AdminAPIKey))
		{
			monitoring.GET("/metrics", httpHandler.GetSystemMetrics)
			monitoring.GET("/metrics/realtime", httpHandler.GetRealtimeMetrics)
			monitoring.GET("/security", httpHandler.GetSecurityMetrics)
			monitoring.GET("/logs", httpHandler.GetApplicationLogs)
			monitoring.GET("/performance", httpHandler.GetPerformanceMetrics)
			monitoring.GET("/uptime", httpHandler.GetUptimeStatus)
		}

		// Main API endpoints (protected)
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.AdminAPIKey))
		{
			// Packages
			protected.GET("/packages", httpHandler.GetAllPackages)
			protected.POST("/packages/search", httpHandler.SearchPackages)
			protected.GET("/packages/stock", httpHandler.GetPackageStock)

			// OTP
			protected.POST("/otp/request", httpHandler.RequestOTP)
			protected.POST("/otp/verify", httpHandler.VerifyOTP)

			// Purchase
			protected.POST("/purchase", httpHandler.PurchasePackage)

			// Card Management
			protected.POST("/card/status", httpHandler.CheckCardStatus)
			protected.POST("/card/packages", httpHandler.CheckActivePackages)

			// Wallet
			protected.GET("/balance", httpHandler.GetBalance)
			protected.GET("payment-methods", httpHandler.GetPaymentMethods)

			// Transaction
			protected.POST("/transaction/check", httpHandler.CheckTransaction)
			protected.GET("/transactions", httpHandler.GetTransactions)

			// Invoice
			protected.GET("/invoices", httpHandler.GetInvoices)
			protected.GET("/invoices/:id", httpHandler.GetInvoiceDetail)
			protected.GET("/invoice/stats", httpHandler.GetInvoiceStatsHandler)

			// Analytics
			protected.GET("/stats/daily", httpHandler.GetDailyStats)

			// Export
			protected.GET("/export/transactions", httpHandler.ExportTransactions)
			protected.GET("/export/invoices", httpHandler.ExportInvoices)
		}

		// User endpoints (API Key or JWT protected, for end users with manipulated prices)
		userGroup := api.Group("/user")
		userGroup.Use(middleware.HybridAuthMiddleware(cfg.AdminAPIKey, cfg.JWTSecret))
		{
			// Products (with manipulated prices +1500)
			userGroup.GET("/products", httpHandler.GetAllProducts)
			userGroup.POST("/products/search", httpHandler.SearchProducts)
			userGroup.GET("/products/stock", httpHandler.GetProductStock)
		}

		// Reseller endpoints (API Key or JWT protected, for resellers with manipulated prices)
		resellerGroup := api.Group("/reseller")
		resellerGroup.Use(middleware.HybridAuthMiddleware(cfg.AdminAPIKey, cfg.JWTSecret))
		{
			// Products (with manipulated prices +500)
			resellerGroup.GET("/products", httpHandler.GetAllResellerProducts)
			resellerGroup.POST("/products/search", httpHandler.SearchResellerProducts)
			resellerGroup.GET("/products/stock", httpHandler.GetResellerProductStock)
		}

		// Dashboard & other data endpoints (publicly accessible data, but might need auth in real life)
		api.GET("/dashboard", httpHandler.GetDashboardData)
		api.GET("/transactions/:id", httpHandler.GetTransactionDetail)
	}

	// Serve Static HTML files & Swagger
	serveStatic(r)

	// Log startup information
	log.Printf("Starting Nadia API Server v2.0 on %s", cfg.ServerAddress)
	log.Printf("Environment: %s", cfg.Environment)

	// Log Swagger URL
	port := "8080"
	if cfg.ServerAddress != "" && cfg.ServerAddress != ":8080" {
		if len(cfg.ServerAddress) > 1 && cfg.ServerAddress[0] == ':' {
			port = cfg.ServerAddress[1:]
		}
	}
	log.Printf("Swagger Docs: http://localhost:%s/swagger/index.html (or use your server's IP/domain)", port)

	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func serveStatic(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/dashboard")
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("dashboard_complete.html")
	})

	r.GET("/dashboard/old", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("dashboard.html")
	})

	r.GET("/api-flow", func(c *gin.Context) {
		c.Header("Content-Type", "text/html")
		c.File("api_flow_example.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Serve favicon
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("favicon.ico")
	})
}

// setupSwaggerHost configures Swagger host dynamically
func setupSwaggerHost(cfg *config.Config) {
	// If SWAGGER_HOST is set, use it
	if cfg.SwaggerHost != "" {
		docs.SwaggerInfo.Host = cfg.SwaggerHost
		return
	}

	// Always leave host empty so Swagger UI uses window.location.host
	// This allows it to work with localhost, 127.0.0.1, or any other host
	docs.SwaggerInfo.Host = ""
}
