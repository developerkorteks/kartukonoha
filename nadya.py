import requests
import json

# --- KONFIGURASI ---
AUTH_TOKEN = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiI2NzliNWRmZC0xZjMyLTQyYjQtYjhmMC01ZGIzNDJmOTE5MDciLCJ1bm0iOiJrb3J0a2VrcyIsIm5hbWUiOiJrb3J0a2VrcyIsIm1haWwiOiJrb3J0ZWtzbG9sQGdtYWlsLmNvbSIsInJvbGUiOiJtZW1iZXIiLCJncmF2YXRhciI6Imh0dHBzOi8vZ3JhdmF0YXIuY29tL2F2YXRhci81MWUxMTIzM2QwYjQ4YTZkOWU4YTM5YjFmNmQ2MWJkZSIsImhhc2giOiJ3PT1NNU02TlRaWElXMWlwdFozT1ROREEiLCJpYXQiOjE3NTYwNTM1MzMsImV4cCI6MTc1NjA4MjMzM30.FFtPvvJos5Wz-rcNNdIUj2plb_7RHW8baSVl299fw25A2ROZOC-q00MgcEJJBnJCq_KcskUF6AEnPMkVFv8cyieYWEeFELCHmBCLJi2jqkzQVkrPgR7DNmjbpzFnwKK_iMWFiej15o9TbK8gDNLL3b4S7zuU5PZBnraqiL1tm9Zx--K2kotNSOQ225ggI_7ca3BwpDjeKu1nF-iZMbMlgwd69gOkmonnIHHQeU8C2C2CQxfj3KazAv95Ho4_BT2yniptczzYI_xEg_Rk_yt51N1hdFd1LcLaZNJ1zhRCN9wnr6t_tKZRbstZdfWpfk03tdRaCKsNAZlpBMHC_CMBSA"

API_BASE_URL = "https://putri-veronica.my.id"
API_PRODUK_URL = f"{API_BASE_URL}/api/user/limited/xl/package-list-all.json"
API_SALDO_URL = f"{API_BASE_URL}/api/user/wallet/balance.json"

headers = {
    'User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36',
    'Content-Type': 'application/json',
    'x-token': AUTH_TOKEN, 
    'Referer': 'https://putri-veronica.my.id/limited/tembak-paket-xl'
}

def cek_saldo():
    """Fungsi untuk mengecek saldo."""
    print("Mengecek saldo...")
    try:
        response = requests.get(API_SALDO_URL, headers=headers)
        response.raise_for_status()
        data = response.json()
        
        print("\n--- INFO SALDO ---")
        print(json.dumps(data, indent=2))
        print("------------------\n")
        return data

    except requests.exceptions.RequestException as e:
        print(f"Error saat mengecek saldo: {e}")
        if "401" in str(e):
             print("Kemungkinan token sudah kedaluwarsa!")
        return None

def dapatkan_produk():
    """Fungsi untuk mendapatkan semua produk."""
    print("Mengambil daftar produk dari API...")
    try:
        response = requests.get(API_PRODUK_URL, headers=headers)
        response.raise_for_status()
        data = response.json()

        print(f"Berhasil mendapatkan {len(data.get('data', []))} produk.")
        return data.get('data', [])

    except requests.exceptions.RequestException as e:
        print(f"Error saat mengambil produk: {e}")
        if "401" in str(e):
             print("Kemungkinan token sudah kedaluwarsa!")
        return None

# --- FUNGSI UTAMA DIJALANKAN DI SINI ---
if __name__ == "__main__":
    cek_saldo()
    
    daftar_produk = dapatkan_produk()
    
    if daftar_produk:
        print("\n--- DAFTAR 5 PRODUK PERTAMA ---")
        for produk in daftar_produk[:5]: # Tampilkan 5 produk pertama saja
            # Gunakan nama key yang benar dari hasil "intipan" JSON
            nama = produk.get('package_name')
            harga = produk.get('package_harga_int') # Kita ambil harga integer agar mudah diolah
            kode = produk.get('package_code')
            
            print(f"- Kode: {kode}, Nama: {nama}, Harga: {harga}")
        print("----------------------------")
