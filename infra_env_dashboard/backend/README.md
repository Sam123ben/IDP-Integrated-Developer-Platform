# Backend API Documentation

## Overview
This backend provides APIs to fetch company details, customer environment details, and internal environment details. Each service has specific endpoints to handle data retrieval and updates, enabling efficient management of infrastructure and environments.

---

## Table of Contents
1. [fetch_company_details](#fetch_company_details)
2. [fetch_customer_env_details](#fetch_customer_env_details)
3. [fetch_internal_env_details](#fetch_internal_env_details)
4. [Setup and Installation](#setup-and-installation)
5. [Running the Application](#running-the-application)
6. [Contact and Support](#contact-and-support)

---

## fetch_company_details

### Purpose
Fetches company details, including the company name and other metadata.

### Endpoint

#### `GET /api/company-details`
Fetches the company name and related metadata.

- **Response**
  - **Status 200**: Successfully retrieves the company details.
  - **Status 500**: Internal server error.

**Example Response**
```json
{
  "company": {
    "id": 1,
    "name": "My Company"
  }
}
```

---

## fetch_customer_env_details

### Purpose
Manages customer-specific environment details, including fetching, updating, or adding environment data.

### Endpoints

#### `GET /api/customer-env-details`
Fetches environment details for a specific customer and product.

**Query Parameters**
- `customer` (string, required): Customer name.
- `product` (string, required): Product name.

**Response**
- **Status 200**: Successfully retrieves environment details.
- **Status 400**: Missing query parameters.
- **Status 500**: Internal server error.

**Example Request**
```http
GET /api/customer-env-details?customer=VendorA&product=Product1
```

**Example Response**
```json
{
  "environmentDetails": [
    {
      "id": 1,
      "name": "QA Environment",
      "url": "https://qa.example.com",
      "status": "Online",
      "contact": "John Doe",
      "appVersion": "v1.2.3",
      "dbVersion": "v2.3.4",
      "lastUpdated": "2024-11-17T12:34:56",
      "comments": "Testing environment for Vendor A, Product 1."
    }
  ]
}
```

---

#### `PUT /api/customer-env-details`
Updates or inserts a new customer environment.

**Request Body**
```json
{
  "id": 1,
  "customer_name": "Vendor A",
  "product_name": "Product 1",
  "name": "QA Environment",
  "url": "https://qa.example.com",
  "lastUpdated": "2024-11-17T12:34:56",
  "status": "Online",
  "contact": "John Doe",
  "appVersion": "v1.2.3",
  "dbVersion": "v2.3.4",
  "comments": "Testing environment for Vendor A, Product 1."
}
```

**Response**
- **Status 200**: Environment details updated or added successfully.
- **Status 400**: Invalid request payload or parameters.
- **Status 500**: Internal server error.

**Example Response**
```json
{
  "message": "Environment details updated successfully."
}
```

---

## fetch_internal_env_details

### Purpose
Handles internal environment details for products, including fetching, updating, and managing specific environments.

### Endpoints

#### `GET /api/internal-env-details`
Fetches environment details for a specific product and environment group.

**Query Parameters**
- `product` (string, required): Product name.
- `group` (string, required): Environment group (e.g., DEV, QA, PROD).

**Response**
- **Status 200**: Successfully retrieves environment details.
- **Status 400**: Missing or invalid query parameters.
- **Status 500**: Internal server error.

**Example Request**
```http
GET /api/internal-env-details?product=Product1&group=QA
```

**Example Response**
```json
{
  "environmentDetails": [
    {
      "id": 4,
      "name": "QA Environment",
      "url": "https://qa.internal.example.com",
      "status": "In Progress",
      "contact": "Alice",
      "appVersion": "v1.0.0",
      "dbVersion": "v1.5.2",
      "lastUpdated": "2024-11-17T12:00:00",
      "comments": "Automated QA environment."
    }
  ]
}
```

---

#### `PUT /api/internal-env-details`
Updates or inserts internal environment details.

**Request Body**
```json
{
  "id": 4,
  "product_name": "Product 1",
  "environment_group": "QA",
  "name": "QA Environment",
  "url": "https://qa.internal.example.com",
  "lastUpdated": "2024-11-17T12:00:00",
  "status": "Online",
  "contact": "Alice",
  "appVersion": "v1.0.0",
  "dbVersion": "v1.5.2",
  "comments": "Updated QA environment details."
}
```

**Response**
- **Status 200**: Environment details updated or added successfully.
- **Status 400**: Invalid request payload or parameters.
- **Status 500**: Internal server error.

**Example Response**
```json
{
  "message": "Internal environment details updated successfully."
}
```

---

## Setup and Installation

### Prerequisites
- Node.js
- Go (Golang) for backend development
- PostgreSQL or another supported database

### Clone Repository
```bash
git clone https://github.com/your-organization/your-repo.git
cd your-repo/backend
```

### Configure Environment
1. Copy `.env.example` to `.env`.
2. Update database connection details and other configurations in `.env`.

### Install Dependencies
```bash
go mod tidy
```

### Migrate Database
Use your migration tool (e.g., Flyway or Go-Migrate) to apply migrations.

---

## Running the Application

### Start the Backend Server
```bash
go run main.go
```

### API Documentation
If you're using Swagger:
- Access the Swagger UI at `http://localhost:8080/swagger/index.html`.

---

## Contact and Support
For support or questions, please reach out to:

- **Email**: support@yourcompany.com
- **Documentation**: [Your Documentation Link](https://docs.yourcompany.com)
- **Issues**: [GitHub Issues](https://github.com/your-organization/your-repo/issues)

---

This **README.md** provides a detailed, easy-to-follow guide for users and developers to understand and work with the APIs. The documentation also includes endpoint descriptions, request examples, and setup instructions to help get you started quickly.