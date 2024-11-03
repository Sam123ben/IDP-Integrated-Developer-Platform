# Step 1: Build Stage
FROM golang:1.21 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./

# Download all dependencies (this will be cached if the go.mod and go.sum files are not changed)
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application binary
RUN go build -o main cmd/server/main.go

# Step 2: Run Stage
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Install certificates (for HTTPS support) and other needed packages
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the HTML templates, static assets, and any other needed files from the builder stage
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Set environment variables if needed (e.g., for database connection details)
ENV PORT=8080

# Expose the application's port (8080 in this example)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]