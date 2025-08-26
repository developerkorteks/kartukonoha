#!/bin/bash

# Start script for Nadia API
# This script can be used for deployment on various hosting platforms

# Set default environment if not provided
export ENVIRONMENT=${ENVIRONMENT:-production}

# Set default port if not provided
export PORT=${PORT:-8080}

# Create data directory if it doesn't exist
mkdir -p ./data

# Set database path
export DB_PATH=${DB_PATH:-./data/nadia_transactions.db}

# Log startup information
echo "Starting Nadia API..."
echo "Environment: $ENVIRONMENT"
echo "Port: $PORT"
echo "Database: $DB_PATH"

# Start the application
exec ./nadia