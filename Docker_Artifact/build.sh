#!/bin/bash

# ============================
# Build Script for Go Project
# ============================

# Exit if any command fails
set -e

# Create output directory if it doesn't exist
mkdir -p bin

# Output binary name
OUTPUT="bin/main"

echo "🔨 Building the Go application..."

# Build the Go application
go build -o $OUTPUT src/main.go

echo "✅ Build completed successfully."
echo "👉 Binary is located at: $OUTPUT"
