import requests
import json
import os

# --- KONFIGURASI ---
# Token ini didapatkan SETELAH Anda login secara manual di website
# dan menyalinnya dari Developer Tools (F12) -> Network -> Headers.
# PENTING: Token ini memiliki masa kedaluwarsa. Jika skrip gagal dengan error 401 atau 403,
# Anda perlu mendapatkan token yang baru.
AUTH_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJrb3J0a2VrcyIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJ3PT1NNU02TlRaWElXMWlwdFozT1ROREEiLCJpYXQiOjE3NTYwNTM1MzMsImV4cCI6MTc1NjA4MjMzM30.FFtPvvJos5Wz-rcNNdIUj2plb_7RHW8baSVl299fw25A2ROZOC-q00MgcEJJBnJCq_KcskUF6AEnPMkVFv8cyieYWEeFELCHmBCLJi2jqkzQVkrPgR7DNmjbpzFnwKK_iMWFiej15o9TbK8gDNLL3b4S7zuU5PZBnraqiL1tm9Zx--K2kotNSOQ225ggI_7ca3BwpDjeKu1nF-iZMbMlgwd69gOkmonnIHHQeU8C2C2CQxfj3KazAv95Ho4_BT2yniptczzYI_xEg_Rk_yt51N1hdFd1LcLaZNJ1zhRCN9wnr6t_tKZRbstZdfWpfk03tdRaCKsNAZlpBMHC_CMBSA"

API_BASE_URL = "https://putri-veronica.my.id"

# Siapkan header yang dibutuhkan untuk semua request
HEADERS = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36',
    'Content-Type': 'application/json',
    'x-token': AUTH_TOKEN, 
    'Referer': 'https://putri-veronica.my.id/limited/tembak-paket-xl'
}

# Daftar endpoint yang akan kita ambil, dan nama file untuk menyimpannya
ENDPOINTS = {
    "balance_response": "/api/user/wallet/balance.json",
    "payment_methods_response": "/api/user/wallet/payment-methods.json",
    "products_response": "/api/user/limited/xl/package-list-all.json"
}
# --------------------

def fetch_and_save_json():
    """Loop melalui setiap endpoint, ambil datanya, dan simpan ke file."""
    
    # Buat folder untuk menyimpan output jika belum ada
    output_folder = "api_responses"
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)
        
    for filename_prefix, path in ENDPOINTS.items():
        url = API_BASE_URL + path
        output_filename = os.path.join(output_folder, f"{filename_prefix}.json")
        
        print(f"Mengambil data dari: {url}...")
        
        try:
            # Lakukan GET request dengan header otentikasi
            response = requests.get(url, headers=HEADERS)
            
            # Cek jika ada error HTTP (seperti 401 Unauthorized, 404 Not Found, dll.)
            response.raise_for_status()
            
            # Ambil data JSON dari respons
            data = response.json()
            
            # Simpan data JSON ke file dengan format yang rapi (pretty-print)
            with open(output_filename, 'w', encoding='utf-8') as f:
                json.dump(data, f, indent=2, ensure_ascii=False)
            
            print(f"Berhasil! Data telah disimpan ke file: '{output_filename}'\n")

        except requests.exceptions.HTTPError as http_err:
            print(f"HTTP Error: {http_err}")
            if response.status_code == 401:
                print("Error 401 Unauthorized: Token Anda kemungkinan besar sudah kedaluwarsa. Harap perbarui AUTH_TOKEN.\n")
            else:
                print(f"Status Code: {response.status_code}\n")
        except requests.exceptions.RequestException as e:
            print(f"Error saat melakukan request ke {url}: {e}\n")
        except json.JSONDecodeError:
            print(f"Gagal mem-parsing JSON dari respons. Mungkin responsnya bukan JSON yang valid.\n")

# --- Jalankan Fungsi ---
if __name__ == "__main__":
    fetch_and_save_json()
    print("Semua proses selesai.")
