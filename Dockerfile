# Use a lightweight Go base image
FROM golang:1.23.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download all dependencies. Caching layers are leveraged.
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main ./

# Expose any ports your app uses (if any)
# EXPOSE 8080

# Run the application
CMD ["./main"]
