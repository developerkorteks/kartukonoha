package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// --- KONFIGURASI ---
// GANTI DENGAN TOKEN JWT ANDA YANG VALID. JANGAN DIBIARKAN KOSONG.
const (
	authToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJuYWJpbCIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJRPT1jME42TmpaWElXMWlwdFp6T0ROakkiLCJpYXQiOjE3NTYxMTMxMDgsImV4cCI6MTc1NjE0MTkwOH0.FPBQWs8srQm95CXsrAEg9EGT9HXGTChbBAPDXNW48vX7ie1XWNSKcprqqM89uJPGGb2IxQ08bTPVdrJ9nzePNXeXaSbUa5-j7QOk0B_y6d7QmTXv6L6yrtH7-kJlZtIO5apwxQegqyq7VSUmY0Hmm0f0rz9j_8d2SKkFzBJbLmn-9yaRzRCUT5EFg90tJmPdmweEaK_3KGvVKHtYW_AhjYD9sCic2i6EnnSADZpuwydPNyBHL0WJ6q4sLc_efqWBHIe7b273CYWaWFSH_ZDPRAqsBbvi1w8yF24N-DD9B-enhDLY6WQixhExlsJNo30de7RU3hVAP3YgdHxmu_ZRjg"
	baseURL   = "https://putri-veronica.my.id"
)

// --------------------

// --- DEFINISI STRUCT UNTUK SEMUA API ---

// Universal
type GenericResponse map[string]interface{}

// API Metode Pembayaran (untuk di dalam Produk)
type PaymentMethod struct {
	Method      string `json:"payment_method"`
	DisplayName string `json:"payment_method_display_name"`
}

// API Produk (Struktur diperbarui)
type Product struct {
	PackageCode             string          `json:"package_code"`
	PackageName             string          `json:"package_name"`
	PackageHargaInt         int             `json:"package_harga_int"`
	AvailablePaymentMethods []PaymentMethod `json:"available_payment_methods"`
}
type ProductAPIResponse struct {
	Data []Product `json:"data"`
}

// API Transaksi
type RequestOTPPayload struct {
	Phone string `json:"phone"`
}
type RequestOTPResponse struct {
	Data struct {
		AuthID string `json:"auth_id"`
	} `json:"data"`
}
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
type AccessTokenPayload struct {
	AccessToken string `json:"access_token"`
}
type BuyPackagePayload struct {
	PackageCode   string `json:"package_code"`
	Phone         string `json:"phone"`
	AccessToken   string `json:"access_token"`
	PaymentMethod string `json:"payment_method"`
}

// =================================================================
// FUNGSI UTAMA (MAIN LOOP)
// =================================================================
func main() {
	if authToken == "TOKEN_JWT_ANDA_YANG_VALID" || authToken == "" {
		log.Fatal("Error: Harap ganti nilai `authToken` di dalam kode dengan token JWT Anda yang valid.")
	}

	reader := bufio.NewReader(os.Stdin)

	// Loop menu utama
	for {
		fmt.Println("\n==============================================")
		fmt.Println("             MENU UTAMA APLIKASI")
		fmt.Println("==============================================")
		fmt.Println("1. Cek Saldo Akun")
		fmt.Println("2. Tampilkan Semua Produk XL")
		fmt.Println("3. Cari Produk Berdasarkan Metode Pembayaran")
		fmt.Println("4. Cek Paket Aktif di Nomor XL")
		fmt.Println("5. Lakukan Transaksi Tembak Paket XL")
		fmt.Println("0. Keluar")
		fmt.Println("==============================================")
		fmt.Print("Masukkan pilihan Anda: ")

		choiceStr, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceStr)

		switch choice {
		case "1":
			cekSaldo()
		case "2":
			tampilkanSemuaProduk()
		case "3":
			cariProdukByPayment()
		case "4":
			cekPaketAktif()
		case "5":
			lakukanTransaksi()
		case "0":
			fmt.Println("Terima kasih! Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

// =================================================================
// FUNGSI-FUNGSI UNTUK SETIAP MENU
// =================================================================

func cekSaldo() {
	fmt.Println("\n--- Menu: Cek Saldo ---")
	var responseData GenericResponse
	err := apiCall("GET", "/api/user/wallet/balance.json", nil, &responseData)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	prettyJSON, _ := json.MarshalIndent(responseData, "", "  ")
	fmt.Println("Hasil Cek Saldo:")
	fmt.Println(string(prettyJSON))
}

func tampilkanSemuaProduk() {
	fmt.Println("\n--- Menu: Daftar Semua Produk & Metode Pembayaran ---")
	var responseData ProductAPIResponse
	err := apiCall("GET", "/api/user/limited/xl/package-list-all.json", nil, &responseData)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Berhasil mendapatkan %d produk:\n", len(responseData.Data))
	for i, produk := range responseData.Data {
		fmt.Printf("\n%d. Nama: %s\n", i+1, produk.PackageName)
		fmt.Printf("   Kode: %s\n", produk.PackageCode)
		fmt.Printf("   Harga: %d\n", produk.PackageHargaInt)
		fmt.Print("   Metode Bayar: ")
		for j, metode := range produk.AvailablePaymentMethods {
			fmt.Print(metode.Method)
			if j < len(produk.AvailablePaymentMethods)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println()
	}
}

func cariProdukByPayment() {
	fmt.Println("\n--- Menu: Cari Produk Berdasarkan Metode Pembayaran ---")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan metode pembayaran (BALANCE, QRIS, DANA): ")
	method, _ := reader.ReadString('\n')
	method = strings.ToUpper(strings.TrimSpace(method))

	var responseData ProductAPIResponse
	err := apiCall("GET", "/api/user/limited/xl/package-list-all.json", nil, &responseData)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("\nMenampilkan produk yang mendukung metode '%s':\n", method)
	found := false
	for _, produk := range responseData.Data {
		for _, pm := range produk.AvailablePaymentMethods {
			if pm.Method == method {
				fmt.Printf("\n- Nama: %s\n", produk.PackageName)
				fmt.Printf("  Kode: %s\n", produk.PackageCode)
				fmt.Printf("  Harga: %d\n", produk.PackageHargaInt)
				found = true
				break
			}
		}
	}
	if !found {
		fmt.Println("Tidak ada produk yang ditemukan dengan metode pembayaran tersebut.")
	}
}

func cekPaketAktif() {
	fmt.Println("\n--- Menu: Cek Paket Aktif di Nomor XL ---")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nomor telepon target: ")
	targetPhone, _ := reader.ReadString('\n')
	targetPhone = strings.TrimSpace(targetPhone)

	fmt.Println("\nLangkah 1: Meminta OTP...")
	authID := requestOTP(targetPhone)
	fmt.Printf("Berhasil mendapatkan auth_id.\n")

	fmt.Print("\n>>> OTP telah dikirim, silakan masukkan kode: ")
	otpCode, _ := reader.ReadString('\n')
	otpCode = strings.TrimSpace(otpCode)

	fmt.Println("\nLangkah 2: Memverifikasi OTP...")
	accessToken := verifyOTP(targetPhone, authID, otpCode)
	fmt.Printf("Berhasil mendapatkan access_token.\n")

	fmt.Println("\nLangkah 3: Mengambil Daftar Paket Aktif...")
	payload := AccessTokenPayload{AccessToken: accessToken}
	var responseData GenericResponse
	err := apiCall("POST", "/api/user/limited/xl/package-active-list.json", payload, &responseData)

	outputFilename := "paket_aktif_response.json"
	saveResponseToFile(outputFilename, responseData)

	if err != nil {
		log.Printf("Gagal Cek Paket Aktif: %v", err)
		log.Printf("Respons error tersimpan di %s", outputFilename)
	} else {
		fmt.Println("Berhasil mendapatkan daftar paket aktif!")
		fmt.Printf("Respons lengkap telah disimpan di: %s\n", outputFilename)
	}
}

func lakukanTransaksi() {
	fmt.Println("\n--- Menu: Lakukan Transaksi Tembak Paket XL ---")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan nomor telepon target: ")
	targetPhone, _ := reader.ReadString('\n')
	targetPhone = strings.TrimSpace(targetPhone)

	fmt.Print("Masukkan kode paket yang ingin dibeli: ")
	packageCode, _ := reader.ReadString('\n')
	packageCode = strings.TrimSpace(packageCode)

	// BARIS BARU: Meminta input metode pembayaran
	fmt.Print("Masukkan metode pembayaran (BALANCE, DANA, QRIS): ")
	paymentMethod, _ := reader.ReadString('\n')
	paymentMethod = strings.ToUpper(strings.TrimSpace(paymentMethod))

	fmt.Println("\nLangkah 1: Meminta OTP...")
	authID := requestOTP(targetPhone)
	fmt.Printf("Berhasil mendapatkan auth_id.\n")

	fmt.Print("\n>>> OTP telah dikirim, silakan masukkan kode: ")
	otpCode, _ := reader.ReadString('\n')
	otpCode = strings.TrimSpace(otpCode)

	fmt.Println("\nLangkah 2: Memverifikasi OTP...")
	accessToken := verifyOTP(targetPhone, authID, otpCode)
	fmt.Printf("Berhasil mendapatkan access_token.\n")

	fmt.Println("\nLangkah 3: Mengirim permintaan pembelian paket...")
	phoneForPayload := targetPhone
	if strings.HasPrefix(phoneForPayload, "0") {
		phoneForPayload = "62" + phoneForPayload[1:]
	}
	payload := BuyPackagePayload{
		PackageCode:   packageCode,
		Phone:         phoneForPayload,
		AccessToken:   accessToken,
		PaymentMethod: paymentMethod, // PERUBAHAN: Menggunakan input dari pengguna
	}
	var responseData GenericResponse
	err := apiCall("POST", "/api/user/limited/xl/beli-paket-otp.json", payload, &responseData)

	outputFilename := "hasil_transaksi.json"
	saveResponseToFile(outputFilename, responseData)

	if err != nil {
		log.Printf("Gagal melakukan transaksi: %v", err)
		log.Printf("Respons error tersimpan di %s", outputFilename)
		return
	}

	// Cek 'success' flag di dalam JSON respons
	if success, ok := responseData["success"].(bool); ok && success {
		fmt.Println("TRANSAKSI BERHASIL!")
		fmt.Printf("Respons lengkap dari server telah disimpan di file: %s\n", outputFilename)
	} else {
		message := "Pesan error tidak diketahui."
		if msg, ok := responseData["message"].(string); ok {
			message = msg
		}
		fmt.Printf("TRANSAKSI GAGAL: %s\n", message)
		fmt.Printf("Respons lengkap dari server telah disimpan di file: %s\n", outputFilename)
	}
}

// =================================================================
// FUNGSI HELPER UNTUK LOGIKA API (YANG DIPAKAI BERSAMA)
// =================================================================

func requestOTP(phone string) string {
	payload := RequestOTPPayload{Phone: phone}
	var responseData RequestOTPResponse
	err := apiCall("POST", "/api/user/limited/xl/request-otp.json", payload, &responseData)
	if err != nil {
		log.Fatalf("Gagal pada Langkah 1 (Request OTP): %v", err)
	}
	if responseData.Data.AuthID == "" {
		log.Fatalf("Gagal mendapatkan auth_id dari respons. Cek kembali nomor telepon atau token.")
	}
	return responseData.Data.AuthID
}

func verifyOTP(phone, authID, otp string) string {
	payload := VerifyOTPPayload{Phone: phone, AuthID: authID, OTP: otp}
	var responseData VerifyOTPResponse
	err := apiCall("POST", "/api/user/limited/xl/request-login.json", payload, &responseData)
	if err != nil {
		log.Fatalf("Gagal pada Langkah 2 (Verifikasi OTP): %v. Mungkin OTP salah.", err)
	}
	if responseData.Data.AccessToken == "" {
		log.Fatalf("Gagal mendapatkan access_token dari respons. OTP mungkin salah.")
	}
	return responseData.Data.AccessToken
}

func apiCall(method, apiPath string, payload interface{}, responseStruct interface{}) error {
	url := baseURL + apiPath
	var reqBody io.Reader = nil
	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("gagal marshal payload: %w", err)
		}
		reqBody = bytes.NewBuffer(payloadBytes)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return fmt.Errorf("gagal membuat request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-token", authToken)
	req.Header.Set("User-Agent", "Go-Client/1.0")
	req.Header.Set("Referer", "https://putri-veronica.my.id/")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal mengirim request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("gagal membaca respons: %w", err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code tidak valid: %d - %s", resp.StatusCode, string(body))
	}
	err = json.Unmarshal(body, responseStruct)
	if err != nil {
		return fmt.Errorf("gagal unmarshal respons JSON: %w. Body: %s", err, string(body))
	}
	return nil
}

func saveResponseToFile(filename string, data interface{}) {
	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Gagal format JSON untuk file %s: %v", filename, err)
		return
	}
	_ = os.WriteFile(filename, fileBytes, 0644)
}
