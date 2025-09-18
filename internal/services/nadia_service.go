package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"

	"github.com/nabilulilalbab/nadia/internal/config"
)

// TokenManager handles the lifecycle of the API token.
type TokenManager struct {
	cfg       *config.Config
	token     string
	expiresAt time.Time
	mutex     sync.RWMutex
}

// NewTokenManager creates a new TokenManager.
func NewTokenManager(cfg *config.Config) *TokenManager {
	return &TokenManager{
		cfg: cfg,
	}
}

// GetValidToken returns a valid token, refreshing it if necessary.
func (tm *TokenManager) GetValidToken() (string, error) {
	tm.mutex.RLock()
	if tm.token != "" && time.Now().Before(tm.expiresAt) {
		token := tm.token
		tm.mutex.RUnlock()
		return token, nil
	}
	tm.mutex.RUnlock()

	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// Double check after acquiring the lock
	if tm.token != "" && time.Now().Before(tm.expiresAt) {
		return tm.token, nil
	}

	newToken, err := tm.refreshToken()
	if err != nil {
		return "", err
	}

	tm.token = newToken
	tm.expiresAt = time.Now().Add(tm.cfg.TokenExpiry)
	return newToken, nil
}

// refreshToken performs the complete login flow to get a new token.
func (tm *TokenManager) refreshToken() (string, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", fmt.Errorf("failed to create cookie jar: %v", err)
	}
	httpClient := &http.Client{Jar: jar}

	sessionToken, err := tm.getSessionToken(httpClient)
	if err != nil {
		return "", fmt.Errorf("failed to get session token: %v", err)
	}

	finalJWT, err := tm.performApiLogin(httpClient, sessionToken)
	if err != nil {
		return "", fmt.Errorf("failed to perform login: %v", err)
	}

	return finalJWT, nil
}

// getSessionToken requests a temporary session token.
func (tm *TokenManager) getSessionToken(client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", tm.cfg.SSOApiBaseURL+"/oauth/request_token", nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("access-key", tm.cfg.SSOStaticKey)
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

	if resp.StatusCode != http.StatusOK {
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

// performApiLogin uses the session token to perform the final login.
func (tm *TokenManager) performApiLogin(client *http.Client, sessionToken string) (string, error) {
	payload := map[string]string{
		"username": tm.cfg.Username,
		"password": tm.cfg.Password,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", tm.cfg.SSOApiBaseURL+"/user/login", bytes.NewBuffer(payloadBytes))
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

	if resp.StatusCode != http.StatusOK {
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

// NadiaService provides methods to interact with the external Nadia API.
type NadiaService struct {
	cfg          *config.Config
	tokenManager *TokenManager
}

// NewNadiaService creates a new NadiaService.
func NewNadiaService(cfg *config.Config, tm *TokenManager) *NadiaService {
	return &NadiaService{
		cfg:          cfg,
		tokenManager: tm,
	}
}

// MakeRequest executes a request to the Nadia API.
func (s *NadiaService) MakeRequest(method, endpoint string, payload interface{}) (*http.Response, error) {
	return s.makeRequestWithRetry(method, endpoint, payload, 0)
}

func (s *NadiaService) makeRequestWithRetry(method, endpoint string, payload interface{}, retryCount int) (*http.Response, error) {
	token, err := s.tokenManager.GetValidToken()
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

	req, err := http.NewRequest(method, s.cfg.NadiaApiURL+endpoint, reqBody)
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
	if err == nil && resp.StatusCode == http.StatusUnauthorized && retryCount < 2 {
		resp.Body.Close()
		log.Printf("Received 401 Unauthorized, attempting JWT refresh (retry %d/2)", retryCount+1)

		// Force token refresh
		s.tokenManager.mutex.Lock()
		s.tokenManager.token = ""
		s.tokenManager.expiresAt = time.Time{}
		s.tokenManager.mutex.Unlock()

		// Retry the request
		return s.makeRequestWithRetry(method, endpoint, payload, retryCount+1)
	}

	return resp, err
}

// GetPackagePrice fetches the price for a specific package code.
func (s *NadiaService) GetPackagePrice(packageCode string) int {
	resp, err := s.MakeRequest("POST", "/limited/xl/price-list-all.json", nil)
	if err != nil {
		log.Printf("Failed to fetch price list for price lookup: %v", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read price response: %v", err)
		return 0
	}

	var nadiaResp struct {
		Data []struct {
			PackageCode string `json:"package_code"`
			PackageName string `json:"package_name"`
			Price       int    `json:"price"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &nadiaResp); err != nil {
		log.Printf("Failed to parse price response: %v", err)
		return 0
	}

	for _, priceData := range nadiaResp.Data {
		if priceData.PackageCode == packageCode {
			return priceData.Price
		}
	}

	log.Printf("Package price not found: %s", packageCode)
	return 0
}
