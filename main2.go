package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// --- KONFIGURASI ---
// Ganti dengan token JWT Anda yang valid.
const authToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJrb3J0a2VrcyIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJ3PT1NNU02TlRaWElXMWlwdFozT1ROREEiLCJpYXQiOjE3NTYwNTM1MzMsImV4cCI6MTc1NjA4MjMzM30.FFtPvvJos5Wz-rcNNdIUj2plb_7RHW8baSVl299fw25A2ROZOC-q00MgcEJJBnJCq_KcskUF6AEnPMkVFv8cyieYWEeFELCHmBCLJi2jqkzQVkrPgR7DNmjbpzFnwKK_iMWFiej15o9TbK8gDNLL3b4S7zuU5PZBnraqiL1tm9Zx--K2kotNSOQ225ggI_7ca3BwpDjeKu1nF-iZMbMlgwd69gOkmonnIHHQeU8C2C2CQxfj3KazAv95Ho4_BT2yniptczzYI_xEg_Rk_yt51N1hdFd1LcLaZNJ1zhRCN9wnr6t_tKZRbstZdfWpfk03tdRaCKsNAZlpBMHC_CMBSA"
const baseURL = "https://putri-veronica.my.id"

// --------------------

// --- DEFINISI STRUCT UNTUK SETIAP API CALL ---

// Langkah 1: Request OTP
type RequestOTPPayload struct {
	Phone string `json:"phone"`
}
type RequestOTPResponse struct {
	Data struct {
		AuthID string `json:"auth_id"`
	} `json:"data"`
}

// Langkah 2: Verifikasi OTP
type VerifyOTPPayload struct {
	Phone  string `json:"phone"`
	AuthID string `json:"auth_id"`
	OTP    string `json:"otp"`
}
type VerifyOTPResponse struct {
	Data struct {
		AccessToken string `json:"access_token"`
	} `json:"data"`
}

// Langkah 3: Get Active Packages
type GetPackagesPayload struct {
	AccessToken string `json:"access_token"`
}

// Untuk GetPackagesResponse, kita gunakan map[string]interface{} agar fleksibel
// karena kita belum tahu pasti struktur data paket aktifnya.
type GetPackagesResponse map[string]interface{}

func main() {
	// Nomor telepon yang akan di-cek
	targetPhone := "087817739901" // Ganti dengan nomor target

	// =================================================================
	// LANGKAH 1: MEMINTA OTP
	// =================================================================
	fmt.Println("--- LANGKAH 1: Meminta OTP ---")

	// Siapkan payload untuk request OTP
	otpPayload := RequestOTPPayload{Phone: targetPhone}

	// Kirim request dan dapatkan auth_id
	var otpResponseData RequestOTPResponse
	err := apiCall("/api/user/limited/xl/request-otp.json", otpPayload, &otpResponseData)
	if err != nil {
		log.Fatalf("Gagal pada Langkah 1 (Request OTP): %v", err)
	}
	authID := otpResponseData.Data.AuthID
	if authID == "" {
		log.Fatal("Gagal mendapatkan auth_id dari respons.")
	}
	fmt.Printf("Berhasil mendapatkan auth_id: %s\n\n", authID)

	// =================================================================
	// LANGKAH 2: MEMINTA INPUT OTP DARI PENGGUNA
	// =================================================================
	fmt.Print(">>> OTP telah dikirim. Silakan masukkan 6 digit kode OTP: ")
	var otpCode string
	_, err = fmt.Scanln(&otpCode)
	if err != nil {
		log.Fatalf("Gagal membaca input OTP: %v", err)
	}

	// =================================================================
	// LANGKAH 3: VERIFIKASI OTP & DAPATKAN ACCESS TOKEN
	// =================================================================
	fmt.Println("\n--- LANGKAH 2: Memverifikasi OTP ---")

	// Siapkan payload untuk verifikasi OTP
	verifyPayload := VerifyOTPPayload{
		Phone:  targetPhone,
		AuthID: authID,
		OTP:    otpCode,
	}

	var verifyResponseData VerifyOTPResponse
	err = apiCall("/api/user/limited/xl/request-login.json", verifyPayload, &verifyResponseData)
	if err != nil {
		log.Fatalf("Gagal pada Langkah 2 (Verifikasi OTP): %v", err)
	}
	accessToken := verifyResponseData.Data.AccessToken
	if accessToken == "" {
		log.Fatal("Gagal mendapatkan access_token dari respons. OTP mungkin salah.")
	}
	fmt.Printf("Berhasil mendapatkan access_token: %s...\n\n", accessToken[:20]) // Tampilkan sebagian

	// =================================================================
	// LANGKAH 4: MENGAMBIL DAFTAR PAKET AKTIF
	// =================================================================
	fmt.Println("--- LANGKAH 3: Mengambil Daftar Paket Aktif ---")

	// Siapkan payload untuk mengambil daftar paket
	packagesPayload := GetPackagesPayload{AccessToken: accessToken}

	var packagesResponseData GetPackagesResponse
	err = apiCall("/api/user/limited/xl/package-active-list.json", packagesPayload, &packagesResponseData)
	if err != nil {
		log.Fatalf("Gagal pada Langkah 3 (Get Packages): %v", err)
	}

	fmt.Println("Berhasil mendapatkan daftar paket aktif!")

	// Simpan hasil akhir ke file JSON
	outputFilename := "hasil_akhir_transaksi.json"
	fileBytes, err := json.MarshalIndent(packagesResponseData, "", "  ")
	if err != nil {
		log.Fatal("Gagal format JSON untuk file:", err)
	}
	err = os.WriteFile(outputFilename, fileBytes, 0644)
	if err != nil {
		log.Fatal("Gagal menyimpan file JSON:", err)
	}

	fmt.Printf("\n>>> HASIL AKHIR TELAH DISIMPAN KE: %s <<<\n", outputFilename)
}

// apiCall adalah fungsi helper untuk melakukan POST request ke API
func apiCall(apiPath string, payload interface{}, responseStruct interface{}) error {
	url := baseURL + apiPath

	// Marshal payload struct ke JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("gagal marshal payload: %w", err)
	}

	// Buat request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("gagal membuat request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-token", authToken)
	req.Header.Set("User-Agent", "Go-Client/1.0")
	req.Header.Set("Referer", "https://putri-veronica.my.id/limited/tembak-paket-xl")

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirim request: %w", err)
	}
	defer resp.Body.Close()

	// Baca body respons
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("gagal membaca respons: %w", err)
	}

	// Cek status code
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code tidak valid: %d - %s", resp.StatusCode, string(body))
	}

	// Unmarshal body ke response struct
	err = json.Unmarshal(body, responseStruct)
	if err != nil {
		return fmt.Errorf("gagal unmarshal respons JSON: %w", err)
	}

	return nil
}
