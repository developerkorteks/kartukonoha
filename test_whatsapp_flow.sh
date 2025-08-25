#!/bin/bash

# Test script untuk flow pembelian paket dari WhatsApp Bot
# Flow: Request OTP → Verify OTP → Purchase Package

set -e  # Exit on any error

# Configuration
API_BASE="http://localhost:8080/api"
PHONE_NUMBER="087786388052"
PACKAGE_CODE="XL_MASTIF_30D_P_V1"  # Paket masa aktif 30 hari
PAYMENT_METHOD="BALANCE"
SOURCE="whatsapp_bot"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Test Flow Pembelian Paket dari WhatsApp Bot ===${NC}"
echo -e "${BLUE}Phone: ${PHONE_NUMBER}${NC}"
echo -e "${BLUE}Package: ${PACKAGE_CODE}${NC}"
echo -e "${BLUE}Payment: ${PAYMENT_METHOD}${NC}"
echo -e "${BLUE}Source: ${SOURCE}${NC}"
echo ""

# Step 1: Request OTP
echo -e "${YELLOW}Step 1: Request OTP...${NC}"
OTP_RESPONSE=$(curl -s -X POST "${API_BASE}/otp/request" \
  -H "Content-Type: application/json" \
  -d "{
    \"phone_number\": \"${PHONE_NUMBER}\"
  }")

echo "OTP Response:"
echo "$OTP_RESPONSE" | jq '.'

# Check if OTP request was successful
OTP_SUCCESS=$(echo "$OTP_RESPONSE" | jq -r '.success')
if [ "$OTP_SUCCESS" != "true" ]; then
    echo -e "${RED}❌ OTP request failed!${NC}"
    exit 1
fi

echo -e "${GREEN}✅ OTP request successful!${NC}"
echo ""

# Step 2: Manual OTP input
echo -e "${YELLOW}Step 2: Enter OTP code${NC}"
echo -e "${BLUE}Please check your SMS and enter the OTP code:${NC}"
read -p "OTP Code: " OTP_CODE

if [ -z "$OTP_CODE" ]; then
    echo -e "${RED}❌ OTP code cannot be empty!${NC}"
    exit 1
fi

# Step 3: Verify OTP
echo -e "${YELLOW}Step 3: Verify OTP...${NC}"
VERIFY_RESPONSE=$(curl -s -X POST "${API_BASE}/otp/verify" \
  -H "Content-Type: application/json" \
  -d "{
    \"phone_number\": \"${PHONE_NUMBER}\",
    \"otp_code\": \"${OTP_CODE}\"
  }")

echo "Verify Response:"
echo "$VERIFY_RESPONSE" | jq '.'

# Check if OTP verification was successful
VERIFY_SUCCESS=$(echo "$VERIFY_RESPONSE" | jq -r '.success')
if [ "$VERIFY_SUCCESS" != "true" ]; then
    echo -e "${RED}❌ OTP verification failed!${NC}"
    exit 1
fi

# Extract access token
ACCESS_TOKEN=$(echo "$VERIFY_RESPONSE" | jq -r '.data.access_token')
if [ "$ACCESS_TOKEN" == "null" ] || [ -z "$ACCESS_TOKEN" ]; then
    echo -e "${RED}❌ Failed to get access token!${NC}"
    exit 1
fi

echo -e "${GREEN}✅ OTP verification successful!${NC}"
echo -e "${BLUE}Access Token: ${ACCESS_TOKEN}${NC}"
echo ""

# Step 4: Get available packages (optional - for verification)
echo -e "${YELLOW}Step 4: Check available packages...${NC}"
PACKAGES_RESPONSE=$(curl -s -X GET "${API_BASE}/packages?limit=5")
echo "Available packages (first 5):"
echo "$PACKAGES_RESPONSE" | jq '.data[0:5] | .[] | {package_code, package_name, package_price}'
echo ""

# Step 5: Purchase Package
echo -e "${YELLOW}Step 5: Purchase Package...${NC}"
PURCHASE_RESPONSE=$(curl -s -X POST "${API_BASE}/purchase" \
  -H "Content-Type: application/json" \
  -d "{
    \"phone_number\": \"${PHONE_NUMBER}\",
    \"package_code\": \"${PACKAGE_CODE}\",
    \"payment_method\": \"${PAYMENT_METHOD}\",
    \"access_token\": \"${ACCESS_TOKEN}\",
    \"source\": \"${SOURCE}\"
  }")

echo "Purchase Response:"
echo "$PURCHASE_RESPONSE" | jq '.'

# Check if purchase was successful
PURCHASE_SUCCESS=$(echo "$PURCHASE_RESPONSE" | jq -r '.success')
if [ "$PURCHASE_SUCCESS" == "true" ]; then
    echo -e "${GREEN}✅ Package purchase successful!${NC}"
    
    # Extract transaction details
    TRX_ID=$(echo "$PURCHASE_RESPONSE" | jq -r '.data.trx_id // "N/A"')
    PACKAGE_NAME=$(echo "$PURCHASE_RESPONSE" | jq -r '.data.package_name // "N/A"')
    PROCESSING_FEE=$(echo "$PURCHASE_RESPONSE" | jq -r '.data.package_processing_fee // "N/A"')
    
    echo -e "${BLUE}Transaction ID: ${TRX_ID}${NC}"
    echo -e "${BLUE}Package Name: ${PACKAGE_NAME}${NC}"
    echo -e "${BLUE}Processing Fee: ${PROCESSING_FEE}${NC}"
else
    echo -e "${RED}❌ Package purchase failed!${NC}"
    ERROR_MSG=$(echo "$PURCHASE_RESPONSE" | jq -r '.message')
    echo -e "${RED}Error: ${ERROR_MSG}${NC}"
fi

echo ""

# Step 6: Check transaction status (if successful)
if [ "$PURCHASE_SUCCESS" == "true" ] && [ "$TRX_ID" != "N/A" ]; then
    echo -e "${YELLOW}Step 6: Check transaction status...${NC}"
    sleep 2  # Wait a bit before checking
    
    STATUS_RESPONSE=$(curl -s -X POST "${API_BASE}/transaction/check" \
      -H "Content-Type: application/json" \
      -d "{
        \"transaction_id\": \"${TRX_ID}\"
      }")
    
    echo "Transaction Status:"
    echo "$STATUS_RESPONSE" | jq '.'
fi

echo ""

# Step 7: Check dashboard stats (to verify source tracking)
echo -e "${YELLOW}Step 7: Check dashboard stats (source tracking)...${NC}"
DASHBOARD_RESPONSE=$(curl -s -X GET "${API_BASE}/dashboard")

echo "Source Breakdown:"
echo "$DASHBOARD_RESPONSE" | jq '.data.source_breakdown'

echo ""
echo -e "${BLUE}=== Test Completed ===${NC}"

# Summary
echo -e "${YELLOW}=== Summary ===${NC}"
echo -e "Phone Number: ${PHONE_NUMBER}"
echo -e "Package Code: ${PACKAGE_CODE}"
echo -e "Payment Method: ${PAYMENT_METHOD}"
echo -e "Source: ${SOURCE}"
echo -e "OTP Success: $([ "$OTP_SUCCESS" == "true" ] && echo "✅" || echo "❌")"
echo -e "Verify Success: $([ "$VERIFY_SUCCESS" == "true" ] && echo "✅" || echo "❌")"
echo -e "Purchase Success: $([ "$PURCHASE_SUCCESS" == "true" ] && echo "✅" || echo "❌")"

if [ "$PURCHASE_SUCCESS" == "true" ]; then
    echo -e "${GREEN}🎉 WhatsApp Bot purchase flow completed successfully!${NC}"
else
    echo -e "${RED}❌ Purchase flow failed. Check the error messages above.${NC}"
fi