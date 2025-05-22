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
