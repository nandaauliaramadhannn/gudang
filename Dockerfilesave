# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o gudang .

# Stage 2: Create a minimal final image
FROM debian:bookworm-slim

# Copy the built binary from the build stage
COPY --from=builder /app/gudang /app/gudang

WORKDIR /app

# Set executable permissions
RUN chmod +x ./gudang

# Expose port 8080
EXPOSE 8080

# Set the entrypoint
CMD ["./gudang"]
