#!/bin/bash

# Quick test script untuk WhatsApp Bot flow
# Untuk testing cepat tanpa interaksi manual

set -e

API_BASE="http://localhost:8080/api"
PHONE_NUMBER="087786388052"
PACKAGE_CODE="XL_MASTIF_30D_P_V1"
SOURCE="whatsapp_bot"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}=== Quick WhatsApp Bot Flow Test ===${NC}"

# Test 1: Request OTP
echo -e "${YELLOW}Testing OTP Request...${NC}"
OTP_RESPONSE=$(curl -s -X POST "${API_BASE}/otp/request" \
  -H "Content-Type: application/json" \
  -d "{\"phone_number\": \"${PHONE_NUMBER}\"}")

OTP_SUCCESS=$(echo "$OTP_RESPONSE" | jq -r '.success')
echo -e "OTP Request: $([ "$OTP_SUCCESS" == "true" ] && echo -e "${GREEN}✅ SUCCESS${NC}" || echo -e "${RED}❌ FAILED${NC}")"

if [ "$OTP_SUCCESS" == "true" ]; then
    CAN_RESEND=$(echo "$OTP_RESPONSE" | jq -r '.data.can_resend_in')
    echo -e "${BLUE}  → Can resend in: ${CAN_RESEND} seconds${NC}"
fi

echo ""

# Test 2: Check packages availability
echo -e "${YELLOW}Testing Package List...${NC}"
PACKAGES_RESPONSE=$(curl -s -X GET "${API_BASE}/packages?limit=3")
PACKAGES_SUCCESS=$(echo "$PACKAGES_RESPONSE" | jq -r '.success')
echo -e "Package List: $([ "$PACKAGES_SUCCESS" == "true" ] && echo -e "${GREEN}✅ SUCCESS${NC}" || echo -e "${RED}❌ FAILED${NC}")"

if [ "$PACKAGES_SUCCESS" == "true" ]; then
    PACKAGE_COUNT=$(echo "$PACKAGES_RESPONSE" | jq -r '.data | length')
    echo -e "${BLUE}  → Found ${PACKAGE_COUNT} packages${NC}"
    
    # Show first package as example
    FIRST_PACKAGE=$(echo "$PACKAGES_RESPONSE" | jq -r '.data[0].package_name // "N/A"')
    FIRST_PRICE=$(echo "$PACKAGES_RESPONSE" | jq -r '.data[0].package_harga_int // "N/A"')
    echo -e "${BLUE}  → Example: ${FIRST_PACKAGE} - Rp ${FIRST_PRICE}${NC}"
fi

echo ""

# Test 3: Check balance
echo -e "${YELLOW}Testing Balance Check...${NC}"
BALANCE_RESPONSE=$(curl -s -X GET "${API_BASE}/balance")
BALANCE_SUCCESS=$(echo "$BALANCE_RESPONSE" | jq -r '.success')
echo -e "Balance Check: $([ "$BALANCE_SUCCESS" == "true" ] && echo -e "${GREEN}✅ SUCCESS${NC}" || echo -e "${RED}❌ FAILED${NC}")"

if [ "$BALANCE_SUCCESS" == "true" ]; then
    BALANCE=$(echo "$BALANCE_RESPONSE" | jq -r '.data.balance // "N/A"')
    echo -e "${BLUE}  → Current Balance: Rp ${BALANCE}${NC}"
fi

echo ""

# Test 4: Check payment methods
echo -e "${YELLOW}Testing Payment Methods...${NC}"
PAYMENT_RESPONSE=$(curl -s -X GET "${API_BASE}/payment-methods")
PAYMENT_SUCCESS=$(echo "$PAYMENT_RESPONSE" | jq -r '.success')
echo -e "Payment Methods: $([ "$PAYMENT_SUCCESS" == "true" ] && echo -e "${GREEN}✅ SUCCESS${NC}" || echo -e "${RED}❌ FAILED${NC}")"

if [ "$PAYMENT_SUCCESS" == "true" ]; then
    METHODS=$(echo "$PAYMENT_RESPONSE" | jq -r '.data[].payment_method_display_name' | tr '\n' ', ' | sed 's/,$//')
    echo -e "${BLUE}  → Available: ${METHODS}${NC}"
fi

echo ""

# Test 5: Dashboard stats
echo -e "${YELLOW}Testing Dashboard Stats...${NC}"
DASHBOARD_RESPONSE=$(curl -s -X GET "${API_BASE}/dashboard")
DASHBOARD_SUCCESS=$(echo "$DASHBOARD_RESPONSE" | jq -r '.success')
echo -e "Dashboard: $([ "$DASHBOARD_SUCCESS" == "true" ] && echo -e "${GREEN}✅ SUCCESS${NC}" || echo -e "${RED}❌ FAILED${NC}")"

if [ "$DASHBOARD_SUCCESS" == "true" ]; then
    TOTAL_TXN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.total_transactions // 0')
    SUCCESS_RATE=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.success_rate // 0')
    echo -e "${BLUE}  → Total Transactions: ${TOTAL_TXN}${NC}"
    echo -e "${BLUE}  → Success Rate: ${SUCCESS_RATE}%${NC}"
    
    # Check WhatsApp bot stats
    WA_STATS=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.source_breakdown[] | select(.source == "whatsapp_bot")')
    if [ -n "$WA_STATS" ]; then
        WA_COUNT=$(echo "$WA_STATS" | jq -r '.count')
        WA_REVENUE=$(echo "$WA_STATS" | jq -r '.revenue')
        echo -e "${BLUE}  → WhatsApp Bot: ${WA_COUNT} transactions, Rp ${WA_REVENUE} revenue${NC}"
    else
        echo -e "${BLUE}  → WhatsApp Bot: No transactions yet${NC}"
    fi
fi

echo ""

# Test 6: Health check
echo -e "${YELLOW}Testing API Health...${NC}"
HEALTH_RESPONSE=$(curl -s -X GET "${API_BASE}/health")
HEALTH_SUCCESS=$(echo "$HEALTH_RESPONSE" | jq -r '.success')
echo -e "Health Check: $([ "$HEALTH_SUCCESS" == "true" ] && echo -e "${GREEN}✅ SUCCESS${NC}" || echo -e "${RED}❌ FAILED${NC}")"

if [ "$HEALTH_SUCCESS" == "true" ]; then
    STATUS=$(echo "$HEALTH_RESPONSE" | jq -r '.data.status // "N/A"')
    echo -e "${BLUE}  → API Status: ${STATUS}${NC}"
fi

echo ""
echo -e "${BLUE}=== Quick Test Summary ===${NC}"
echo -e "API Base URL: ${API_BASE}"
echo -e "Test Phone: ${PHONE_NUMBER}"
echo -e "Target Package: ${PACKAGE_CODE}"
echo -e "Source: ${SOURCE}"
echo ""
echo -e "${YELLOW}Note: This is a quick connectivity test.${NC}"
echo -e "${YELLOW}For full purchase flow test, run: ./test_whatsapp_flow.sh${NC}"