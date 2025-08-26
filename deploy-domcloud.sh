#!/bin/bash

# DomCloud Deployment Script for Nadia API
echo "🚀 Preparing Nadia API for DomCloud deployment..."

# Create public directory structure
echo "📁 Creating directory structure..."
mkdir -p public_html/public

# Build the application
echo "🔨 Building application..."
go mod download
go build -o app ./cmd/nadia

# Copy static files to public directory
echo "📄 Copying static files..."
cp *.html public_html/public/ 2>/dev/null || echo "No HTML files found"
cp favicon.ico public_html/public/ 2>/dev/null || echo "No favicon found"

# Create .htaccess for proper routing (if needed)
cat > public_html/public/.htaccess << 'EOF'
# Enable CORS for API requests
Header always set Access-Control-Allow-Origin "*"
Header always set Access-Control-Allow-Methods "GET, POST, PUT, DELETE, OPTIONS"
Header always set Access-Control-Allow-Headers "Content-Type, Authorization"

# Handle preflight requests
RewriteEngine On
RewriteCond %{REQUEST_METHOD} OPTIONS
RewriteRule ^(.*)$ $1 [R=200,L]
EOF

# Set executable permissions
chmod +x app

echo "✅ DomCloud deployment preparation complete!"
echo ""
echo "📋 Next steps:"
echo "1. Push your code to GitHub"
echo "2. Use this DomCloud configuration:"
echo ""
echo "source: https://github.com/yourusername/yourrepo.git"
echo "features:"
echo "  - go"
echo "nginx:"
echo "  root: public_html/public"
echo "  passenger:"
echo "    enabled: \"on\""
echo "    app_start_command: env PORT=\$PORT ./app"
echo "commands:"
echo "  - go mod download"
echo "  - go build -o app ./cmd/nadia"
echo "  - mkdir -p public_html/public"
echo "  - cp *.html public_html/public/ 2>/dev/null || true"
echo "  - cp favicon.ico public_html/public/ 2>/dev/null || true"
echo ""
echo "🌐 Your API will be available at: https://yourdomain.domcloud.co/api"
echo "📚 Swagger docs at: https://yourdomain.domcloud.co/swagger/index.html"