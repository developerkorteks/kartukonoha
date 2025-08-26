# Nadia API Deployment Guide

## Environment Configuration

### API Base URL Configuration

For different environments, you can configure the API base URL using environment variables:

#### JavaScript Example
```bash
# For local development
export NADIA_API_BASE="http://localhost:8080/api"
# or
export NADIA_API_BASE="http://127.0.0.1:8080/api"

# For production
export NADIA_API_BASE="https://your-domain.com/api"

# Then run
node examples/javascript_example.js
```

#### Python Example
```bash
# For local development
export NADIA_API_BASE="http://localhost:8080/api"
# or
export NADIA_API_BASE="http://127.0.0.1:8080/api"

# For production
export NADIA_API_BASE="https://your-domain.com/api"

# Then run
python3 examples/python_example.py
```

### Swagger Configuration

The Swagger UI automatically adapts to the host you're accessing it from:

- **Local development**: Access via `http://localhost:8080/swagger/index.html` or `http://127.0.0.1:8080/swagger/index.html`
- **Production**: Access via `https://your-domain.com/swagger/index.html`

The API calls in Swagger UI will automatically use the same host and protocol as the page you're viewing.

### Server Configuration

You can override the Swagger host by setting the `SWAGGER_HOST` environment variable:

```bash
# For specific host configuration
export SWAGGER_HOST="api.yourdomain.com:8080"

# For production with custom domain
export SWAGGER_HOST="api.yourdomain.com"
```

## Deployment Examples

### DomCloud Deployment

DomCloud is a popular hosting platform that supports Go applications. Here's how to deploy Nadia API:

#### 1. Prepare Your Repository

Make sure your repository has the correct structure and push to GitHub:

```bash
git add .
git commit -m "Prepare for DomCloud deployment"
git push origin main
```

#### 2. DomCloud Configuration

Use this configuration in your DomCloud dashboard:

```yaml
source: https://github.com/yourusername/yourrepo.git
features:
  - go
nginx:
  root: public_html/public
  passenger:
    enabled: "on"
    app_start_command: env PORT=$PORT ./app
commands:
  - go mod download
  - go build -o app ./cmd/nadia
  - mkdir -p public_html/public
  - cp *.html public_html/public/ 2>/dev/null || true
  - cp favicon.ico public_html/public/ 2>/dev/null || true
```

#### 3. Environment Variables for DomCloud

Set these environment variables in your DomCloud dashboard:

```bash
ENVIRONMENT=production
SSO_API_BASE_URL=https://sso.putri-veronica.my.id/api
NADIA_API_URL=https://putri-veronica.my.id/api/user
SSO_STATIC_KEY=your-sso-static-key
NADIA_USERNAME=your-username
NADIA_PASSWORD=your-password
ADMIN_API_KEY=your-admin-api-key
JWT_SECRET=your-jwt-secret
TOKEN_EXPIRY_HOURS=8
DB_PATH=./nadia_transactions.db
```

#### 4. Access Your API

After deployment, your API will be available at:
- **API Base**: `https://yourdomain.domcloud.co/api`
- **Swagger Docs**: `https://yourdomain.domcloud.co/swagger/index.html`
- **Dashboard**: `https://yourdomain.domcloud.co/dashboard`

#### 5. Testing the Deployment

Test your API endpoints:

```bash
# Health check
curl https://yourdomain.domcloud.co/api/health

# Get packages
curl -X POST https://yourdomain.domcloud.co/api/packages \
  -H "Content-Type: application/json" \
  -d '{"query": "masa aktif", "max_price": 10000}'
```

### Docker Deployment
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o nadia ./cmd/nadia

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/nadia .
COPY --from=builder /app/*.html .

# Set environment variables
ENV ENVIRONMENT=production
ENV SERVER_ADDRESS=:8080
ENV SWAGGER_HOST=""

EXPOSE 8080
CMD ["./nadia"]
```

### Kubernetes Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nadia-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nadia-api
  template:
    metadata:
      labels:
        app: nadia-api
    spec:
      containers:
      - name: nadia-api
        image: nadia-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: ENVIRONMENT
          value: "production"
        - name: SERVER_ADDRESS
          value: ":8080"
        - name: SWAGGER_HOST
          value: ""  # Let Swagger UI auto-detect
---
apiVersion: v1
kind: Service
metadata:
  name: nadia-api-service
spec:
  selector:
    app: nadia-api
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

### Nginx Reverse Proxy
```nginx
server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

## Environment Variables Reference

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `ENVIRONMENT` | Application environment | `development` | `production` |
| `SERVER_ADDRESS` | Server bind address | `:8080` | `:3000` |
| `SWAGGER_HOST` | Override Swagger host | `""` (auto-detect) | `api.domain.com` |
| `NADIA_API_BASE` | API base URL for examples | `http://localhost:8080/api` | `https://api.domain.com/api` |

## Best Practices

1. **Never hardcode URLs** in production code
2. **Use environment variables** for configuration
3. **Let Swagger UI auto-detect** the host when possible
4. **Use HTTPS** in production
5. **Configure proper CORS** for web applications
6. **Use reverse proxy** for production deployments