#!/bin/bash
set -e

# Redirect Go cache and sumdb to local writable directory
export GOPATH=$(pwd)/go
export GOMODCACHE=$GOPATH/pkg/mod
export GOCACHE=$GOPATH/pkg/cache

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
