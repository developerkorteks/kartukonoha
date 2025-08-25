package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

// Definisikan struct untuk menampung data produk agar lebih mudah diolah.
// Nama field (misal: PackageCode) harus diawali huruf besar agar bisa diakses (exported).
// `json:"..."` adalah tag yang memberitahu Go cara memetakan field JSON ke field struct.
type Product struct {
	PackageCode        string `json:"package_code"`
	PackageName        string `json:"package_name"`
	PackageHargaInt    int    `json:"package_harga_int"`
	PackageDescription string `json:"package_description"`
}

// Struct untuk menampung seluruh respons dari API produk
type ProductAPIResponse struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Success    bool      `json:"success"`
	Data       []Product `json:"data"`
}

func main() {
	// --- KONFIGURASI ---
	// Token ini didapatkan SETELAH Anda login dan menyalinnya dari Developer Tools.
	// PENTING: Token ini memiliki masa kedaluwarsa!
	authToken := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJrb3J0a2VrcyIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJ3PT1NNU02TlRaWElXMWlwdFozT1ROREEiLCJpYXQiOjE3NTYwNTM1MzMsImV4cCI6MTc1NjA4MjMzM30.FFtPvvJos5Wz-rcNNdIUj2plb_7RHW8baSVl299fw25A2ROZOC-q00MgcEJJBnJCq_KcskUF6AEnPMkVFv8cyieYWEeFELCHmBCLJi2jqkzQVkrPgR7DNmjbpzFnwKK_iMWFiej15o9TbK8gDNLL3b4S7zuU5PZBnraqiL1tm9Zx--K2kotNSOQ225ggI_7ca3BwpDjeKu1nF-iZMbMlgwd69gOkmonnIHHQeU8C2C2CQxfj3KazAv95Ho4_BT2yniptczzYI_xEg_Rk_yt51N1hdFd1LcLaZNJ1zhRCN9wnr6t_tKZRbstZdfWpfk03tdRaCKsNAZlpBMHC_CMBSA"

	apiProdukURL := "https://putri-veronica.my.id/api/user/limited/xl/package-list-all.json"
	// --------------------

	// Inisialisasi Colly Collector
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36"),
	)

	// Callback ini akan dijalankan SEBELUM setiap request dibuat.
	// Di sinilah kita menambahkan header otentikasi.
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Mengunjungi API Endpoint:", r.URL.String())
		r.Headers.Set("x-token", authToken)
		r.Headers.Set("Referer", "https://putri-veronica.my.id/limited/tembak-paket-xl")
		r.Headers.Set("Content-Type", "application/json")
	})

	// Callback ini akan dijalankan SETELAH respons diterima.
	// Di sinilah kita memproses data JSON.
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Respons diterima dengan status code:", r.StatusCode)

		// Cek jika status code sukses (200 OK)
		if r.StatusCode != 200 {
			log.Println("Gagal mendapatkan data, status code:", r.StatusCode)
			if r.StatusCode == 401 {
				log.Println("Error 401: Token otentikasi kemungkinan tidak valid atau sudah kedaluwarsa.")
			}
			return
		}

		// Buat variabel untuk menampung hasil unmarshal (parsing) JSON
		var apiResponse ProductAPIResponse

		// Unmarshal body respons (yang berupa byte) ke dalam struct kita
		err := json.Unmarshal(r.Body, &apiResponse)
		if err != nil {
			log.Fatal("Gagal mem-parsing JSON:", err)
		}

		// Sekarang data sudah ada di dalam struct dan mudah diakses
		fmt.Printf("Pesan dari API: %s\n", apiResponse.Message)
		fmt.Printf("Berhasil mendapatkan %d produk.\n", len(apiResponse.Data))

		// Simpan hasil ke file JSON
		outputFilename := "products_response.json"
		file, err := json.MarshalIndent(apiResponse, "", "  ")
		if err != nil {
			log.Fatal("Gagal membuat format JSON untuk file:", err)
		}

		err = os.WriteFile(outputFilename, file, 0644)
		if err != nil {
			log.Fatal("Gagal menyimpan file JSON:", err)
		}
		fmt.Printf("Data lengkap telah disimpan di file: %s\n", outputFilename)

		// Contoh menampilkan 5 produk pertama dari data yang sudah diparsing
		fmt.Println("\n--- DAFTAR 5 PRODUK PERTAMA ---")
		for i, produk := range apiResponse.Data {
			if i >= 5 {
				break
			}
			fmt.Printf("- Nama: %s, Harga: %d, Kode: %s\n", produk.PackageName, produk.PackageHargaInt, produk.PackageCode)
		}
		fmt.Println("-----------------------------")
	})

	// Callback untuk menangani error
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Terjadi error:", err)
	})

	// Mulai proses dengan mengunjungi URL API produk
	c.Visit(apiProdukURL)
}
