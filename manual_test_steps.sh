#!/bin/bash

# Manual step-by-step testing untuk WhatsApp Bot flow
# Setiap step bisa dijalankan satu per satu

API_BASE="http://localhost:8080/api"
PHONE_NUMBER="087817739901"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

show_help() {
    echo -e "${BLUE}=== Manual Testing Steps for WhatsApp Bot Flow ===${NC}"
    echo ""
    echo "Usage: $0 [step]"
    echo ""
    echo "Available steps:"
    echo "  1 | otp       - Request OTP"
    echo "  2 | verify    - Verify OTP (requires OTP code input)"
    echo "  3 | packages  - List available packages"
    echo "  4 | search    - Search packages"
    echo "  5 | balance   - Check balance"
    echo "  6 | purchase  - Purchase package (requires access token)"
    echo "  7 | dashboard - Check dashboard stats"
    echo "  8 | health    - API health check"
    echo "  all           - Run all steps (interactive)"
    echo ""
    echo "Examples:"
    echo "  $0 1          # Request OTP"
    echo "  $0 otp        # Request OTP"
    echo "  $0 verify     # Verify OTP"
    echo "  $0 all        # Run all steps"
}

step_otp() {
    echo -e "${YELLOW}=== Step 1: Request OTP ===${NC}"
    echo -e "${BLUE}Phone: ${PHONE_NUMBER}${NC}"
    echo ""
    
    echo "curl -X POST \"${API_BASE}/otp/request\" \\"
    echo "  -H \"Content-Type: application/json\" \\"
    echo "  -d '{\"phone_number\": \"${PHONE_NUMBER}\"}'"
    echo ""
    
    OTP_RESPONSE=$(curl -s -X POST "${API_BASE}/otp/request" \
      -H "Content-Type: application/json" \
      -d "{\"phone_number\": \"${PHONE_NUMBER}\"}")
    
    echo "Response:"
    echo "$OTP_RESPONSE" | jq '.'
    
    OTP_SUCCESS=$(echo "$OTP_RESPONSE" | jq -r '.success')
    if [ "$OTP_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ OTP request successful!${NC}"
        CAN_RESEND=$(echo "$OTP_RESPONSE" | jq -r '.data.can_resend_in')
        echo -e "${BLUE}Can resend in: ${CAN_RESEND} seconds${NC}"
    else
        echo -e "${RED}❌ OTP request failed!${NC}"
    fi
}

step_verify() {
    echo -e "${YELLOW}=== Step 2: Verify OTP ===${NC}"
    echo -e "${BLUE}Phone: ${PHONE_NUMBER}${NC}"
    echo ""
    
    read -p "Enter OTP code: " OTP_CODE
    
    if [ -z "$OTP_CODE" ]; then
        echo -e "${RED}❌ OTP code cannot be empty!${NC}"
        return 1
    fi
    
    echo ""
    echo "curl -X POST \"${API_BASE}/otp/verify\" \\"
    echo "  -H \"Content-Type: application/json\" \\"
    echo "  -d '{"
    echo "    \"phone_number\": \"${PHONE_NUMBER}\","
    echo "    \"otp_code\": \"${OTP_CODE}\""
    echo "  }'"
    echo ""
    
    VERIFY_RESPONSE=$(curl -s -X POST "${API_BASE}/otp/verify" \
      -H "Content-Type: application/json" \
      -d "{
        \"phone_number\": \"${PHONE_NUMBER}\",
        \"otp_code\": \"${OTP_CODE}\"
      }")
    
    echo "Response:"
    echo "$VERIFY_RESPONSE" | jq '.'
    
    VERIFY_SUCCESS=$(echo "$VERIFY_RESPONSE" | jq -r '.success')
    if [ "$VERIFY_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ OTP verification successful!${NC}"
        ACCESS_TOKEN=$(echo "$VERIFY_RESPONSE" | jq -r '.data.access_token')
        echo -e "${BLUE}Access Token: ${ACCESS_TOKEN}${NC}"
        echo ""
        echo -e "${YELLOW}💾 Save this access token for purchase step!${NC}"
        echo "$ACCESS_TOKEN" > /tmp/nadia_access_token.txt
        echo -e "${BLUE}Token saved to: /tmp/nadia_access_token.txt${NC}"
    else
        echo -e "${RED}❌ OTP verification failed!${NC}"
    fi
}

step_packages() {
    echo -e "${YELLOW}=== Step 3: List Available Packages ===${NC}"
    echo ""
    
    echo "curl -X GET \"${API_BASE}/packages?limit=5\""
    echo ""
    
    PACKAGES_RESPONSE=$(curl -s -X GET "${API_BASE}/packages?limit=5")
    
    echo "Response (first 5 packages):"
    echo "$PACKAGES_RESPONSE" | jq '.'
    
    PACKAGES_SUCCESS=$(echo "$PACKAGES_RESPONSE" | jq -r '.success')
    if [ "$PACKAGES_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ Package list retrieved!${NC}"
        PACKAGE_COUNT=$(echo "$PACKAGES_RESPONSE" | jq -r '.data | length')
        echo -e "${BLUE}Found ${PACKAGE_COUNT} packages${NC}"
        
        echo ""
        echo -e "${BLUE}Package details:${NC}"
        echo "$PACKAGES_RESPONSE" | jq -r '.data[] | "  • \(.package_name) - Rp \(.package_harga_int) (\(.package_code))"'
    else
        echo -e "${RED}❌ Failed to get packages!${NC}"
    fi
}

step_search() {
    echo -e "${YELLOW}=== Step 4: Search Packages ===${NC}"
    echo ""
    
    echo "Searching for packages under Rp 5,000 with BALANCE payment:"
    echo ""
    echo "curl -X POST \"${API_BASE}/packages/search\" \\"
    echo "  -H \"Content-Type: application/json\" \\"
    echo "  -d '{"
    echo "    \"max_price\": 5000,"
    echo "    \"payment_method\": \"BALANCE\""
    echo "  }'"
    echo ""
    
    SEARCH_RESPONSE=$(curl -s -X POST "${API_BASE}/packages/search" \
      -H "Content-Type: application/json" \
      -d '{
        "max_price": 5000,
        "payment_method": "BALANCE"
      }')
    
    echo "Response:"
    echo "$SEARCH_RESPONSE" | jq '.'
    
    SEARCH_SUCCESS=$(echo "$SEARCH_RESPONSE" | jq -r '.success')
    if [ "$SEARCH_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ Package search successful!${NC}"
        PACKAGE_COUNT=$(echo "$SEARCH_RESPONSE" | jq -r '.data | length')
        echo -e "${BLUE}Found ${PACKAGE_COUNT} packages under Rp 5,000${NC}"
        
        if [ "$PACKAGE_COUNT" -gt 0 ]; then
            echo ""
            echo -e "${BLUE}Cheap packages:${NC}"
            echo "$SEARCH_RESPONSE" | jq -r '.data[0:3] | .[] | "  • \(.package_name) - Rp \(.package_harga_int)"'
        fi
    else
        echo -e "${RED}❌ Package search failed!${NC}"
    fi
}

step_balance() {
    echo -e "${YELLOW}=== Step 5: Check Balance ===${NC}"
    echo ""
    
    echo "curl -X GET \"${API_BASE}/balance\""
    echo ""
    
    BALANCE_RESPONSE=$(curl -s -X GET "${API_BASE}/balance")
    
    echo "Response:"
    echo "$BALANCE_RESPONSE" | jq '.'
    
    BALANCE_SUCCESS=$(echo "$BALANCE_RESPONSE" | jq -r '.success')
    if [ "$BALANCE_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ Balance check successful!${NC}"
        BALANCE=$(echo "$BALANCE_RESPONSE" | jq -r '.data.balance // "N/A"')
        echo -e "${BLUE}Current Balance: Rp ${BALANCE}${NC}"
    else
        echo -e "${RED}❌ Balance check failed!${NC}"
    fi
}

step_purchase() {
    echo -e "${YELLOW}=== Step 6: Purchase Package ===${NC}"
    echo ""
    
    # Try to load saved access token
    if [ -f "/tmp/nadia_access_token.txt" ]; then
        SAVED_TOKEN=$(cat /tmp/nadia_access_token.txt)
        echo -e "${BLUE}Found saved access token: ${SAVED_TOKEN}${NC}"
        read -p "Use saved token? (y/n): " USE_SAVED
        if [ "$USE_SAVED" == "y" ] || [ "$USE_SAVED" == "Y" ]; then
            ACCESS_TOKEN="$SAVED_TOKEN"
        fi
    fi
    
    if [ -z "$ACCESS_TOKEN" ]; then
        read -p "Enter access token: " ACCESS_TOKEN
    fi
    
    if [ -z "$ACCESS_TOKEN" ]; then
        echo -e "${RED}❌ Access token cannot be empty!${NC}"
        echo -e "${YELLOW}Please run step 2 (verify) first to get access token${NC}"
        return 1
    fi
    
    PACKAGE_CODE="XL_MASTIF_30D_P_V1"
    
    echo ""
    echo -e "${BLUE}Purchase details:${NC}"
    echo -e "${BLUE}  Phone: ${PHONE_NUMBER}${NC}"
    echo -e "${BLUE}  Package: ${PACKAGE_CODE}${NC}"
    echo -e "${BLUE}  Payment: BALANCE${NC}"
    echo -e "${BLUE}  Source: whatsapp_bot${NC}"
    echo ""
    
    echo "curl -X POST \"${API_BASE}/purchase\" \\"
    echo "  -H \"Content-Type: application/json\" \\"
    echo "  -d '{"
    echo "    \"phone_number\": \"${PHONE_NUMBER}\","
    echo "    \"package_code\": \"${PACKAGE_CODE}\","
    echo "    \"payment_method\": \"BALANCE\","
    echo "    \"access_token\": \"${ACCESS_TOKEN}\","
    echo "    \"source\": \"whatsapp_bot\""
    echo "  }'"
    echo ""
    
    read -p "Proceed with purchase? (y/n): " CONFIRM
    if [ "$CONFIRM" != "y" ] && [ "$CONFIRM" != "Y" ]; then
        echo -e "${YELLOW}Purchase cancelled${NC}"
        return 0
    fi
    
    PURCHASE_RESPONSE=$(curl -s -X POST "${API_BASE}/purchase" \
      -H "Content-Type: application/json" \
      -d "{
        \"phone_number\": \"${PHONE_NUMBER}\",
        \"package_code\": \"${PACKAGE_CODE}\",
        \"payment_method\": \"BALANCE\",
        \"access_token\": \"${ACCESS_TOKEN}\",
        \"source\": \"whatsapp_bot\"
      }")
    
    echo "Response:"
    echo "$PURCHASE_RESPONSE" | jq '.'
    
    PURCHASE_SUCCESS=$(echo "$PURCHASE_RESPONSE" | jq -r '.success')
    if [ "$PURCHASE_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ Package purchase successful!${NC}"
        TRX_ID=$(echo "$PURCHASE_RESPONSE" | jq -r '.data.trx_id // "N/A"')
        PACKAGE_NAME=$(echo "$PURCHASE_RESPONSE" | jq -r '.data.package_name // "N/A"')
        echo -e "${BLUE}Transaction ID: ${TRX_ID}${NC}"
        echo -e "${BLUE}Package Name: ${PACKAGE_NAME}${NC}"
    else
        echo -e "${RED}❌ Package purchase failed!${NC}"
        ERROR_MSG=$(echo "$PURCHASE_RESPONSE" | jq -r '.message')
        echo -e "${RED}Error: ${ERROR_MSG}${NC}"
    fi
}

step_dashboard() {
    echo -e "${YELLOW}=== Step 7: Check Dashboard Stats ===${NC}"
    echo ""
    
    echo "curl -X GET \"${API_BASE}/dashboard\""
    echo ""
    
    DASHBOARD_RESPONSE=$(curl -s -X GET "${API_BASE}/dashboard")
    
    DASHBOARD_SUCCESS=$(echo "$DASHBOARD_RESPONSE" | jq -r '.success')
    if [ "$DASHBOARD_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ Dashboard data retrieved!${NC}"
        echo ""
        
        # Overall stats
        TOTAL_TXN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.total_transactions // 0')
        SUCCESS_RATE=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.success_rate // 0')
        TOTAL_REVENUE=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.stats.total_revenue // 0')
        
        echo -e "${BLUE}Overall Statistics:${NC}"
        echo -e "${BLUE}  Total Transactions: ${TOTAL_TXN}${NC}"
        echo -e "${BLUE}  Success Rate: ${SUCCESS_RATE}%${NC}"
        echo -e "${BLUE}  Total Revenue: Rp ${TOTAL_REVENUE}${NC}"
        echo ""
        
        # Source breakdown
        echo -e "${BLUE}Source Breakdown:${NC}"
        SOURCE_BREAKDOWN=$(echo "$DASHBOARD_RESPONSE" | jq -r '.data.source_breakdown')
        if [ "$SOURCE_BREAKDOWN" != "null" ] && [ "$SOURCE_BREAKDOWN" != "[]" ]; then
            echo "$SOURCE_BREAKDOWN" | jq -r '.[] | "  \(.source): \(.count) transactions, Rp \(.revenue) revenue"'
        else
            echo -e "${BLUE}  No transactions recorded yet${NC}"
        fi
        
    else
        echo -e "${RED}❌ Failed to get dashboard data!${NC}"
        echo "Response:"
        echo "$DASHBOARD_RESPONSE" | jq '.'
    fi
}

step_health() {
    echo -e "${YELLOW}=== Step 8: API Health Check ===${NC}"
    echo ""
    
    echo "curl -X GET \"${API_BASE}/health\""
    echo ""
    
    HEALTH_RESPONSE=$(curl -s -X GET "${API_BASE}/health")
    
    echo "Response:"
    echo "$HEALTH_RESPONSE" | jq '.'
    
    HEALTH_SUCCESS=$(echo "$HEALTH_RESPONSE" | jq -r '.success')
    if [ "$HEALTH_SUCCESS" == "true" ]; then
        echo -e "${GREEN}✅ API is healthy!${NC}"
        STATUS=$(echo "$HEALTH_RESPONSE" | jq -r '.data.status // "N/A"')
        echo -e "${BLUE}Status: ${STATUS}${NC}"
    else
        echo -e "${RED}❌ API health check failed!${NC}"
    fi
}

run_all() {
    echo -e "${BLUE}=== Running All Steps Interactively ===${NC}"
    echo ""
    
    step_health
    echo ""; read -p "Press Enter to continue..."
    
    step_packages
    echo ""; read -p "Press Enter to continue..."
    
    step_search
    echo ""; read -p "Press Enter to continue..."
    
    step_balance
    echo ""; read -p "Press Enter to continue..."
    
    step_otp
    echo ""; read -p "Press Enter to continue..."
    
    step_verify
    echo ""; read -p "Press Enter to continue..."
    
    step_purchase
    echo ""; read -p "Press Enter to continue..."
    
    step_dashboard
    
    echo ""
    echo -e "${GREEN}🎉 All steps completed!${NC}"
}

# Main script logic
case "${1:-help}" in
    1|otp)
        step_otp
        ;;
    2|verify)
        step_verify
        ;;
    3|packages)
        step_packages
        ;;
    4|search)
        step_search
        ;;
    5|balance)
        step_balance
        ;;
    6|purchase)
        step_purchase
        ;;
    7|dashboard)
        step_dashboard
        ;;
    8|health)
        step_health
        ;;
    all)
        run_all
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        echo -e "${RED}Unknown step: $1${NC}"
        echo ""
        show_help
        exit 1
        ;;
esac