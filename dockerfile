# Use Go 1.22.1 or higher (Choose the version you prefer)
FROM golang:1.22.1-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download the Go modules dependencies
RUN go mod download

# Copy the entire application
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port your application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
