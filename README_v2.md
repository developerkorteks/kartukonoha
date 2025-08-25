# 🚀 Nadia Package Purchase API v2.0

API yang disederhanakan untuk pembelian paket XL dengan flow yang lebih mudah dan user-friendly.

## ✨ New Features v2.0

- ✅ **Simplified OTP Flow** - Tidak perlu manage auth_id manual
- ✅ **Automatic Phone Number Formatting** - Support berbagai format nomor
- ✅ **Session Management** - OTP session otomatis dengan auto-cleanup
- ✅ **Package Search & Filtering** - Cari paket berdasarkan nama, harga, metode bayar
- ✅ **Complete XL Card Management** - Cek status kartu dan paket aktif
- ✅ **User-friendly Request/Response** - Format yang lebih mudah dipahami
- ✅ **Enhanced Swagger Documentation** - Dokumentasi yang lebih lengkap

## 🚀 Quick Start

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Run Server

```bash
go run nadia_api.go
```

Server akan berjalan di `http://localhost:8080`

### 3. Akses Dokumentasi

Buka browser dan akses: `http://localhost:8080/swagger/index.html`

## 📋 API Endpoints

### System
- `GET /api/health` - Health check

### Packages
- `POST /api/packages` - Get packages with optional filtering
- `GET /api/packages/stock` - Get package stock information

### OTP & Purchase (Simplified Flow)
- `POST /api/otp/request` - Request OTP (step 1)
- `POST /api/purchase` - Purchase package with OTP (step 2)

### Card Management
- `POST /api/card/status` - Check XL card status and balance
- `POST /api/card/packages` - Check active packages on card

### Wallet
- `GET /api/balance` - Get account balance
- `GET /api/payment-methods` - Get available payment methods

### Transaction
- `POST /api/transaction/check` - Check transaction status

### Invoice
- `GET /api/invoices` - Get invoice list with pagination

## 🎯 Simplified Purchase Flow

### Step 1: Request OTP

```bash
curl -X POST http://localhost:8080/api/otp/request \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052"
  }'
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Berhasil mengirim kode OTP, silakan cek SMS di nomor XL kamu!",
  "success": true,
  "data": {
    "phone_number": "087786388052",
    "can_resend_in": 60,
    "expires_in": 300
  }
}
```

### Step 2: Purchase Package

```bash
curl -X POST http://localhost:8080/api/purchase \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "package_code": "XL_MASTIF_30D_P_V1",
    "payment_method": "BALANCE",
    "otp_code": "123456"
  }'
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Paket berhasil dibeli. Silakan cek kuotanya via aplikasi MyXL",
  "success": true,
  "data": {
    "msisdn": "6287786388052",
    "package_code": "XL_MASTIF_30D_P_V1",
    "package_name": "XL Masa Aktif 30 Hari",
    "package_processing_fee": 1000,
    "trx_id": "a1e3b046-3dda-4eb5-93c2-f0eb01a1a453"
  }
}
```

## 🔍 Package Search & Filtering

### Search Packages

```bash
curl -X POST http://localhost:8080/api/packages \
  -H "Content-Type: application/json" \
  -d '{
    "query": "masa aktif",
    "payment_method": "BALANCE",
    "max_price": 10000,
    "min_price": 1000
  }'
```

**Available Filters:**
- `query` - Search by package name or description
- `payment_method` - Filter by payment method (BALANCE, DANA, QRIS)
- `max_price` - Maximum price filter
- `min_price` - Minimum price filter

## 📱 Phone Number Format Support

API mendukung berbagai format nomor telepon:

```json
{
  "phone_number": "087786388052"    // ✅ Format lokal
}
{
  "phone_number": "6287786388052"   // ✅ Dengan kode negara
}
{
  "phone_number": "+6287786388052"  // ✅ Dengan + dan kode negara
}
{
  "phone_number": "0877-8638-8052" // ✅ Dengan separator
}
```

Semua format akan dinormalisasi secara otomatis.

## 🏥 XL Card Management

### Check Card Status

```bash
curl -X POST http://localhost:8080/api/card/status \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "otp_code": "123456"
  }'
```

### Check Active Packages

```bash
curl -X POST http://localhost:8080/api/card/packages \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "otp_code": "123456"
  }'
```

## 💰 Wallet & Payment

### Check Balance

```bash
curl -X GET http://localhost:8080/api/balance
```

### Get Payment Methods

```bash
curl -X GET http://localhost:8080/api/payment-methods
```

## 📊 Transaction Management

### Check Transaction Status

```bash
curl -X POST http://localhost:8080/api/transaction/check \
  -H "Content-Type: application/json" \
  -d '{
    "transaction_id": "a1e3b046-3dda-4eb5-93c2-f0eb01a1a453"
  }'
```

## 🧾 Invoice Management

### Get Invoices

```bash
curl -X GET "http://localhost:8080/api/invoices?search=topup&limit=10"
```

## 📝 Response Format

Semua response menggunakan format JSON yang konsisten:

```json
{
  "statusCode": 200,
  "message": "Success message",
  "success": true,
  "data": {
    // Response data here
  }
}
```

## 🚨 Error Handling

Jika terjadi error, response akan berformat:

```json
{
  "statusCode": 400,
  "message": "Error description",
  "success": false
}
```

## 🔧 Key Improvements dari v1.0

### 1. Simplified OTP Flow
**v1.0 (Manual):**
```json
// Step 1: Request OTP
{"phone_number": "087786388052"}
// Response: {"auth_id": "abc123"}

// Step 2: Purchase
{
  "package_code": "XL_MASTIF_30D_P_V1",
  "phone_number": "6287786388052",
  "access_token": "123456:abc123",  // Manual format
  "payment_method": "BALANCE"
}
```

**v2.0 (Simplified):**
```json
// Step 1: Request OTP
{"phone_number": "087786388052"}

// Step 2: Purchase (auth_id handled automatically)
{
  "phone_number": "087786388052",
  "package_code": "XL_MASTIF_30D_P_V1",
  "payment_method": "BALANCE",
  "otp_code": "123456"
}
```

### 2. Automatic Phone Number Formatting
- Support berbagai format nomor telepon
- Normalisasi otomatis
- Tidak perlu manual format ke 62xxx

### 3. Session Management
- OTP session disimpan otomatis selama 5 menit
- Auto-cleanup setelah purchase berhasil
- Tidak perlu manage auth_id manual

### 4. Enhanced Package Search
- Filter berdasarkan nama, harga, metode bayar
- Search dalam deskripsi paket
- Response yang lebih terstruktur

### 5. Complete Card Management
- Cek status kartu dengan OTP
- Lihat paket aktif
- Informasi pulsa dan masa aktif

## 🛠️ Development

### Generate Swagger Documentation

```bash
~/go/bin/swag init -g nadia_api.go
```

### Run in Production Mode

```bash
export GIN_MODE=release
go run nadia_api.go
```

## 📚 API Documentation

Dokumentasi lengkap tersedia di Swagger UI: `http://localhost:8080/swagger/index.html`

## 🎯 Use Cases

### 1. Simple Package Purchase
```bash
# 1. Request OTP
curl -X POST http://localhost:8080/api/otp/request \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "087786388052"}'

# 2. Purchase (after receiving OTP via SMS)
curl -X POST http://localhost:8080/api/purchase \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "package_code": "XL_MASTIF_30D_P_V1",
    "payment_method": "BALANCE",
    "otp_code": "123456"
  }'
```

### 2. Find Cheap Packages
```bash
curl -X POST http://localhost:8080/api/packages \
  -H "Content-Type: application/json" \
  -d '{
    "max_price": 5000,
    "payment_method": "BALANCE"
  }'
```

### 3. Check Card Status
```bash
# 1. Request OTP
curl -X POST http://localhost:8080/api/otp/request \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "087786388052"}'

# 2. Check status
curl -X POST http://localhost:8080/api/card/status \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "otp_code": "123456"
  }'
```

## 🤝 Contributing

1. Fork repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## 📄 License

MIT License

---

**Server running on:** [http://localhost:8080](http://localhost:8080)  
**Swagger Documentation:** [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

**Version:** 2.0.0  
**Last Updated:** August 2025