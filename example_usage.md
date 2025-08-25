# Nadia API v2.0 - Usage Examples

## Dashboard Monitoring

### Akses Dashboard
- **Dashboard**: http://localhost:8080/dashboard
- **API Flow Tester**: http://localhost:8080/api-flow
- **Swagger**: http://localhost:8080/swagger/index.html

## Complete Purchase Flow

### Step 1: Request OTP
```bash
curl -X POST "http://localhost:8080/api/otp/request" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052"
  }'
```

Response:
```json
{
  "status_code": 200,
  "message": "OTP sent successfully",
  "success": true,
  "data": {
    "phone_number": "087786388052",
    "can_resend_in": 60,
    "expires_in": 300
  }
}
```

### Step 2: Verify OTP & Get Access Token
```bash
curl -X POST "http://localhost:8080/api/otp/verify" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "otp_code": "123456"
  }'
```

Response:
```json
{
  "status_code": 200,
  "message": "OTP verified successfully",
  "success": true,
  "data": {
    "access_token": "1097690:ec321a89-6593-4b01-9a9e-383744301283",
    "phone_number": "087786388052"
  }
}
```

### Step 3: Purchase Package (with Source Tracking)

## API Endpoints dengan Source Tracking

### 1. Purchase Package dengan Source Tracking

```bash
curl -X POST "http://localhost:8080/api/purchase" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052",
    "package_code": "XL_MASTIF_30D_P_V1",
    "payment_method": "BALANCE",
    "access_token": "your_access_token_here",
    "source": "telegram_bot"
  }'
```

**Source Options:**
- `telegram_bot` - Transaksi dari Telegram Bot
- `whatsapp_bot` - Transaksi dari WhatsApp Bot  
- `web_interface` - Transaksi dari Web Interface
- `api_direct` - Transaksi langsung via API (default)

### 2. Dashboard Data

```bash
curl -X GET "http://localhost:8080/api/dashboard"
```

Response:
```json
{
  "status_code": 200,
  "message": "Dashboard data retrieved successfully",
  "success": true,
  "data": {
    "stats": {
      "total_transactions": 150,
      "successful_transactions": 142,
      "failed_transactions": 8,
      "total_revenue": 2500000,
      "success_rate": 94.7,
      "last_updated": "2024-01-15T10:30:00Z"
    },
    "recent_transactions": [...],
    "source_breakdown": [
      {
        "source": "telegram_bot",
        "count": 85,
        "revenue": 1500000,
        "success_rate": 95.3
      },
      {
        "source": "whatsapp_bot", 
        "count": 45,
        "revenue": 750000,
        "success_rate": 93.3
      }
    ],
    "balance": {...},
    "system_status": "OPERATIONAL"
  }
}
```

### 3. Transaction History dengan Filter

```bash
# Filter by source
curl -X GET "http://localhost:8080/api/transactions?source=telegram_bot&limit=20"

# Filter by status
curl -X GET "http://localhost:8080/api/transactions?status=SUCCESS&limit=50"

# Filter by date range
curl -X GET "http://localhost:8080/api/transactions?from=2024-01-01&to=2024-01-31"

# Combined filters
curl -X GET "http://localhost:8080/api/transactions?source=telegram_bot&status=SUCCESS&from=2024-01-15"
```

### 4. Export Transactions to CSV

```bash
# Export all transactions
curl -X GET "http://localhost:8080/api/export/transactions" -o transactions.csv

# Export with filters
curl -X GET "http://localhost:8080/api/export/transactions?source=telegram_bot&status=SUCCESS" -o telegram_success.csv
```

### 5. Statistics Endpoints

```bash
# System statistics
curl -X GET "http://localhost:8080/api/stats"

# Hourly statistics (last 24 hours)
curl -X GET "http://localhost:8080/api/stats/hourly"

# Daily statistics (last 30 days)
curl -X GET "http://localhost:8080/api/stats/daily"
```

## Bot Integration Examples

### Telegram Bot Integration

```python
import requests
import json

def purchase_package_telegram(phone, package_code, access_token):
    url = "http://localhost:8080/api/purchase"
    payload = {
        "phone_number": phone,
        "package_code": package_code,
        "payment_method": "BALANCE",
        "access_token": access_token,
        "source": "telegram_bot"  # Track source
    }
    
    response = requests.post(url, json=payload)
    return response.json()

# Usage
result = purchase_package_telegram("087786388052", "XL_MASTIF_30D_P_V1", "token_here")
print(f"Transaction Status: {result['success']}")
```

### WhatsApp Bot Integration

```javascript
const axios = require('axios');

async function purchasePackageWhatsApp(phone, packageCode, accessToken) {
    try {
        const response = await axios.post('http://localhost:8080/api/purchase', {
            phone_number: phone,
            package_code: packageCode,
            payment_method: 'BALANCE',
            access_token: accessToken,
            source: 'whatsapp_bot'  // Track source
        });
        
        return response.data;
    } catch (error) {
        console.error('Purchase failed:', error.response?.data || error.message);
        return null;
    }
}

// Usage
purchasePackageWhatsApp('087786388052', 'XL_MASTIF_30D_P_V1', 'token_here')
    .then(result => console.log('Result:', result));
```

## Monitoring Features

### 1. Real-time Dashboard
- **System Status**: OPERATIONAL/DEGRADED
- **Transaction Statistics**: Total, Success, Failed, Success Rate
- **Revenue Tracking**: Total revenue with breakdown by source
- **Source Analytics**: Performance comparison between different sources
- **Recent Transactions**: Latest transactions with status

### 2. Advanced Analytics
- **Hourly Statistics**: Transaction patterns by hour
- **Daily Statistics**: Transaction trends over 30 days
- **Source Performance**: Success rates by source (Telegram vs WhatsApp vs Direct API)
- **Error Tracking**: Failed transactions with error messages

### 3. Export & Reporting
- **CSV Export**: Full transaction history with filters
- **Date Range Filtering**: Custom date ranges for analysis
- **Source Filtering**: Analyze specific bot performance
- **Status Filtering**: Focus on successful or failed transactions

## Error Handling & JWT Auto-Refresh

The API automatically handles JWT token expiration:

1. **Automatic Retry**: If a request returns 401 Unauthorized, the system automatically refreshes the JWT token and retries the request
2. **Transparent to Users**: No manual token management required
3. **Fallback Handling**: Comprehensive error handling for all scenarios
4. **Transaction Tracking**: All transactions are tracked regardless of success/failure

## Dashboard Features

### Real-time Monitoring
- Auto-refresh every 30 seconds
- Live transaction status updates
- System health monitoring
- Balance tracking

### Interactive Features
- **Filter Transactions**: By status, source, date range
- **Export Data**: Download CSV reports
- **Visual Analytics**: Charts and graphs for better insights
- **Responsive Design**: Works on desktop and mobile

### Key Metrics Tracked
- Total transactions
- Success/failure rates
- Revenue by source
- Transaction patterns
- Error analysis
- Performance trends

## Security & Reliability

- **JWT Auto-refresh**: Automatic token management
- **Error Recovery**: Robust error handling with retries
- **Transaction Integrity**: All transactions are tracked and logged
- **Source Validation**: Proper source identification for analytics
- **Data Persistence**: In-memory storage with planned database integration