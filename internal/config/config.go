package config

import (
	"os"
	"strconv"
	"time"
)

// Config stores all configuration for the application.
// Values are read from environment variables with fallback to constants.
type Config struct {
	SSOApiBaseURL string
	NadiaApiURL   string
	SSOStaticKey  string
	Username      string
	Password      string
	AdminAPIKey   string
	JWTSecret     string
	ServerAddress string
	DBPath        string
	TokenExpiry   time.Duration
	Environment   string
	SwaggerHost   string
}

// LoadConfig loads configuration from environment variables with fallback to constants.
func LoadConfig() *Config {
	config := &Config{
		SSOApiBaseURL: getEnv("SSO_API_BASE_URL", "https://sso.putri-veronica.my.id/api"),
		NadiaApiURL:   getEnv("NADIA_API_URL", "https://putri-veronica.my.id/api/user"),
		SSOStaticKey:  getEnv("SSO_STATIC_KEY", "18e0cf008b474614a0b20f9c170cd33f"),
		Username:      getEnv("NADIA_USERNAME", "kortkeks"),
		Password:      getEnv("NADIA_PASSWORD", "nabilalbab78"),
		AdminAPIKey:   getEnv("ADMIN_API_KEY", "nadia-admin-2024-secure-key"),
		JWTSecret:     getEnv("JWT_SECRET", "nadia-jwt-secret-key-2024"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DBPath:        getEnv("DB_PATH", "./nadia_transactions.db"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		SwaggerHost:   getEnv("SWAGGER_HOST", ""),
		TokenExpiry:   8 * time.Hour,
	}

	// Handle PORT environment variable (common in cloud hosting)
	if port := os.Getenv("PORT"); port != "" {
		// Bind to all interfaces for cloud hosting
		config.ServerAddress = "0.0.0.0:" + port
	}

	// Set token expiry from environment if provided
	if tokenExpiryStr := os.Getenv("TOKEN_EXPIRY_HOURS"); tokenExpiryStr != "" {
		if hours, err := strconv.Atoi(tokenExpiryStr); err == nil {
			config.TokenExpiry = time.Duration(hours) * time.Hour
		}
	}

	return config
}

// getEnv gets environment variable with fallback to default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// IsProduction returns true if running in production environment
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// IsDevelopment returns true if running in development environment
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}
