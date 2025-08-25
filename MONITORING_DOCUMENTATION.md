# 📊 Nadia API - Monitoring & Authentication Documentation

## 🔐 Authentication System

### API Key Authentication
- **Admin API Key**: `nadia-admin-2024-secure-key`
- **Header**: `X-API-Key: nadia-admin-2024-secure-key`
- **Query Parameter**: `?api_key=nadia-admin-2024-secure-key`

### Protected Endpoints
All API endpoints (except `/api/health` and `/swagger/*`) require authentication.

### Authentication Login
```bash
curl -X POST "http://localhost:8080/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"api_key": "nadia-admin-2024-secure-key"}'
```

## 📈 Monitoring Endpoints

### 1. System Metrics
**Endpoint**: `GET /api/monitoring/metrics`
**Description**: Comprehensive system monitoring data

**Response includes**:
- Service Status (running/stopped)
- CPU Usage (%)
- Memory Usage (bytes, %, leak detection)
- Disk Usage (bytes, %, I/O stats)
- Network Traffic (bytes sent/received, packets)
- Load Average (1min, 5min, 15min)
- Response Time Metrics (avg, P95, P99)
- Error Rate (%)
- Throughput (requests/sec)
- Database Performance
- Security Metrics
- Uptime Check

### 2. Performance Metrics
**Endpoint**: `GET /api/monitoring/performance`
**Description**: Detailed performance analysis

**Metrics**:
- Response Time (average, P95, P99, min, max)
- Throughput (requests per second)
- Error Rate (percentage)
- CPU Usage (real-time)
- Memory Statistics (allocated, total, system, GC runs)
- Uptime (seconds)

### 3. Security Metrics
**Endpoint**: `GET /api/monitoring/security`
**Description**: Security monitoring and threat detection

**Features**:
- Unauthorized access attempts
- Suspicious traffic detection
- Recent security threats log
- IP-based threat tracking
- Severity classification (low, medium, high)

### 4. Application Logs
**Endpoint**: `GET /api/monitoring/logs`
**Description**: Recent application request logs

**Log Data**:
- Timestamp
- HTTP Method
- Request Path
- Status Code
- Response Time
- Client IP
- User Agent

### 5. Uptime Status
**Endpoint**: `GET /api/monitoring/uptime`
**Description**: Service health and uptime monitoring

**Status Info**:
- Service Up/Down status
- Health check results
- Uptime duration (human readable)
- Start time
- Current time
- Version information
- Environment details

### 6. Realtime Metrics
**Endpoint**: `GET /api/monitoring/metrics/realtime`
**Description**: Live monitoring data (last 10 minutes)

**Real-time Data**:
- Recent requests (last 10 minutes)
- Live performance metrics
- Active connections
- Current resource usage

## 🛡️ Security Features

### Threat Detection
- **Unauthorized Access**: Tracks invalid API key attempts
- **Slow Response**: Detects responses > 5 seconds
- **Suspicious Traffic**: Monitors unusual request patterns
- **IP Tracking**: Logs source IPs for security events

### Security Logging
All security events are logged with:
- IP Address
- Threat Type
- Severity Level
- Timestamp
- Detailed Description

### Automatic Protection
- Rate limiting through monitoring
- Automatic threat classification
- Real-time security alerts
- Request pattern analysis

## 💾 Database Persistence

### SQLite Database
- **File**: `./nadia_transactions.db`
- **Tables**: `transactions`
- **Features**: 
  - Automatic schema creation
  - Transaction persistence
  - Data recovery on restart
  - Indexed queries for performance

### Transaction Tracking
- All transactions saved to database
- In-memory cache for fast access
- Automatic sync between memory and database
- Transaction history preservation

## 📊 Monitoring Metrics Details

### CPU Monitoring
- Real-time CPU usage percentage
- Alert when CPU > 80% for extended periods
- Multi-core CPU support
- Historical CPU trends

### Memory Monitoring
- Used/Total memory tracking
- Memory leak detection (>90% usage)
- Garbage collection statistics
- Memory allocation patterns

### Disk Monitoring
- Disk usage percentage
- Free space monitoring
- I/O operations tracking
- Storage bottleneck detection

### Network Monitoring
- Bytes sent/received
- Packet statistics
- Network latency measurement
- Bandwidth utilization

### Response Time Analysis
- Average response time
- P95 and P99 percentiles
- Min/Max response times
- Response time trends

## 🔧 Configuration

### Environment Variables
```bash
# API Configuration
API_PORT=8080
API_KEY=nadia-admin-2024-secure-key

# Database
DB_PATH=./nadia_transactions.db

# Monitoring
METRICS_RETENTION=1000  # Keep last 1000 requests
SECURITY_RETENTION=100  # Keep last 100 threats
```

### Monitoring Intervals
- **CPU/Memory**: Real-time (1 second intervals)
- **Disk Usage**: On-demand
- **Network Stats**: Real-time
- **Request Metrics**: Per request
- **Security Checks**: Per request

## 📋 Usage Examples

### Basic Monitoring Check
```bash
# Check system health
curl -H "X-API-Key: nadia-admin-2024-secure-key" \
  http://localhost:8080/api/monitoring/uptime

# Get performance metrics
curl -H "X-API-Key: nadia-admin-2024-secure-key" \
  http://localhost:8080/api/monitoring/performance
```

### Security Monitoring
```bash
# Check security threats
curl -H "X-API-Key: nadia-admin-2024-secure-key" \
  http://localhost:8080/api/monitoring/security

# View application logs
curl -H "X-API-Key: nadia-admin-2024-secure-key" \
  http://localhost:8080/api/monitoring/logs
```

### Real-time Monitoring
```bash
# Get live metrics
curl -H "X-API-Key: nadia-admin-2024-secure-key" \
  http://localhost:8080/api/monitoring/metrics/realtime

# Full system metrics
curl -H "X-API-Key: nadia-admin-2024-secure-key" \
  http://localhost:8080/api/monitoring/metrics
```

## 🚨 Alerts & Thresholds

### Performance Alerts
- **CPU Usage > 80%**: High CPU usage warning
- **Memory Usage > 90%**: Memory leak detection
- **Response Time > 5s**: Slow response alert
- **Error Rate > 5%**: High error rate warning

### Security Alerts
- **Unauthorized Access**: Invalid API key attempts
- **Suspicious Traffic**: Unusual request patterns
- **High Response Time**: Potential DoS indicators
- **Multiple Failed Attempts**: Brute force detection

## 📈 Dashboard Integration

The monitoring data can be integrated with:
- **Grafana**: For visual dashboards
- **Prometheus**: For metrics collection
- **ELK Stack**: For log analysis
- **Custom Dashboards**: Using the API endpoints

## 🔄 Maintenance

### Database Maintenance
- Automatic cleanup of old records
- Database optimization
- Backup recommendations
- Recovery procedures

### Log Rotation
- Automatic log cleanup
- Configurable retention periods
- Log compression
- Archive management

---

## 🎯 Key Benefits

✅ **Real-time Monitoring**: Live system metrics and performance data
✅ **Security Protection**: Comprehensive threat detection and logging
✅ **Data Persistence**: SQLite database for transaction history
✅ **API Authentication**: Secure access control with API keys
✅ **Performance Analysis**: Detailed response time and throughput metrics
✅ **Resource Monitoring**: CPU, memory, disk, and network tracking
✅ **Automated Alerts**: Built-in threshold monitoring and alerting
✅ **Easy Integration**: RESTful API endpoints for external tools

This monitoring system provides enterprise-level observability for the Nadia API, ensuring optimal performance, security, and reliability.