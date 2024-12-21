# Build Stage: Build the Go binary
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o hospital-doctor-service .

FROM ubuntu:22.04

# Install necessary dependencies including CA certificates for TLS verification
RUN apt-get update && apt-get install -y libc6 libc6-dev ca-certificates

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/hospital-doctor-service .

# Copy the .env file (if required)
COPY .env ./

# Set the command to run the binary
CMD ["./hospital-doctor-service"]