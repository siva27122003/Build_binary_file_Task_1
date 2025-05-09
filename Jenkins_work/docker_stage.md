Single-Stage Docker Build

**Definition:** The Dockerfile does everything—build and run—in one single image layer.

**Result:** The final image includes build tools, source code, and the compiled binary.
```bash
FROM golang:1.22.2

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY ./src ./src

# Build the Go app
WORKDIR /app/src
RUN go build -o /app/bookapi main.go

# Set working directory to binary location
WORKDIR /app
EXPOSE 8088

# Run the binary
CMD ["./bookapi"]

````

Multi-Stage Docker Build

**Definition:** The Dockerfile has two or more stages. The first stage compiles the binary. The second stage copies only the compiled binary into a minimal image (like debian:slim).

**Result:** The final image contains only what's needed to run the app—no source code, no compiler.

```bash
# Use the official Golang image as the builder
FROM golang:1.22.2 AS builder


# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Make the build script executable
RUN chmod +x build.sh

# Run the build script
RUN ./build.sh

# ============================
# Create a smaller final image
# ============================
FROM debian:bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/bookapi .

# Expose the port your app uses (optional)
EXPOSE 8088

# Run the binary
CMD ["./bookapi"]

```
