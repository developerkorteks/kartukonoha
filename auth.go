package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

// --- KONFIGURASI ---
const (
	ssoApiBaseURL = "https://sso.putri-veronica.my.id/api"
	ssoStaticKey  = "18e0cf008b474614a0b20f9c170cd33f" // Token dari URL login
	username      = "kortkeks"
	password      = "nabilalbab78"
)

// --------------------

// Struct untuk menampung token sesi sementara dari /request_token
type SessionTokenResponse struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
	Message     string `json:"message"`
}

// Struct untuk payload login ke /user/login
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Struct untuk menampung token JWT final dari /user/login
type FinalAuthResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
	Message string `json:"message"`
}

func main() {
	fmt.Println("Memulai proses login API SSO...")

	// 1. Buat HTTP client dengan cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Gagal membuat cookie jar: %v", err)
	}
	httpClient := &http.Client{Jar: jar}

	// 2. Langkah 1: Dapatkan token sesi sementara
	fmt.Println("Langkah 1: Meminta token sesi sementara...")
	sessionToken, err := getSessionToken(httpClient)
	if err != nil {
		log.Fatalf("Gagal mendapatkan token sesi: %v", err)
	}
	fmt.Printf("Berhasil mendapatkan token sesi: %s...\n", sessionToken[:20])

	// 3. Langkah 2: Lakukan login untuk mendapatkan JWT final
	fmt.Println("\nLangkah 2: Melakukan login dengan username/password...")
	finalJWT, err := performApiLogin(httpClient, sessionToken)
	if err != nil {
		log.Fatalf("Gagal melakukan login akhir: %v", err)
	}

	fmt.Println("\n==============================================")
	fmt.Println("         LOGIN SSO BERHASIL!")
	fmt.Println("==============================================")
	fmt.Println("Token JWT Anda adalah:")
	fmt.Println(finalJWT)
	fmt.Println("==============================================")
	fmt.Println("Gunakan token ini untuk konstanta 'authToken' di skrip menu utama Anda.")
}

func getSessionToken(client *http.Client) (string, error) {
	req, err := http.NewRequest("GET", ssoApiBaseURL+"/oauth/request_token", nil)
	if err != nil {
		return "", err
	}

	// Tambahkan header penting
	req.Header.Set("User-Agent", "Go-Auth-Client/1.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("access-key", ssoStaticKey)
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

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("gagal, server merespons dengan status %d. Respons: %s", resp.StatusCode, string(body))
	}

	var sessionResp SessionTokenResponse
	if err := json.Unmarshal(body, &sessionResp); err != nil {
		return "", fmt.Errorf("gagal parsing JSON token sesi: %w. Respons: %s", err, string(body))
	}

	if !sessionResp.Success || sessionResp.AccessToken == "" {
		return "", fmt.Errorf("API tidak memberikan token sesi yang valid. Pesan: %s", sessionResp.Message)
	}

	return sessionResp.AccessToken, nil
}

func performApiLogin(client *http.Client, sessionToken string) (string, error) {
	payload := LoginPayload{
		Username: username,
		Password: password,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", ssoApiBaseURL+"/user/login", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Tambahkan header penting
	req.Header.Set("User-Agent", "Go-Auth-Client/1.0")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("access-token", sessionToken) // Gunakan token sesi di sini
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

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("login akhir gagal dengan status %d. Respons: %s", resp.StatusCode, string(body))
	}

	var finalResp FinalAuthResponse
	if err := json.Unmarshal(body, &finalResp); err != nil {
		return "", fmt.Errorf("gagal parsing JSON final: %w. Respons: %s", err, string(body))
	}

	if !finalResp.Success || finalResp.Token == "" {
		return "", fmt.Errorf("API tidak memberikan token JWT final. Pesan: %s", finalResp.Message)
	}

	return finalResp.Token, nil
}
