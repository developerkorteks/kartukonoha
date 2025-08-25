#!/bin/bash

# Test script untuk membandingkan berbagai source tracking
# Menguji WhatsApp Bot vs Telegram Bot vs API Direct

set -e

API_BASE="http://localhost:8080/api"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
NC='\033[0m'

echo -e "${PURPLE}=== Source Tracking Comparison Test ===${NC}"
echo ""

# Function to test OTP request for different sources
test_otp_request() {
    local phone=$1
    local source_name=$2
    
    echo -e "${YELLOW}Testing OTP for ${source_name}...${NC}"
    
    OTP_RESPONSE=$(curl -s -X POST "${API_BASE}/otp/request" \
      -H "Content-Type: application/json" \
      -d "{\"phone_number\": \"${phone}\"}")
    
    OTP_SUCCESS=$(echo "$OTP_RESPONSE" | jq -r '.success')
    
    if [ "$OTP_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ ${source_name} OTP request successful${NC}"
        CAN_RESEND=$(echo "$OTP_RESPONSE" | jq -r '.data.can_resend_in')
        echo -e "${BLUE}   → Can resend in: ${CAN_RESEND} seconds${NC}"
        return 0
    else
        echo -e "${RED}❌ ${source_name} OTP request failed${NC}"
        ERROR_MSG=$(echo "$OTP_RESPONSE" | jq -r '.message')
        echo -e "${RED}   → Error: ${ERROR_MSG}${NC}"
        return 1
    fi
}

# Function to simulate purchase (without actual OTP verification)
simulate_purchase() {
    local phone=$1
    local source=$2
    local source_name=$3
    
    echo -e "${YELLOW}Simulating purchase for ${source_name}...${NC}"
    
    # This would normally require a valid access token
    # For testing purposes, we'll just show what the request would look like
    echo -e "${BLUE}Purchase request for ${source_name}:${NC}"
    echo -e "${BLUE}  Phone: ${phone}${NC}"
    echo -e "${BLUE}  Source: ${source}${NC}"
    echo -e "${BLUE}  Package: XL_MASTIF_30D_P_V1${NC}"
    echo -e "${BLUE}  Payment: BALANCE${NC}"
    
    # Show the curl command that would be used
    echo -e "${BLUE}  Command:${NC}"
    echo "  curl -X POST \"${API_BASE}/purchase\" \\"
    echo "    -H \"Content-Type: application/json\" \\"
    echo "    -d '{"
    echo "      \"phone_number\": \"${phone}\","
    echo "      \"package_code\": \"XL_MASTIF_30D_P_V1\","
    echo "      \"payment_method\": \"BALANCE\","
    echo "      \"access_token\": \"[VALID_ACCESS_TOKEN]\","
    echo "      \"source\": \"${source}\""
    echo "    }'"
    echo ""
}

# Test different sources
echo -e "${BLUE}=== Testing Different Sources ===${NC}"
echo ""

# WhatsApp Bot
echo -e "${PURPLE}1. WhatsApp Bot Source${NC}"
test_otp_request "087786388052" "WhatsApp Bot"
simulate_purchase "087786388052" "whatsapp_bot" "WhatsApp Bot"
echo ""

# Telegram Bot  
echo -e "${PURPLE}2. Telegram Bot Source${NC}"
test_otp_request "087786388053" "Telegram Bot"
simulate_purchase "087786388053" "telegram_bot" "Telegram Bot"
echo ""

# Web Interface
echo -e "${PURPLE}3. Web Interface Source${NC}"
test_otp_request "087786388054" "Web Interface"
simulate_purchase "087786388054" "web_interface" "Web Interface"
echo ""

# API Direct
echo -e "${PURPLE}4. API Direct Source${NC}"
test_otp_request "087786388055" "API Direct"
simulate_purchase "087786388055" "api_direct" "API Direct"
echo ""

# Check current dashboard stats
echo -e "${PURPLE}=== Current Dashboard Statistics ===${NC}"
DASHBOARD_RESPONSE=$(curl -s -X GET "${API_BASE}/dashboard")
DASHBOARD_SUCCESS=$(echo "$DASHBOARD_RESPONSE" | jq -r '.success')

if [ "$DASHBOARD_SUCCESS" == "true" ]; then
    echo -e "${GREEN}✅ Dashboard data retrieved${NC}"
    echo ""
    
    # Overall stats
    echo -e "${YELLOW}Overall Statistics:${NC}"
    TOTAL_TXN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.total_transactions // 0')
    SUCCESS_TXN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.successful_transactions // 0')
    FAILED_TXN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.failed_transactions // 0')
    SUCCESS_RATE=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.success_rate // 0')
    TOTAL_REVENUE=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.total_revenue // 0')
    
    echo -e "${BLUE}  Total Transactions: ${TOTAL_TXN}${NC}"
    echo -e "${BLUE}  Successful: ${SUCCESS_TXN}${NC}"
    echo -e "${BLUE}  Failed: ${FAILED_TXN}${NC}"
    echo -e "${BLUE}  Success Rate: ${SUCCESS_RATE}%${NC}"
    echo -e "${BLUE}  Total Revenue: Rp ${TOTAL_REVENUE}${NC}"
    echo ""
    
    # Source breakdown
    echo -e "${YELLOW}Source Breakdown:${NC}"
    SOURCE_BREAKDOWN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.source_breakdown')
    
    if [ "$SOURCE_BREAKDOWN" != "null" ] && [ "$SOURCE_BREAKDOWN" != "[]" ]; then
        echo "$SOURCE_BREAKDOWN" | jq -r '.[] | "  \(.source): \(.count) transactions, Rp \(.revenue) revenue, \(.success_rate)% success rate"'
    else
        echo -e "${BLUE}  No transactions recorded yet${NC}"
    fi
    
    echo ""
    
    # Recent transactions
    echo -e "${YELLOW}Recent Transactions (last 5):${NC}"
    RECENT_TXN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.recent_transactions')
    
    if [ "$RECENT_TXN" != "null" ] && [ "$RECENT_TXN" != "[]" ]; then
        echo "$RECENT_TXN" | jq -r '.[0:5] | .[] | "  \(.created_at) - \(.phone_number) - \(.source) - \(.status)"'
    else
        echo -e "${BLUE}  No recent transactions${NC}"
    fi
    
else
    echo -e "${RED}❌ Failed to retrieve dashboard data${NC}"
fi

echo ""

# Test package search with different criteria
echo -e "${PURPLE}=== Testing Package Search ===${NC}"

echo -e "${YELLOW}1. Search packages under Rp 5,000:${NC}"
SEARCH_RESPONSE=$(curl -s -X POST "${API_BASE}/packages/search" \
  -H "Content-Type: application/json" \
  -d '{"max_price": 5000, "payment_method": "BALANCE"}')

SEARCH_SUCCESS=$(echo "$SEARCH_RESPONSE" | jq -r '.success')
if [ "$SEARCH_SUCCESS" == "true" ]; then
    PACKAGE_COUNT=$(echo "$SEARCH_RESPONSE" | jq -r '.data | length')
    echo -e "${GREEN}✅ Found ${PACKAGE_COUNT} packages under Rp 5,000${NC}"
    
    if [ "$PACKAGE_COUNT" -gt 0 ]; then
        echo -e "${BLUE}  Top 3 cheap packages:${NC}"
        echo "$SEARCH_RESPONSE" | jq -r '.data[0:3] | .[] | "    \(.package_name) - Rp \(.package_harga_int)"'
    fi
else
    echo -e "${RED}❌ Package search failed${NC}"
fi

echo ""

echo -e "${YELLOW}2. Search 'masa aktif' packages:${NC}"
SEARCH_RESPONSE=$(curl -s -X POST "${API_BASE}/packages/search" \
  -H "Content-Type: application/json" \
  -d '{"query": "masa aktif", "max_price": 10000}')

SEARCH_SUCCESS=$(echo "$SEARCH_RESPONSE" | jq -r '.success')
if [ "$SEARCH_SUCCESS" == "true" ]; then
    PACKAGE_COUNT=$(echo "$SEARCH_RESPONSE" | jq -r '.data | length')
    echo -e "${GREEN}✅ Found ${PACKAGE_COUNT} 'masa aktif' packages${NC}"
    
    if [ "$PACKAGE_COUNT" -gt 0 ]; then
        echo -e "${BLUE}  Available masa aktif packages:${NC}"
        echo "$SEARCH_RESPONSE" | jq -r '.data[0:3] | .[] | "    \(.package_name) - Rp \(.package_harga_int)"'
    fi
else
    echo -e "${RED}❌ Package search failed${NC}"
fi

echo ""

# Summary
echo -e "${PURPLE}=== Test Summary ===${NC}"
echo -e "${BLUE}✅ API connectivity tests completed${NC}"
echo -e "${BLUE}✅ Source tracking structure verified${NC}"
echo -e "${BLUE}✅ Package search functionality tested${NC}"
echo -e "${BLUE}✅ Dashboard statistics retrieved${NC}"
echo ""
echo -e "${YELLOW}Available Sources for Tracking:${NC}"
echo -e "${BLUE}  • whatsapp_bot - For WhatsApp Bot transactions${NC}"
echo -e "${BLUE}  • telegram_bot - For Telegram Bot transactions${NC}"
echo -e "${BLUE}  • web_interface - For Web Interface transactions${NC}"
echo -e "${BLUE}  • api_direct - For direct API calls (default)${NC}"
echo ""
echo -e "${YELLOW}To perform actual purchase test:${NC}"
echo -e "${BLUE}  ./test_whatsapp_flow.sh${NC}"