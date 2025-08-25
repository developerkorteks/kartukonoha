package config

import "time"

// Config stores all configuration for the application.
// Values are read from constants but could be expanded to read from a file or environment variables.
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
}

// LoadConfig loads configuration from constants.
func LoadConfig() *Config {
	return &Config{
		SSOApiBaseURL: "https://sso.putri-veronica.my.id/api",
		NadiaApiURL:   "https://putri-veronica.my.id/api/user",
		SSOStaticKey:  "18e0cf008b474614a0b20f9c170cd33f",
		Username:      "kortkeks",
		Password:      "nabilalbab78",
		AdminAPIKey:   "nadia-admin-2024-secure-key",
		JWTSecret:     "nadia-jwt-secret-key-2024",
		ServerAddress: ":8080",
		DBPath:        "./nadia_transactions.db",
		TokenExpiry:   8 * time.Hour,
	}
}
