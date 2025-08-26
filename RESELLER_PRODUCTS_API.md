# Reseller Products API Documentation

## Overview
Endpoint khusus untuk reseller dengan harga yang telah dimanipulasi (+500 rupiah dari harga asli). Semua endpoint memerlukan authentication menggunakan API Key atau JWT token.

## Authentication
Endpoint reseller mendukung dua metode authentication:

### 1. API Key Authentication
```
X-API-Key: nadia-admin-2024-secure-key
```

### 2. JWT Bearer Token Authentication
```
Authorization: Bearer <jwt_token>
```

**Note**: Anda dapat menggunakan salah satu dari kedua metode authentication di atas.

## Price Manipulation
- **Harga Admin (asli)**: Harga sebenarnya dari Nadia API
- **Harga Reseller**: Harga asli + Rp 500
- Reseller **TIDAK** dapat melihat harga asli

## Endpoints

### 1. Get All Products

**GET** `/api/reseller/products`

Mengambil semua produk yang tersedia dengan harga yang telah dimanipulasi untuk reseller.

**Query Parameters:**
- `limit` (optional): Jumlah produk yang dikembalikan (default: 100)

**Example Request (API Key):**
```bash
curl -X GET "http://localhost:8080/api/reseller/products?limit=3" \
  -H "X-API-Key: nadia-admin-2024-secure-key"
```

**Example Request (JWT Token):**
```bash
curl -X GET "http://localhost:8080/api/reseller/products?limit=3" \
  -H "Authorization: Bearer <jwt_token>"
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Retrieved 3 products",
  "success": true,
  "data": [
    {
      "package_code": "AKRAB_L75_PENGELOLA_PULSA",
      "package_name": "[Metode Pulsa] Pengelola Akrab L Kuber 75GB dan Bonus Paket Akrab untuk 3 Anggota 28 Hari",
      "package_name_alias_short": "Pengelola Akrab L Kuber 75GB dan Bonus Paket Akrab untuk 3 Anggota 28 Hari",
      "package_description": "Buat yang tau-tau saja hehe...",
      "package_harga_int": 4000,
      "package_harga": "Rp. 4.000,00",
      "have_daily_limit": true,
      "daily_limit_details": {
        "max_daily_transaction_limit": 2000,
        "current_daily_transaction_count": 0
      },
      "no_need_login": false,
      "can_multi_trx": true,
      "can_scheduled_trx": true,
      "have_cut_off_time": false,
      "cut_off_time": {
        "prohibited_hour_starttime": "23:00",
        "prohibited_hour_endtime": "02:00"
      },
      "need_check_stock": false,
      "is_show_payment_method": true,
      "available_payment_methods": [
        {
          "order": 1,
          "payment_method": "BALANCE",
          "payment_method_display_name": "Metode Pulsa (BALANCE)",
          "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa"
        }
      ]
    }
  ]
}
```

### 2. Search Products

**POST** `/api/reseller/products/search`

Mencari dan memfilter produk berdasarkan nama, harga, metode pembayaran, dll.

**Request Body:**
```json
{
  "query": "masa aktif",
  "min_price": 1000,
  "max_price": 7000,
  "payment_method": "BALANCE"
}
```

**Example Request (API Key):**
```bash
curl -X POST "http://localhost:8080/api/reseller/products/search" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: nadia-admin-2024-secure-key" \
  -d '{
    "query": "masa aktif",
    "max_price": 7000
  }'
```

**Example Request (JWT Token):**
```bash
curl -X POST "http://localhost:8080/api/reseller/products/search" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt_token>" \
  -d '{
    "query": "masa aktif",
    "max_price": 7000
  }'
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Found 54 products matching your criteria",
  "success": true,
  "data": [
    {
      "package_code": "AXISMA1BLN_PREV",
      "package_name": "[RESMI] Tambah Masa Aktif Kartu AXIS 1 Bulan (30 Hari)",
      "package_name_alias_short": "Tambah Masa Aktif Kartu AXIS 1 Bulan",
      "package_description": "MISI.. NUMPANG JUALAN MASA AKTIF AXIS :v...",
      "package_harga_int": 7000,
      "package_harga": "Rp. 7.000,00",
      "have_daily_limit": true,
      "daily_limit_details": {
        "max_daily_transaction_limit": 2000,
        "current_daily_transaction_count": 0
      },
      "no_need_login": true,
      "can_multi_trx": false,
      "can_scheduled_trx": false,
      "have_cut_off_time": false,
      "cut_off_time": {
        "prohibited_hour_starttime": "00:26",
        "prohibited_hour_endtime": "23:36"
      },
      "need_check_stock": false,
      "is_show_payment_method": false,
      "available_payment_methods": [
        {
          "order": 1,
          "payment_method": "BALANCE",
          "payment_method_display_name": "Metode Pulsa (BALANCE)",
          "desc": "Langsung memotong pulsa kamu jika paket membutuhkan Pulsa"
        }
      ]
    }
  ]
}
```

### 3. Get Product Stock

**GET** `/api/reseller/products/stock`

Mengambil informasi stok produk.

**Example Request (API Key):**
```bash
curl -X GET "http://localhost:8080/api/reseller/products/stock" \
  -H "X-API-Key: nadia-admin-2024-secure-key"
```

**Example Request (JWT Token):**
```bash
curl -X GET "http://localhost:8080/api/reseller/products/stock" \
  -H "Authorization: Bearer <jwt_token>"
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Stock information retrieved successfully",
  "success": true,
  "data": {
    // Stock information from Nadia API
  }
}
```

## Search Parameters

### Available Search Filters:
- `query` (string): Pencarian berdasarkan nama atau deskripsi produk
- `min_price` (integer): Harga minimum (dalam rupiah)
- `max_price` (integer): Harga maksimum (dalam rupiah)
- `payment_method` (string): Filter berdasarkan metode pembayaran (BALANCE, DANA, QRIS, dll)

### Payment Methods:
- `BALANCE`: Metode Pulsa
- `DANA`: E-Wallet DANA
- `QRIS`: Pembayaran QRIS
- Dan metode pembayaran lainnya yang tersedia

## Error Responses

### 401 Unauthorized
```json
{
  "statusCode": 401,
  "message": "Unauthorized: Invalid API key or token",
  "success": false
}
```

### 400 Bad Request
```json
{
  "statusCode": 400,
  "message": "Invalid search request: [error details]",
  "success": false
}
```

### 500 Internal Server Error
```json
{
  "statusCode": 500,
  "message": "Failed to fetch products: [error details]",
  "success": false
}
```

## Notes

1. **Price Manipulation**: Semua harga yang ditampilkan sudah ditambah Rp 500 dari harga asli Nadia API
2. **Authentication**: Gunakan salah satu dari API Key atau JWT token
3. **Rate Limiting**: Endpoint ini mungkin memiliki rate limiting, gunakan dengan bijak
4. **Data Freshness**: Data produk diambil real-time dari Nadia API
5. **Stock Information**: Informasi stok diambil langsung dari Nadia API tanpa manipulasi