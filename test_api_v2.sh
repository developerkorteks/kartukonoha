#!/bin/bash

# Nadia API v2.0 Test Script
# Usage: ./test_api_v2.sh

API_BASE="http://localhost:8080/api"

echo "🚀 Testing Nadia Package Purchase API v2.0"
echo "=========================================="

# Test 1: Health Check
echo ""
echo "1️⃣  Testing Health Check..."
curl -s -X GET "$API_BASE/health" | jq '.'

# Test 2: Get All Packages (GET)
echo ""
echo "2️⃣  Testing Get All Packages (GET with limit)..."
curl -s -X GET "$API_BASE/packages?limit=10" | jq '.message, .success, (.data | length)'

# Test 3: Search Packages with Filters (POST)
echo ""
echo "3️⃣  Testing Package Search (POST with filters)..."
curl -s -X POST "$API_BASE/packages/search" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "masa aktif",
    "max_price": 10000,
    "payment_method": "BALANCE"
  }' | jq '.message, .success, (.data | length)'

# Test 4: Get Package Stock
echo ""
echo "4️⃣  Testing Package Stock..."
curl -s -X GET "$API_BASE/packages/stock" | jq '.message, .success, (.data | length)'

# Test 5: Get Balance
echo ""
echo "5️⃣  Testing Get Balance..."
curl -s -X GET "$API_BASE/balance" | jq '.message, .success'

# Test 6: Get Payment Methods
echo ""
echo "6️⃣  Testing Get Payment Methods..."
curl -s -X GET "$API_BASE/payment-methods" | jq '.message, .success'

# Test 7: Get Invoices
echo ""
echo "7️⃣  Testing Get Invoices..."
curl -s -X GET "$API_BASE/invoices?limit=5" | jq '.message, .success'

# Test 8: Request OTP (Example - will fail without valid XL number)
echo ""
echo "8️⃣  Testing Request OTP (Example - will fail without valid XL number)..."
curl -s -X POST "$API_BASE/otp/request" \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "087786388052"
  }' | jq '.message, .success'

echo ""
echo "✅ API v2.0 Testing Complete!"
echo ""
echo "📖 Key Improvements in v2.0:"
echo "   🎯 Simplified OTP flow - no manual auth_id management"
echo "   📱 Automatic phone number formatting"
echo "   🔄 Session management with auto-cleanup"
echo "   🔍 Package search and filtering"
echo "   🏥 Complete XL card management"
echo ""
echo "📖 For full documentation, visit: http://localhost:8080/swagger/index.html"
echo ""
echo "🔄 Simplified Purchase Flow:"
echo "1. POST /api/otp/request - Request OTP"
echo "2. POST /api/purchase - Purchase with OTP code (auth_id handled automatically)"
echo ""
echo "📦 Package Endpoints:"
echo "• GET /api/packages?limit=50 - Get all packages (simple, fast)"
echo "• POST /api/packages/search - Search with filters (advanced, flexible)"
echo ""
echo "🏥 Card Management Flow:"
echo "1. POST /api/otp/request - Request OTP"
echo "2. POST /api/card/status - Check card status with OTP"
echo "3. POST /api/card/packages - Check active packages with OTP"
echo ""
echo "🔍 Package Search Examples:"
echo "   • Search by name: {\"query\": \"masa aktif\"}"
echo "   • Filter by price: {\"max_price\": 10000, \"min_price\": 1000}"
echo "   • Filter by payment: {\"payment_method\": \"BALANCE\"}"
echo "   • Combined filters: {\"query\": \"combo\", \"payment_method\": \"DANA\", \"max_price\": 5000}"