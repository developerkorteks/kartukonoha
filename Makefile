# Makefile for Nadia API

# Default build (for local development)
.PHONY: build
build:
	go build -o nadia ./cmd/nadia

# Build for DomCloud deployment
.PHONY: build-domcloud
build-domcloud:
	go build -o app ./cmd/nadia

# Build and prepare static files for DomCloud
.PHONY: deploy-domcloud
deploy-domcloud: build-domcloud
	mkdir -p public_html/public
	cp *.html public_html/public/ 2>/dev/null || true
	cp favicon.ico public_html/public/ 2>/dev/null || true

# Clean build artifacts
.PHONY: clean
clean:
	rm -f nadia app
	rm -rf public_html

# Run locally
.PHONY: run
run: build
	./nadia

# Run with DomCloud-like environment
.PHONY: run-domcloud
run-domcloud: build-domcloud
	PORT=8080 ./app

# Test build
.PHONY: test
test:
	go test ./...

# Install dependencies
.PHONY: deps
deps:
	go mod download
	go mod tidy

# Generate swagger docs
.PHONY: swagger
swagger:
	swag init -g cmd/nadia/main.go -o docs

.DEFAULT_GOAL := build