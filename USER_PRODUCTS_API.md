# User Products API Documentation

## Overview
Endpoint khusus untuk user dengan harga yang telah dimanipulasi (+1500 rupiah dari harga asli). Semua endpoint memerlukan authentication menggunakan API Key atau JWT token.

## Authentication
Endpoint user mendukung dua metode authentication:

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
- **Harga User**: Harga asli + Rp 1.500
- User **TIDAK** dapat melihat harga asli

## Endpoints

### 1. Get All Products
**GET** `/api/user/products`

Mengambil semua produk dengan harga yang dimanipulasi.

**Query Parameters:**
- `limit` (optional): Jumlah produk yang dikembalikan (default: 100)

**Example Request (API Key):**
```bash
curl -X GET "http://localhost:8080/api/user/products?limit=3" \
  -H "X-API-Key: nadia-admin-2024-secure-key"
```

**Example Request (JWT Token):**
```bash
curl -X GET "http://localhost:8080/api/user/products?limit=3" \
  -H "Authorization: Bearer <jwt_token>"
```

**Example Response:**
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
      "package_harga_int": 3500,
      "package_harga": "Rp. 3.500,00",
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
**POST** `/api/user/products/search`

Mencari dan memfilter produk dengan harga yang dimanipulasi.

**Request Body:**
```json
{
  "query": "masa aktif",
  "max_price": 7000,
  "min_price": 1000,
  "payment_method": "BALANCE"
}
```

**Example Request (API Key):**
```bash
curl -X POST "http://localhost:8080/api/user/products/search" \
  -H "Content-Type: application/json" \
  -H "X-API-Key: nadia-admin-2024-secure-key" \
  -d '{
    "query": "masa aktif",
    "max_price": 7000
  }'
```

**Example Request (JWT Token):**
```bash
curl -X POST "http://localhost:8080/api/user/products/search" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <jwt_token>" \
  -d '{
    "query": "masa aktif",
    "max_price": 7000
  }'
```

**Example Response:**
```json
{
  "statusCode": 200,
  "message": "Found 54 products matching your criteria",
  "success": true,
  "data": [
    {
      "package_code": "AXISMA1BLN_PREV",
      "package_name": "[RESMI] Tambah Masa Aktif Kartu AXIS 1 Bulan (30 Hari)",
      "package_harga_int": 6500,
      "package_harga": "Rp. 6.500,00"
    }
  ]
}
```

### 3. Get Product Stock
**GET** `/api/user/products/stock`

Mengambil informasi stok produk.

**Example Request (API Key):**
```bash
curl -X GET "http://localhost:8080/api/user/products/stock" \
  -H "X-API-Key: nadia-admin-2024-secure-key"
```

**Example Request (JWT Token):**
```bash
curl -X GET "http://localhost:8080/api/user/products/stock" \
  -H "Authorization: Bearer <jwt_token>"
```

**Example Response:**
```json
{
  "statusCode": 200,
  "message": "Berhasil mendapatkan data stok paket global!",
  "success": true,
  "data": [
    {
      "name": "Akrab Anggota Super MINI Promo",
      "package_id": "XLMINISUPERV2",
      "stok": 65
    }
  ]
}
```

## Price Comparison Examples

| Product | Admin Price (Original) | User Price (Manipulated) | Markup |
|---------|----------------------|-------------------------|---------|
| AKRAB_L75_PENGELOLA_PULSA | Rp. 2.000,00 | Rp. 3.500,00 | +Rp. 1.500 |
| AKRAB_L75_PENGELOLA_DANA | Rp. 2.000,00 | Rp. 3.500,00 | +Rp. 1.500 |
| AXISMA1BLN_PREV | Rp. 5.000,00 | Rp. 6.500,00 | +Rp. 1.500 |
| XL_MASTIF_30D_P_V1 | Rp. 1.000,00 | Rp. 2.500,00 | +Rp. 1.500 |

## Error Responses

### 401 Unauthorized - Missing Token
```json
{
  "statusCode": 401,
  "message": "Authorization header required",
  "success": false
}
```

### 401 Unauthorized - Invalid Token
```json
{
  "statusCode": 401,
  "message": "Invalid or expired token",
  "success": false
}
```

### 400 Bad Request - Invalid Search
```json
{
  "statusCode": 400,
  "message": "Invalid search request: ...",
  "success": false
}
```

### 500 Internal Server Error
```json
{
  "statusCode": 500,
  "message": "Failed to fetch products: ...",
  "success": false
}
```

## Security Features

1. **JWT Authentication**: Semua endpoint memerlukan valid JWT token
2. **Price Manipulation**: User tidak dapat melihat harga asli
3. **Security Logging**: Invalid token attempts dicatat dalam security log
4. **Rate Limiting**: Mengikuti daily limit dari Nadia API

## Implementation Details

- **Price Markup**: +Rp 1.500 ditambahkan ke semua harga
- **Price Formatting**: Format Indonesia dengan titik sebagai pemisah ribuan
- **Data Source**: Mengambil data dari Nadia API `/limited/xl/package-list-all.json`
- **Stock Data**: Mengambil dari `/limited/xl/check-stock-package-global.json`
- **Authentication**: JWT dengan HMAC-SHA256 signature
- **Token Expiry**: Configurable (default: 1 hour)

## Notes

- Harga yang ditampilkan ke user adalah harga yang sudah dimanipulasi
- User tidak memiliki akses ke endpoint admin yang menampilkan harga asli
- Semua filter price (min_price, max_price) menggunakan harga yang sudah dimanipulasi
- JWT token harus valid dan belum expired untuk mengakses endpoint