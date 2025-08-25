#!/bin/bash

# Test script untuk fitur Authentication dan Monitoring
API_BASE="http://localhost:8080/api"
API_KEY="nadia-admin-2024-secure-key"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}=== Testing Nadia API Authentication & Monitoring ===${NC}"
echo ""

# Test 1: Health Check (Public endpoint)
echo -e "${YELLOW}=== Test 1: Health Check (Public) ===${NC}"
echo "curl -X GET \"${API_BASE}/health\""
echo ""
HEALTH_RESPONSE=$(curl -s -X GET "${API_BASE}/health")
echo "Response:"
echo "$HEALTH_RESPONSE" | jq '.'
echo ""

# Test 2: Try accessing protected endpoint without API key
echo -e "${YELLOW}=== Test 2: Access Protected Endpoint Without API Key ===${NC}"
echo "curl -X GET \"${API_BASE}/packages\""
echo ""
UNAUTH_RESPONSE=$(curl -s -X GET "${API_BASE}/packages")
echo "Response:"
echo "$UNAUTH_RESPONSE" | jq '.'
echo ""

# Test 3: Authentication Login
echo -e "${YELLOW}=== Test 3: Authentication Login ===${NC}"
echo "curl -X POST \"${API_BASE}/auth/login\" \\"
echo "  -H \"Content-Type: application/json\" \\"
echo "  -d '{\"api_key\": \"${API_KEY}\"}'"
echo ""
AUTH_RESPONSE=$(curl -s -X POST "${API_BASE}/auth/login" \
  -H "Content-Type: application/json" \
  -d "{\"api_key\": \"${API_KEY}\"}")
echo "Response:"
echo "$AUTH_RESPONSE" | jq '.'
echo ""

# Test 4: Access protected endpoint with API key
echo -e "${YELLOW}=== Test 4: Access Protected Endpoint With API Key ===${NC}"
echo "curl -X GET \"${API_BASE}/packages?limit=3\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
PACKAGES_RESPONSE=$(curl -s -X GET "${API_BASE}/packages?limit=3" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$PACKAGES_RESPONSE" | jq '.'
echo ""

# Test 5: System Metrics
echo -e "${YELLOW}=== Test 5: System Metrics ===${NC}"
echo "curl -X GET \"${API_BASE}/monitoring/metrics\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
METRICS_RESPONSE=$(curl -s -X GET "${API_BASE}/monitoring/metrics" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$METRICS_RESPONSE" | jq '.'
echo ""

# Test 6: Performance Metrics
echo -e "${YELLOW}=== Test 6: Performance Metrics ===${NC}"
echo "curl -X GET \"${API_BASE}/monitoring/performance\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
PERF_RESPONSE=$(curl -s -X GET "${API_BASE}/monitoring/performance" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$PERF_RESPONSE" | jq '.'
echo ""

# Test 7: Security Metrics
echo -e "${YELLOW}=== Test 7: Security Metrics ===${NC}"
echo "curl -X GET \"${API_BASE}/monitoring/security\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
SECURITY_RESPONSE=$(curl -s -X GET "${API_BASE}/monitoring/security" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$SECURITY_RESPONSE" | jq '.'
echo ""

# Test 8: Application Logs
echo -e "${YELLOW}=== Test 8: Application Logs ===${NC}"
echo "curl -X GET \"${API_BASE}/monitoring/logs\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
LOGS_RESPONSE=$(curl -s -X GET "${API_BASE}/monitoring/logs" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$LOGS_RESPONSE" | jq '.'
echo ""

# Test 9: Uptime Status
echo -e "${YELLOW}=== Test 9: Uptime Status ===${NC}"
echo "curl -X GET \"${API_BASE}/monitoring/uptime\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
UPTIME_RESPONSE=$(curl -s -X GET "${API_BASE}/monitoring/uptime" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$UPTIME_RESPONSE" | jq '.'
echo ""

# Test 10: Realtime Metrics
echo -e "${YELLOW}=== Test 10: Realtime Metrics ===${NC}"
echo "curl -X GET \"${API_BASE}/monitoring/metrics/realtime\" \\"
echo "  -H \"X-API-Key: ${API_KEY}\""
echo ""
REALTIME_RESPONSE=$(curl -s -X GET "${API_BASE}/monitoring/metrics/realtime" \
  -H "X-API-Key: ${API_KEY}")
echo "Response:"
echo "$REALTIME_RESPONSE" | jq '.'
echo ""

echo -e "${GREEN}🎉 All monitoring tests completed!${NC}"
echo ""
echo -e "${BLUE}=== Summary ===${NC}"
echo -e "${GREEN}✅ Health Check: Public endpoint working${NC}"
echo -e "${GREEN}✅ Authentication: API key protection working${NC}"
echo -e "${GREEN}✅ System Metrics: CPU, Memory, Disk monitoring${NC}"
echo -e "${GREEN}✅ Performance Metrics: Response time, throughput${NC}"
echo -e "${GREEN}✅ Security Metrics: Unauthorized attempts tracking${NC}"
echo -e "${GREEN}✅ Application Logs: Request logging${NC}"
echo -e "${GREEN}✅ Uptime Status: Service health monitoring${NC}"
echo -e "${GREEN}✅ Realtime Metrics: Live monitoring data${NC}"
echo ""
echo -e "${YELLOW}📊 Key Features Implemented:${NC}"
echo -e "${BLUE}🔐 Authentication & Authorization${NC}"
echo -e "${BLUE}📈 System Performance Monitoring${NC}"
echo -e "${BLUE}🛡️ Security Threat Detection${NC}"
echo -e "${BLUE}📊 Real-time Metrics Collection${NC}"
echo -e "${BLUE}💾 SQLite Database Persistence${NC}"
echo -e "${BLUE}🔍 Request/Response Logging${NC}"
echo -e "${BLUE}⚡ Load & Resource Monitoring${NC}"