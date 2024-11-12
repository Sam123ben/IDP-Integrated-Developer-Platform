# Infra Environment Dashboard

[![Dashboard Infra Build](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/manual:%20Build-Dashboard-Infra.yaml/badge.svg)](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/manual:%20Build-Dashboard-Infra.yaml)
[![Backend Services CI](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Backend-Services.yaml/badge.svg)](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Backend-Services.yaml)
[![Frontend Services CI](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Frontend-Services.yaml/badge.svg)](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Frontend-Services.yaml)

Infra Environment Dashboard is a web application that helps teams manage and visualize infrastructure environments across different infra types like **INTERNAL** (for internal development) and **CUSTOMER** (for customer-specific environments). The dashboard is highly configurable using the database, meaning if the correct data is entered into the database, it will automatically flow into the dashboard without additional configuration. The dashboard makes it easy to see each environment’s status, grouped by products or customers.

## Features

![Infra Environment Dashboard Screenshot](./docs/images/Devops-Dashboard.png)

- **Company Service**: Provides company information.
- **Environment Service**: Lists environment types, sections, and groups under INTERNAL and CUSTOMER infra types.
- **Frontend Dashboard**: User-friendly interface to view and manage environments in one place.

## Project Structure

The project is divided into multiple services and components:

```plaintext
# Project Folder Structure

This document provides a detailed overview of the project's folder structure. Each directory and file is organized to enhance modularity, readability, and maintainability.

.
├── README.md                     # Main project documentation
├── docs                          # Documentation assets
│   └── images                    # Images used in documentation
│       ├── Deployment-Version.png
│       ├── Devops-Dashboard.png
│       └── dashboard.png
├── infra_env_dashboard           # Primary application directory
│   ├── README.md                 # README for infra_env_dashboard specifics
│   ├── backend                   # Backend services and configurations
│   │   ├── common                # Common utilities and services used across backend
│   │   │   ├── configs           # Configuration files
│   │   │   │   └── config.yaml   # Main configuration file
│   │   │   ├── httpservice       # HTTP service handling
│   │   │   │   └── http_service.go
│   │   │   ├── logger            # Logger utility for backend services
│   │   │   │   └── logger.go
│   │   │   ├── postgress         # PostgreSQL database connection handler
│   │   │   │   └── db.go
│   │   │   └── utils             # Utility functions for error handling, etc.
│   │   │       └── error_handler.go
│   │   ├── go.mod                # Go module file
│   │   ├── go.sum                # Go dependencies
│   │   └── services              # Backend services for specific features
│   │       ├── fetch_company_details # Service to fetch company details
│   │       │   ├── Dockerfile    # Dockerfile for containerizing this service
│   │       │   ├── docs          # Documentation for this service
│   │       │   │   ├── docs.go
│   │       │   │   ├── swagger.json
│   │       │   │   └── swagger.yaml
│   │       │   ├── handlers      # API endpoint handlers
│   │       │   │   └── company_handler.go
│   │       │   ├── main.go       # Entry point for the service
│   │       │   ├── models        # Models related to company details
│   │       │   │   ├── company.go
│   │       │   │   └── response.go
│   │       │   ├── repository    # Database operations
│   │       │   │   └── company_repository.go
│   │       │   └── router        # Router configuration
│   │       │       └── router.go
│   │       ├── fetch_infra_types # Service to fetch infrastructure types
│   │       │   ├── Dockerfile
│   │       │   ├── docs
│   │       │   │   ├── docs.go
│   │       │   │   ├── swagger.json
│   │       │   │   └── swagger.yaml
│   │       │   ├── handlers
│   │       │   │   └── infra_handler.go
│   │       │   ├── main.go
│   │       │   ├── models
│   │       │   │   ├── infra_type.go
│   │       │   │   └── response.go
│   │       │   ├── repository
│   │       │   │   └── infra_repository.go
│   │       │   └── router
│   │       │       └── router.go
│   │       └── fetch_internal_env_details # Service to fetch internal environment details
│   │           ├── Dockerfile
│   │           ├── handlers
│   │           │   └── handler.go
│   │           ├── main.go
│   │           ├── models
│   │           │   ├── models.go
│   │           │   └── response.go
│   │           ├── repository
│   │           │   └── repository.go
│   │           └── router
│   │               └── router.go
│   ├── database                   # Database setup scripts
│   │   └── 000_create_database_schema.sql # SQL script to create initial schema
│   ├── docker-compose.yml         # Docker Compose for the entire application stack
│   └── frontend                   # Frontend application
│       ├── Dockerfile             # Dockerfile for frontend containerization
│       ├── assets                 # Static assets (images, fonts, etc.)
│       ├── dist                   # Compiled frontend assets
│       │   ├── bundle.js
│       │   ├── bundle.js.LICENSE.txt
│       │   └── index.html
│       ├── package-lock.json      # Node package lock file
│       ├── package.json           # Node package configuration
│       ├── public                 # Public static files
│       │   └── config
│       ├── src                    # Source code for the frontend
│       │   ├── App.js             # Main application component
│       │   ├── components         # Reusable React components
│       │   │   ├── AppVersionModal.js
│       │   │   ├── Card.js
│       │   │   ├── Footer.js
│       │   │   ├── Header.js
│       │   │   ├── MainContent.js
│       │   │   ├── Modal.js
│       │   │   ├── PrivacyModal.js
│       │   │   ├── SectionHeader.js
│       │   │   ├── Sidebar.js
│       │   │   └── TileContainer.js
│       │   ├── config.js          # Configuration file for frontend settings
│       │   ├── index.css          # Main CSS file
│       │   ├── index.html         # HTML template
│       │   ├── index.js           # Entry point for React
│       │   ├── services           # API services for frontend
│       │   │   └── api.js
│       │   └── styles             # CSS files for styling components
│       │       ├── App.css
│       │       ├── AppVersionModal.css
│       │       ├── Card.css
│       │       ├── Footer.css
│       │       ├── Header.css
│       │       ├── MainContent.css
│       │       ├── Modal.css
│       │       ├── PrivacyModal.css
│       │       ├── SectionHeader.css
│       │       ├── Sidebar.css
│       │       └── TileContainer.css
│       └── webpack.config.js      # Webpack configuration
└── tree.txt                       # File listing of the project structure
```

## Getting Started

## Prerequisites
Ensure the following are installed on your system:
- **Docker** and **Docker Compose**
- **Go (Golang)** for running backend services
- **Node.js** and **npm** for frontend

---

## Option 1: Manually Start Each Service Locally

### 1. Clone the Repository

```bash
git clone https://github.com/Sam123ben/IDP-Integrated-Developer-Platform.git
cd infra_env_dashboard
```

### 2. Set Up and Start PostgreSQL Database
Run the following Docker command to start a PostgreSQL container locally:

```
docker run --name my_postgres \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  -d postgres
```
> **Note**: This command starts the PostgreSQL process locally without a persistent volume. To retain data, use the `-v` option to map a local directory to the Docker container (e.g., `-v /your/local/dir:/var/lib/postgresql/data`).

To enter the running PostgreSQL service and interact with the database, execute:

```bash
docker exec -it my_postgres psql -U myuser -d mydatabase
```

### 3. Initialize Database Schema and Test Data
To enter the running PostgreSQL service and interact with the database, execute:

```bash
docker exec -it my_postgres psql -U myuser -d mydatabase
```

Once connected, you can initialize the database schema and add some quick test data. Open the SQL script file `database/000_create_database_schema.sql`, copy the SQL commands, and paste them into the PostgreSQL terminal.

This will create the necessary tables and relationships for the application, and you can also add dummy data for testing purposes.

### 4. Start Backend Services
Navigate to the backend directory:

```bash
cd infra_env_dashboard/backend
```

Start each backend service individually using Go:

#### Backend Service

```bash
go run main.go
```

Each service will start and listen on its designated port (e.g., Company Service on port 8081, Environment Service on port 8082).

### 5. Start the Frontend
Navigate to the frontend directory:

```bash
cd infra_env_dashboard/frontend
```

Install dependencies if needed:

```bash
npm install
```

Start the frontend development server:

```bash
npm start
```

The frontend will be accessible at [http://localhost:3000](http://localhost:3000).

### Testing the Dashboard
1. Open [http://localhost:3000](http://localhost:3000) in your browser.
2. Check for correct display of environments grouped by infra type (e.g., **INTERNAL** or **CUSTOMER**).
3. Test each backend service's integration with the frontend by verifying data from each service is displayed correctly in the UI.

## Option 2: Using Docker Compose (Future Testing)
In the future, you can use Docker Compose to simplify starting up all services. The `docker-compose.yml` file is already configured for this:

1. Adjust any environment variables in `docker-compose.yml` as needed.

2. Run the following command to build and start all services together:

   ```bash
   docker-compose up --build
   ```
This will automatically start:

- **PostgreSQL Database** on port 5432
- **Company Service** on port 8081
- **Environment Service** on port 8082
- **Frontend** on port 3000

Access the dashboard at [http://localhost:3000](http://localhost:3000) to test it.

## Infrastructure Automation - Terraform

For detailed instructions on setting up and managing the infrastructure using Terraform, please refer to the [Infrastructure Automation README](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/tree/main/infra_env_dashboard/infra-automation#readme).

## Future Enhancements

To improve the quality and reliability of the application, the following enhancements are planned:

### 1. Implement Unit Testing Frameworks

#### Backend (Go):

- **Framework**: Utilize Go's built-in `testing` package along with `testify` for assertions and mocking.
- **Description**: Write unit tests for backend services to ensure individual components function as expected.
- **Tasks**:
  - Set up the testing environment in each backend service.
  - Write unit tests for models, handlers, and repository functions.
  - Integrate tests into the build pipeline.

#### Frontend (React):

- **Framework**: Implement `Jest` and `React Testing Library` for testing React components.
- **Description**: Write unit tests for React components to verify rendering, user interaction, and state management.
- **Tasks**:
  - Configure Jest and React Testing Library in the frontend project.
  - Write tests for critical components like `Header`, `Footer`, `MainContent`, etc.
  - Ensure coverage for various user interaction scenarios.

### 2. Introduce Integration Testing

#### Backend Integration Tests:

- **Framework**: Use tools like `GoConvey` or continue with Go's `testing` package for integration tests.
- **Description**: Write tests that verify the interaction between different parts of the backend services and the database.
- **Tasks**:
  - Set up integration test suites that run against a test database.
  - Write tests that cover API endpoints and database transactions.
  - Mock external services if necessary.

#### Frontend Integration Tests:

- **Framework**: Utilize `Cypress` or `Selenium` for end-to-end testing.
- **Description**: Write tests that simulate user interactions with the UI and verify the application flow.
- **Tasks**:
  - Configure Cypress in the frontend project.
  - Write end-to-end tests that cover critical user journeys.
  - Integrate these tests into the CI/CD pipeline.

### 3. Continuous Integration and Continuous Deployment (CI/CD)

- **Integrate Testing into CI/CD Pipeline**:
  - **Description**: Ensure that all tests are automatically run during the build and deployment process.
  - **Tasks**:
    - Set up GitHub Actions, Jenkins, or another CI tool to run tests on every push and pull request.
    - Fail builds if any tests fail.
    - Generate and store test coverage reports.

### 4. Documentation and Code Quality

#### Improve Code Documentation:

- **Description**: Add comments and documentation to the codebase for better maintainability.
- **Tasks**:
  - Add GoDoc comments to all public Go functions and types.
  - Document React components and their props.

#### Enforce Coding Standards:

- **Backend**:
  - Use `golint` and `go vet` to enforce coding standards.
- **Frontend**:
  - Implement `ESLint` and `Prettier` for consistent code formatting and linting.

## Contributing

If you’d like to contribute:

1. Clone the repo
2. Create a new branch (`git checkout -b feature-name`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push the branch (`git push origin feature-name`)
5. Open a pull request

## License

This project is licensed under the MIT License.