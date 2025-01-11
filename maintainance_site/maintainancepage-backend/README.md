# Backend Service README

## Overview
This folder contains the **backend service** for the Maintenance Page System. It provides API endpoints, database interactions, and middleware for handling system components, maintenance windows, and updates. The backend is written in **Go** and leverages **PostgreSQL** for data storage. It also includes utilities for database migrations and seeding.

This README explains the folder structure, key files, and how to run the backend locally (with and without Docker).

---

## Folder Structure

### **Root Folder**
- **`Dockerfile`**: Docker configuration file to containerize the backend service.
- **`go.mod`**: Specifies the project dependencies and Go version.
- **`go.sum`**: Dependency checksum file for integrity verification.

### **`cmd/`**
Contains entry points for specific tasks.
- **`migrate/main.go`**: Handles database schema migrations.
- **`seed/main.go`**: Seeds the database with initial dummy data.
- **`server/main.go`**: Entry point for starting the backend server.

### **`data/`**
- **`dummy_data.sql`**: SQL script containing dummy data for testing and development purposes.

### **`internal/`**
This directory contains the core logic and modules for the backend.

#### **`config/`**
- **`config.go`**: Handles application-level configurations, such as database connection details and environment variables.

#### **`cron/`**
- **`cleanup.go`**: Contains scheduled tasks, such as database cleanup.

#### **`database/`**
- **`database.go`**: Contains logic for initializing and interacting with the PostgreSQL database.

#### **`handlers/`**
Defines API endpoints for managing system components and maintenance updates.
- **`maintenance.go`**: Handlers for creating, updating, and retrieving maintenance updates.
- **`system_components.go`**: Handlers for managing system component statuses.

#### **`middleware/`**
- **`middleware.go`**: Defines middleware for request validation, logging, and security.

#### **`models/`**
- **`models.go`**: Defines the data models and database schema for the backend.
- **`response.go`**: Defines the structure of API responses.

#### **`router/`**
- **`router.go`**: Configures the API routes and maps them to their respective handlers.

---

## Key Features

### API Endpoints
The backend exposes RESTful API endpoints for:
- Managing system components (e.g., updating statuses).
- Creating and updating maintenance windows.
- Retrieving maintenance updates.

### Database Management
- Supports database migrations and seeding.
- Automatically initializes tables and relationships on startup.

### Middleware
- Implements request validation and logging.
- Ensures secure access to endpoints.

### Cron Jobs
- Scheduled tasks for periodic database cleanup and maintenance.

---

## Running Locally

### **Prerequisites**
- **Go** installed on your local machine.
- **PostgreSQL** installed and running.
- **Docker** (if running with containers).

### **With Docker**
1. Build the Docker image:
   ```bash
   docker build -t maintenance-backend .
   ```
2. Run the container:
   ```bash
   docker run -p 8080:8080 -e DATABASE_URL=your_database_url maintenance-backend
   ```
3. Access the API at [http://localhost:8080](http://localhost:8080).

### **Without Docker**
1. Set up environment variables:
   - Create a `.env` file with the following:
     ```env
     DATABASE_URL=your_database_url
     ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the server:
   ```bash
   go run cmd/server/main.go
   ```
4. Access the API at [http://localhost:8080](http://localhost:8080).

---

## Key Notes

### **Database Migrations**
To apply database migrations:
```bash
go run cmd/migrate/main.go
```

### **Seeding the Database**
To seed the database with dummy data:
```bash
go run cmd/seed/main.go
```

### **API Endpoints**
- **GET /system-components**: Retrieve all system components.
- **POST /system-components**: Add a new system component.
- **PUT /system-components**: Update an existing system component.
- **POST /maintenance-windows**: Create a new maintenance window.
- **GET /maintenance-windows**: Retrieve all maintenance windows.

### **Testing CORS**
Ensure CORS policies are configured correctly if integrating with a frontend application.

---

## Conclusion
This backend service provides a robust and scalable foundation for managing system maintenance. It is designed for performance and ease of integration, making it suitable for both development and production environments. Let us know if you need additional assistance!