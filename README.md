# 🚀 Nadia Package Purchase API

API untuk pembelian paket XL dengan JWT authentication otomatis dan dukungan OTP verification.

## ✨ Features

- ✅ **JWT Authentication otomatis** - Token di-manage secara otomatis
- ✅ **Auto token refresh** - Token di-refresh otomatis ketika expired
- ✅ **OTP Support** - Dukungan penuh untuk OTP verification
- ✅ **Swagger Documentation** - Dokumentasi API yang lengkap
- ✅ **CORS Support** - Siap untuk frontend integration
- ✅ **Error Handling** - Error handling yang comprehensive
- ✅ **Consistent Response Format** - Format response yang konsisten

## 🚀 Quick Start

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Run Server

```bash
go run api_server.go
```

Server akan berjalan di `http://localhost:8080`

### 3. Akses Dokumentasi

Buka browser dan akses: `http://localhost:8080/swagger/index.html`

## 📋 API Endpoints

### System
- `GET /api/health` - Health check

### Packages
- `GET /api/packages` - Get all available packages

### Wallet
- `GET /api/balance` - Get account balance
- `GET /api/payment-methods` - Get available payment methods

### Purchase
- `POST /api/purchase` - Purchase package (tanpa OTP)
- `POST /api/purchase/otp` - Purchase package dengan OTP

### OTP
- `POST /api/otp/request` - Request OTP code

### Transaction
- `POST /api/transaction/check` - Check transaction status

## 🔐 Authentication Flow

API menggunakan JWT authentication yang di-handle secara otomatis. Anda tidak perlu mengirim token authentication di request.

## 📱 OTP Flow

### 1. Request OTP

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
    "auth_id": "b56a2b90-25b5-4831-a504-970b24331d9c",
    "can_resend_in": 60
  }
}
```

### 2. Purchase dengan OTP

Setelah menerima OTP via SMS, gunakan `auth_id` dan kode OTP untuk membuat `access_token`:

Format access_token: `{otp_code}:{auth_id}`

Contoh: jika OTP code = `1097498` dan auth_id = `7c17952a-343a-49b1-9e2e-4a52b7682b7a`, maka access_token = `1097498:7c17952a-343a-49b1-9e2e-4a52b7682b7a`

```bash
curl -X POST http://localhost:8080/api/purchase/otp \
  -H "Content-Type: application/json" \
  -d '{
    "package_code": "b0012edd21983678eb7ebc08d8f04ecd",
    "phone_number": "6287817739901",
    "access_token": "1097498:7c17952a-343a-49b1-9e2e-4a52b7682b7a",
    "payment_method": "BALANCE"
  }'
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Paket berhasil dibeli. Silakan cek kuotanya via aplikasi MyXL",
  "success": true,
  "data": {
    "msisdn": "6287817739901",
    "package_code": "b0012edd21983678eb7ebc08d8f04ecd",
    "package_name": "Masa Aktif Kartu XL 1 Tahun",
    "package_processing_fee": 5000,
    "trx_id": "a1e3b046-3dda-4eb5-93c2-f0eb01a1a453",
    "have_deeplink": false,
    "deeplink_data": {
      "payment_method": "BALANCE",
      "deeplink_url": ""
    },
    "is_qris": false,
    "qris_data": []
  }
}
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

## 📖 Usage Examples

### Get Available Packages

```bash
curl -X GET http://localhost:8080/api/packages
```

### Check Account Balance

```bash
curl -X GET http://localhost:8080/api/balance
```

### Check Transaction Status

```bash
curl -X POST http://localhost:8080/api/transaction/check \
  -H "Content-Type: application/json" \
  -d '{
    "transaction_id": "a1e3b046-3dda-4eb5-93c2-f0eb01a1a453"
  }'
```

### Get Payment Methods

```bash
curl -X GET http://localhost:8080/api/payment-methods
```

## 🔧 Configuration

Edit konstanta di `api_server.go` untuk konfigurasi:

```go
const (
    ssoApiBaseURL = "https://sso.putri-veronica.my.id/api"
    nadiaApiURL   = "https://putri-veronica.my.id/api/user"
    ssoStaticKey  = "your_static_key"
    username      = "your_username"
    password      = "your_password"
)
```

## 🛠️ Development

### Generate Swagger Documentation

```bash
~/go/bin/swag init -g api_server.go
```

### Run in Production Mode

```bash
export GIN_MODE=release
go run api_server.go
```

## 📚 API Documentation

Dokumentasi lengkap tersedia di Swagger UI: `http://localhost:8080/swagger/index.html`

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