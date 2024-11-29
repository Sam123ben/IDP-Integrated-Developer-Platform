# Use the official Golang image as the base image for building the application
FROM golang:1.23.3 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o terraform-app main.go

# Use a lightweight base image for the runtime
FROM debian:bullseye-slim

# Set the working directory in the container
WORKDIR /app

# Install necessary dependencies, including Terraform
RUN apt-get update && \
    apt-get install -y wget unzip && \
    wget https://releases.hashicorp.com/terraform/1.5.7/terraform_1.5.7_linux_amd64.zip && \
    unzip terraform_1.5.7_linux_amd64.zip -d /usr/local/bin/ && \
    rm terraform_1.5.7_linux_amd64.zip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy the built application from the builder stage
COPY --from=builder /app/terraform-app /app/terraform-app

# Copy necessary configuration files and templates
COPY --from=builder /app/configs/ /app/configs/
COPY --from=builder /app/templates/ /app/templates/

# Expose any required ports (if applicable)
# EXPOSE 8080

# Set the entrypoint to the application binary
ENTRYPOINT ["/app/terraform-app"]