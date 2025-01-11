# Use the official Golang image as a base image
FROM golang:1.23.4 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire project directory
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Use a minimal base image for production
FROM alpine:latest

# Install CA certificates to make HTTPS calls possible
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy other necessary files, e.g., dummy_data.sql, if needed
COPY --from=builder /app/data/dummy_data.sql ./data/dummy_data.sql

# Expose the application port (replace 8080 with your actual port)
EXPOSE 8080

# Command to run the application
CMD ["./main"]