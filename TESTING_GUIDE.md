# 🧪 Testing Guide - WhatsApp Bot Flow

Panduan lengkap untuk testing API pembelian paket dengan source tracking WhatsApp Bot.

## 📋 Overview

API Nadia v2.0 mendukung source tracking untuk membedakan transaksi dari berbagai sumber:
- `whatsapp_bot` - Transaksi dari WhatsApp Bot
- `telegram_bot` - Transaksi dari Telegram Bot  
- `web_interface` - Transaksi dari Web Interface
- `api_direct` - Transaksi langsung via API (default)

## 🔄 Flow Pembelian Paket

### 1. Request OTP
```bash
curl -X POST "http://localhost:8080/api/otp/request" \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "087786388052"}'
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Berhasil mengirim kode OTP",
  "success": true,
  "data": {
    "phone_number": "087786388052",
    "can_resend_in": 60,
    "expires_in": 300
  }
}
```

### 2. Verify OTP & Get Access Token
```bash
curl -X POST "http://localhost:8080/api/otp/verify" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "otp_code": "123456"
  }'
```

**Response:**
```json
{
  "statusCode": 200,
  "message": "Berhasil Login OTP!",
  "success": true,
  "data": {
    "access_token": "1097690:ec321a89-6593-4b01-9a9e-383744301283",
    "phone_number": "087786388052"
  }
}
```

### 3. Purchase Package (dengan Source Tracking)
```bash
curl -X POST "http://localhost:8080/api/purchase" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "package_code": "XL_MASTIF_30D_P_V1",
    "payment_method": "BALANCE",
    "access_token": "1097690:ec321a89-6593-4b01-9a9e-383744301283",
    "source": "whatsapp_bot"
  }'
```

**Response (Success):**
```json
{
  "statusCode": 200,
  "message": "Paket berhasil dibeli",
  "success": true,
  "data": {
    "msisdn": "6287786388052",
    "package_code": "XL_MASTIF_30D_P_V1",
    "package_name": "XL Masa Aktif 30 Hari",
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

## 🛠️ Test Scripts

### 1. Full WhatsApp Flow Test
```bash
./test_whatsapp_flow.sh
```
- ✅ Complete purchase flow dengan OTP manual input
- ✅ Source tracking untuk WhatsApp Bot
- ✅ Transaction status checking
- ✅ Dashboard stats verification

### 2. Quick Connectivity Test
```bash
./test_quick_whatsapp.sh
```
- ✅ API connectivity check
- ✅ Package availability
- ✅ Balance check
- ✅ Payment methods
- ✅ Dashboard stats
- ✅ Health check

### 3. Source Comparison Test
```bash
./test_source_comparison.sh
```
- ✅ Multiple source testing
- ✅ Package search functionality
- ✅ Dashboard analytics
- ✅ Source breakdown comparison

## 📊 Monitoring & Analytics

### Dashboard Stats
```bash
curl -X GET "http://localhost:8080/api/dashboard"
```

**Response includes:**
- Total transactions
- Success/failure rates
- Revenue by source
- Recent transactions
- Source breakdown

### Source Breakdown Example
```json
{
  "source_breakdown": [
    {
      "source": "whatsapp_bot",
      "count": 25,
      "revenue": 150000,
      "success_rate": 96.0
    },
    {
      "source": "telegram_bot",
      "count": 18,
      "revenue": 95000,
      "success_rate": 94.4
    }
  ]
}
```

## 🔍 Package Search

### Search by Price Range
```bash
curl -X POST "http://localhost:8080/api/packages/search" \
  -H "Content-Type: application/json" \
  -d '{
    "max_price": 5000,
    "payment_method": "BALANCE"
  }'
```

### Search by Keyword
```bash
curl -X POST "http://localhost:8080/api/packages/search" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "masa aktif",
    "min_price": 1000,
    "max_price": 10000
  }'
```

## 🚨 Error Scenarios

### 1. Invalid OTP
```json
{
  "statusCode": 400,
  "message": "Kode OTP salah atau sudah kedaluwarsa",
  "success": false
}
```

### 2. Insufficient Balance
```json
{
  "statusCode": 400,
  "message": "Saldo tidak mencukupi untuk membeli paket ini",
  "success": false
}
```

### 3. Invalid Access Token
```json
{
  "statusCode": 400,
  "message": "Access token tidak valid atau sudah kedaluwarsa",
  "success": false
}
```

## 📱 WhatsApp Bot Integration

### Node.js Example
```javascript
const axios = require('axios');

class NadiaWhatsAppBot {
  constructor(apiBase = 'http://localhost:8080/api') {
    this.apiBase = apiBase;
  }

  async requestOTP(phoneNumber) {
    try {
      const response = await axios.post(`${this.apiBase}/otp/request`, {
        phone_number: phoneNumber
      });
      return response.data;
    } catch (error) {
      throw new Error(`OTP request failed: ${error.response?.data?.message || error.message}`);
    }
  }

  async verifyOTP(phoneNumber, otpCode) {
    try {
      const response = await axios.post(`${this.apiBase}/otp/verify`, {
        phone_number: phoneNumber,
        otp_code: otpCode
      });
      return response.data;
    } catch (error) {
      throw new Error(`OTP verification failed: ${error.response?.data?.message || error.message}`);
    }
  }

  async purchasePackage(phoneNumber, packageCode, accessToken, paymentMethod = 'BALANCE') {
    try {
      const response = await axios.post(`${this.apiBase}/purchase`, {
        phone_number: phoneNumber,
        package_code: packageCode,
        payment_method: paymentMethod,
        access_token: accessToken,
        source: 'whatsapp_bot'  // Important: Source tracking
      });
      return response.data;
    } catch (error) {
      throw new Error(`Purchase failed: ${error.response?.data?.message || error.message}`);
    }
  }

  async getPackages(limit = 50) {
    try {
      const response = await axios.get(`${this.apiBase}/packages?limit=${limit}`);
      return response.data;
    } catch (error) {
      throw new Error(`Failed to get packages: ${error.response?.data?.message || error.message}`);
    }
  }

  async searchPackages(filters) {
    try {
      const response = await axios.post(`${this.apiBase}/packages/search`, filters);
      return response.data;
    } catch (error) {
      throw new Error(`Package search failed: ${error.response?.data?.message || error.message}`);
    }
  }
}

// Usage Example
const bot = new NadiaWhatsAppBot();

async function handlePurchaseFlow(phoneNumber, packageCode) {
  try {
    // Step 1: Request OTP
    console.log('Requesting OTP...');
    const otpResult = await bot.requestOTP(phoneNumber);
    console.log('OTP sent successfully');

    // Step 2: Wait for user to provide OTP (in real bot, this would be from user message)
    const otpCode = '123456'; // This would come from user input

    // Step 3: Verify OTP
    console.log('Verifying OTP...');
    const verifyResult = await bot.verifyOTP(phoneNumber, otpCode);
    const accessToken = verifyResult.data.access_token;
    console.log('OTP verified, access token obtained');

    // Step 4: Purchase package
    console.log('Purchasing package...');
    const purchaseResult = await bot.purchasePackage(phoneNumber, packageCode, accessToken);
    
    if (purchaseResult.success) {
      console.log('Package purchased successfully!');
      console.log(`Transaction ID: ${purchaseResult.data.trx_id}`);
      return purchaseResult;
    } else {
      throw new Error(purchaseResult.message);
    }
  } catch (error) {
    console.error('Purchase flow failed:', error.message);
    throw error;
  }
}
```

## 🔧 Testing Checklist

### Pre-Testing
- [ ] API server running on port 8080
- [ ] Valid phone number for testing
- [ ] Sufficient balance in account
- [ ] `jq` installed for JSON parsing

### Basic Tests
- [ ] Health check endpoint
- [ ] Package list retrieval
- [ ] Balance check
- [ ] Payment methods check

### OTP Flow Tests
- [ ] OTP request successful
- [ ] OTP session management
- [ ] OTP verification
- [ ] Access token generation

### Purchase Flow Tests
- [ ] Package purchase with WhatsApp source
- [ ] Transaction tracking
- [ ] Source attribution
- [ ] Dashboard stats update

### Error Handling Tests
- [ ] Invalid OTP handling
- [ ] Expired session handling
- [ ] Insufficient balance handling
- [ ] Invalid access token handling

## 📈 Performance Monitoring

### Key Metrics to Track
- **Success Rate**: Percentage of successful transactions
- **Response Time**: API response times
- **Source Performance**: Success rates by source
- **Error Patterns**: Common failure reasons
- **Revenue Tracking**: Revenue by source

### Dashboard URLs
- **Main Dashboard**: http://localhost:8080/dashboard
- **API Documentation**: http://localhost:8080/swagger/index.html
- **API Flow Tester**: http://localhost:8080/api-flow

## 🚀 Production Considerations

### Security
- Implement rate limiting for OTP requests
- Add request validation and sanitization
- Use HTTPS in production
- Implement proper authentication for dashboard

### Monitoring
- Set up logging for all transactions
- Monitor success rates by source
- Track API performance metrics
- Set up alerts for high failure rates

### Scalability
- Consider database storage for transaction history
- Implement caching for package data
- Add load balancing for high traffic
- Monitor memory usage for in-memory storage

## 📞 Support

Jika ada masalah dengan testing:
1. Check API server status: `curl http://localhost:8080/api/health`
2. Verify phone number format
3. Check balance availability
4. Review error messages in response
5. Check dashboard for transaction history