# Build Stage: Build the Go binary
FROM golang:1.23.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o hospital-booking-service .

# Final Stage: Create a lightweight production image
FROM ubuntu:22.04

# Install runtime dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    libc6 libc6-dev ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /app/hospital-booking-service ./

# Copy the .env file, if it exists
COPY .env ./

# Set the command to run the binary
CMD ["./hospital-booking-service"]
