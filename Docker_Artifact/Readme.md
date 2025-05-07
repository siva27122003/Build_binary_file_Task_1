# Go Project Docker Setup

This project demonstrates how to build and run a Go application using Docker.

## Project Structure
go_project/
├── Dockerfile
├── build.sh
├── main.go
├── go.mod
└── go.sum

## build.sh Script

Ensure your `build.sh` script contains the following:

```bash
#!/bin/bash
set -e
echo "Building the Go application..."
mkdir -p bin
go build -o bin/bookapi main.go
```

## Dockerfile

```dockerfile
# Stage 1: Build
FROM golang:1.22.2 AS builder
WORKDIR /app
COPY . .
RUN chmod +x build.sh
RUN ./build.sh

# Stage 2: Final slim image
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/bin/bookapi .
EXPOSE 8088
CMD ["./bookapi"]
```

## Build and Run

### Step 1: Build the Docker image

```bash
docker build -t my_app_image .
```

### Step 2: Run the Docker container

```bash
docker run -it -p 8088:8088 my_app_image
```

## Common Errors and Fixes

### Error 1: Go version mismatch

```
go: go.mod requires go >= 1.22.2 (running go 1.21.13; GOTOOLCHAIN=local)
```

**Cause:** The Docker image used an outdated Go version.

**Fix:** Update Dockerfile to use the correct Go version:

```dockerfile
FROM golang:1.22.2 AS builder
```

---

### Error 2: Executable not found in final image

```
exec: "./bin/bookapi": stat ./bin/bookapi: no such file or directory
```

**Cause:** The binary was either not built properly or copied incorrectly in the final stage.

**Fix:**

**Option A (Recommended):** Place the binary in the root of the image:

```dockerfile
COPY --from=builder /app/bin/bookapi .
CMD ["./bookapi"]
```

**Option B:** Retain directory structure and adjust the command:

```dockerfile
RUN mkdir -p bin
COPY --from=builder /app/bin/bookapi ./bin/bookapi
CMD ["./bin/bookapi"]
```

