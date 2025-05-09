#!/bin/bash
set -e

# Set Go module cache to local folder
export GOMODCACHE=$(pwd)/gocache

# Navigate to the script's root
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR"

mkdir -p bin

cd src

if [ ! -f "go.mod" ]; then
    echo "🔧 Initializing Go module..."
    go mod init example.com/myapp
fi

echo "📦 Fetching dependencies..."
go get github.com/gin-gonic/gin
go mod tidy

echo "🔨 Building the Go application..."
go build -o ../bin/main main.go

echo "✅ Build completed successfully."
echo "👉 Binary is located at: bin/main"
