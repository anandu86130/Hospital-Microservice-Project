# Build Stage: Build the Go binary
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Copy Go modules and source files
COPY go.mod go.sum ./ 
RUN go mod download
COPY . .

# Build the binary
RUN go build -o main .

# Final Stage: Create a smaller image
FROM alpine:latest
WORKDIR /app

# Copy the binary and templates
COPY --from=builder /app/main .
COPY --from=builder /app/template ./template

# Copy the .env file
COPY .env ./

# Expose port and start the app
EXPOSE 8080
CMD ["./main"]
