Single-Stage Docker Build

**Definition:** The Dockerfile does everything—build and run—in one single image layer.

**Result:** The final image includes build tools, source code, and the compiled binary.
```bash
FROM golang:1.22.2


WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./src ./src
WORKDIR /app/src
RUN go build -o /app/bookapi main.go
WORKDIR /app
EXPOSE 8088
CMD ["./bookapi"]

````

Multi-Stage Docker Build

**Definition:** The Dockerfile has two or more stages. The first stage compiles the binary. The second stage copies only the compiled binary into a minimal image (like debian:slim).

**Result:** The final image contains only what's needed to run the app—no source code, no compiler.

```bash

FROM golang:1.22.2 AS builder
WORKDIR /app
COPY . .
RUN chmod +x build.sh
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
