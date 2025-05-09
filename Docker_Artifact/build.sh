#!/bin/bash

# ============================
# Build Script for Go Project
# ============================

set -e  # Exit on error

# Navigate to the script's root (assuming it's in Docker_Artifact)
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

# Ensure output directory exists
mkdir -p bin

# Navigate to source directory
cd src

# Initialize go module if go.mod doesn't exist
if [ ! -f "go.mod" ]; then
    echo "🔧 Initializing Go module..."
    go mod init example.com/myapp
fi

# Ensure dependencies are downloaded
echo "📦 Fetching dependencies..."
go get github.com/gin-gonic/gin
go mod tidy

# Build the Go application
echo "🔨 Building the Go application..."
go build -o ../bin/main main.go

echo "✅ Build completed successfully."
echo "👉 Binary is located at: bin/main"

