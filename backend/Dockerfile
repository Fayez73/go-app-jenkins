# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy module files and download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o backend main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/backend .

# Expose port
EXPOSE 8080

# Run the backend
CMD ["./backend"]
