#!/bin/bash

set -e

echo " Creating bin directory if it doesn't exist..."
mkdir -p bin

echo " Building the Go application from src/main.go..."
go build -o bin/bookapi ./src/main.go

echo " Build completed! Executable is at bin/bookapi"
