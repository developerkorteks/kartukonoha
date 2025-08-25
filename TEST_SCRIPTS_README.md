# 🧪 Test Scripts untuk WhatsApp Bot Flow

Kumpulan script testing untuk API Nadia v2.0 dengan fokus pada source tracking WhatsApp Bot.

## 📁 Available Test Scripts

### 1. `test_whatsapp_flow.sh` - Complete Purchase Flow
**Fungsi**: Test lengkap flow pembelian paket dari WhatsApp Bot
**Fitur**:
- ✅ Request OTP dengan input manual
- ✅ Verify OTP dan dapatkan access token
- ✅ Purchase package dengan source tracking
- ✅ Check transaction status
- ✅ Verify dashboard stats

**Usage**:
```bash
./test_whatsapp_flow.sh
```

**Output**: Interactive flow dengan input OTP manual

---

### 2. `test_quick_whatsapp.sh` - Quick Connectivity Test
**Fungsi**: Test cepat konektivitas API tanpa purchase
**Fitur**:
- ✅ API health check
- ✅ Package list availability
- ✅ Balance check
- ✅ Payment methods
- ✅ Dashboard stats
- ✅ WhatsApp bot transaction history

**Usage**:
```bash
./test_quick_whatsapp.sh
```

**Output**: Summary report semua endpoint

---

### 3. `test_source_comparison.sh` - Source Tracking Test
**Fungsi**: Test dan compare berbagai source tracking
**Fitur**:
- ✅ Test OTP untuk multiple sources
- ✅ Simulate purchase untuk berbagai source
- ✅ Package search functionality
- ✅ Dashboard analytics comparison
- ✅ Source breakdown analysis

**Usage**:
```bash
./test_source_comparison.sh
```

**Output**: Comparison report antar source

---

### 4. `manual_test_steps.sh` - Step-by-Step Manual Testing
**Fungsi**: Test manual per step untuk debugging
**Fitur**:
- ✅ Individual step testing
- ✅ Interactive mode
- ✅ Token management
- ✅ Detailed curl commands
- ✅ Step-by-step guidance

**Usage**:
```bash
# Run specific step
./manual_test_steps.sh 1        # Request OTP
./manual_test_steps.sh otp      # Request OTP
./manual_test_steps.sh verify   # Verify OTP
./manual_test_steps.sh purchase # Purchase package

# Run all steps interactively
./manual_test_steps.sh all

# Show help
./manual_test_steps.sh help
```

**Available Steps**:
1. `otp` - Request OTP
2. `verify` - Verify OTP
3. `packages` - List packages
4. `search` - Search packages
5. `balance` - Check balance
6. `purchase` - Purchase package
7. `dashboard` - Dashboard stats
8. `health` - Health check

---

## 🚀 Quick Start

### 1. Pastikan API Server Running
```bash
# Check if server is running
curl -s http://localhost:8080/api/health | jq '.'
```

### 2. Run Quick Test
```bash
./test_quick_whatsapp.sh
```

### 3. Run Full Purchase Flow
```bash
./test_whatsapp_flow.sh
```

### 4. Manual Step Testing
```bash
./manual_test_steps.sh all
```

## 📊 Expected Results

### Successful OTP Request
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

### Successful Purchase
```json
{
  "statusCode": 200,
  "message": "Paket berhasil dibeli",
  "success": true,
  "data": {
    "trx_id": "a1e3b046-3dda-4eb5-93c2-f0eb01a1a453",
    "package_name": "XL Masa Aktif 30 Hari",
    "package_processing_fee": 5000
  }
}
```

### Dashboard with WhatsApp Bot Stats
```json
{
  "source_breakdown": [
    {
      "source": "whatsapp_bot",
      "count": 5,
      "revenue": 25000,
      "success_rate": 100.0
    }
  ]
}
```

## 🔧 Configuration

### Environment Variables (Optional)
```bash
export NADIA_API_BASE="http://localhost:8080/api"
export NADIA_TEST_PHONE="087786388052"
export NADIA_TEST_PACKAGE="XL_MASTIF_30D_P_V1"
```

### Script Configuration
Edit variables di awal setiap script:
```bash
API_BASE="http://localhost:8080/api"
PHONE_NUMBER="087786388052"
PACKAGE_CODE="XL_MASTIF_30D_P_V1"
SOURCE="whatsapp_bot"
```

## 🚨 Troubleshooting

### Common Issues

#### 1. API Server Not Running
```bash
# Error: Connection refused
curl: (7) Failed to connect to localhost port 8080

# Solution: Start the API server
go run nadia_api.go
```

#### 2. Invalid Phone Number
```bash
# Error: Invalid phone number format
{
  "statusCode": 400,
  "message": "Invalid phone number format",
  "success": false
}

# Solution: Use valid Indonesian phone number
# Format: 087786388052 (without +62)
```

#### 3. OTP Expired
```bash
# Error: OTP session expired
{
  "statusCode": 400,
  "message": "OTP session expired. Please request a new OTP.",
  "success": false
}

# Solution: Request new OTP
./manual_test_steps.sh otp
```

#### 4. Insufficient Balance
```bash
# Error: Insufficient balance
{
  "statusCode": 400,
  "message": "Saldo tidak mencukupi",
  "success": false
}

# Solution: Check balance first
./manual_test_steps.sh balance
```

### Debug Mode
Add `-x` flag untuk debug:
```bash
bash -x ./test_whatsapp_flow.sh
```

## 📈 Monitoring Results

### Dashboard URLs
- **Main Dashboard**: http://localhost:8080/dashboard
- **API Docs**: http://localhost:8080/swagger/index.html
- **API Flow Tester**: http://localhost:8080/api-flow

### Key Metrics to Monitor
- **Success Rate**: Should be > 95%
- **Response Time**: Should be < 2 seconds
- **Source Attribution**: WhatsApp bot transactions properly tracked
- **Error Patterns**: Monitor common failure reasons

### Export Results
```bash
# Export transaction history
curl -X GET "http://localhost:8080/api/export/transactions?source=whatsapp_bot" -o whatsapp_transactions.csv

# Export dashboard data
curl -X GET "http://localhost:8080/api/dashboard" | jq '.' > dashboard_stats.json
```

## 🔄 Continuous Testing

### Automated Testing (Cron)
```bash
# Add to crontab for hourly health checks
0 * * * * /path/to/test_quick_whatsapp.sh >> /var/log/nadia_health.log 2>&1
```

### Load Testing
```bash
# Run multiple concurrent tests
for i in {1..5}; do
  ./test_quick_whatsapp.sh &
done
wait
```

## 📝 Test Reports

### Generate Test Report
```bash
# Create test report
{
  echo "=== Nadia API Test Report ==="
  echo "Date: $(date)"
  echo ""
  ./test_quick_whatsapp.sh
  echo ""
  echo "=== Dashboard Stats ==="
  curl -s http://localhost:8080/api/dashboard | jq '.data.stats'
} > test_report_$(date +%Y%m%d_%H%M%S).txt
```

### Log Analysis
```bash
# Analyze success rates
grep "SUCCESS" test_report_*.txt | wc -l
grep "FAILED" test_report_*.txt | wc -l
```

## 🎯 Best Practices

### Before Testing
1. ✅ Ensure API server is running
2. ✅ Check balance availability
3. ✅ Verify phone number format
4. ✅ Install required tools (`jq`, `curl`)

### During Testing
1. ✅ Monitor response times
2. ✅ Check error messages carefully
3. ✅ Verify source tracking
4. ✅ Save access tokens for reuse

### After Testing
1. ✅ Review dashboard stats
2. ✅ Check transaction history
3. ✅ Analyze error patterns
4. ✅ Document any issues

## 📞 Support

Jika mengalami masalah:
1. Check API health: `./manual_test_steps.sh health`
2. Review error messages dalam response JSON
3. Check server logs untuk detail error
4. Verify konfigurasi phone number dan package code
5. Pastikan balance mencukupi untuk testing

## 🔗 Related Files

- `TESTING_GUIDE.md` - Comprehensive testing documentation
- `PACKAGES_API_GUIDE.md` - Package API documentation
- `example_usage.md` - API usage examples
- `nadia_api.go` - Main API implementation