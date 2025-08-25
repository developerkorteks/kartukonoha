#!/bin/bash

# Nadia API Test Script
# Usage: ./test_api.sh

API_BASE="http://localhost:8080/api"

echo "🚀 Testing Nadia Package Purchase API"
echo "======================================"

# Test 1: Health Check
echo ""
echo "1️⃣  Testing Health Check..."
curl -s -X GET "$API_BASE/health" | jq '.'

# Test 2: Get Packages
echo ""
echo "2️⃣  Testing Get Packages..."
echo "Note: This will take a moment as it requires JWT authentication..."
curl -s -X GET "$API_BASE/packages" | jq '.message, .success'

# Test 3: Get Balance
echo ""
echo "3️⃣  Testing Get Balance..."
curl -s -X GET "$API_BASE/balance" | jq '.message, .success, .data.balance // empty'

# Test 4: Get Payment Methods
echo ""
echo "4️⃣  Testing Get Payment Methods..."
curl -s -X GET "$API_BASE/payment-methods" | jq '.message, .success'

# Test 5: Request OTP (Example)
echo ""
echo "5️⃣  Testing Request OTP (Example - will fail without valid XL number)..."
curl -s -X POST "$API_BASE/otp/request" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052"
  }' | jq '.message, .success'

echo ""
echo "✅ API Testing Complete!"
echo ""
echo "📖 For full documentation, visit: http://localhost:8080/swagger/index.html"
echo ""
echo "🔄 OTP Flow Example:"
echo "1. Request OTP: POST /api/otp/request"
echo "2. Get OTP code from SMS"
echo "3. Purchase with OTP: POST /api/purchase/otp"
echo "   - access_token format: {otp_code}:{auth_id}"
echo "4. Check transaction: POST /api/transaction/check"