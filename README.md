# Infra Environment Dashboard

Infra Environment Dashboard is a web application that helps teams manage and visualize infrastructure environments across different infra types like **INTERNAL** (for internal development) and **CUSTOMER** (for customer-specific environments). The dashboard is highly configurable using the database, meaning if the correct data is entered into the database, it will automatically flow into the dashboard without additional configuration. The dashboard makes it easy to see each environment’s status, grouped by products or customers. The dashboard is highly configurable using the database, and as long as the data is correct in the database, the information will seamlessly flow into the UI for easy visualization.

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

```plaintext
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

### Prerequisites

- **Docker** and **Docker Compose** installed on your system.

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/infra_env_dashboard.git
   cd infra_env_dashboard
   ```

2. **Setup environment variables** (Optional):
   - You may adjust database user and password in `docker-compose.yml` if needed.

3. **Run the application**:
   ```bash
   docker-compose up --build
   ```

   This will build and run the services:
   - **PostgreSQL Database** on port `5432`
   - **Company Service** (backend) on port `8081`
   - **Environment Service** (backend) on port `8082`
   - **Frontend** on port `3000`

4. **Access the dashboard**:
   Open [http://localhost:3000](http://localhost:3000) in your web browser.

### Project Components

- **Backend Services**:
  - `fetch_company_details`: Provides details about the company.
  - `fetch_infra_types`: Lists infra types, sections, and groups for managing different environments.
- **Frontend**:
  - Built with React. This provides a user interface to display and manage all environments based on infra type.

### Database Setup

- **Database**: PostgreSQL is used as the database for storing company, infra types, sections, environment groups, and environments.
- **Schema**: The `database/000_create_database_schema.sql` file initializes the database tables and relations.

### API Endpoints

Each backend service has its own API endpoints:

1. **Company Service** (`fetch_company_details`):
   - `GET /company` - Retrieves company details.

2. **Environment Service** (`fetch_infra_types`):
   - `GET /infra_types` - Fetches infra types, sections, and groups.

Check the API documentation (Swagger) for detailed endpoint information.

### Environment Variables

- **Best Practices for Credentials**: For local development, you can use `config.yaml` to store database connection details and other configurations. However, for production environments, it is highly recommended to use secure methods like **HashiCorp Vault**, **AWS Secrets Manager**, or another secrets management service to avoid storing sensitive information (such as passwords or API keys) directly in code or configuration files. This helps ensure that true production passwords are not leaked or stored in Git.

- **DATABASE_URL**: Used by backend services to connect to the PostgreSQL database. Set up automatically in `docker-compose.yml`.

### File Structure

- **backend/common**: Common files like database connection (`db.go`), HTTP service handlers (`http_service.go`), and logging setup (`logger.go`).
- **backend/services**: Each service has its own folder with code, handlers, models, repository, and routing.
- **frontend**: Contains the React code for the dashboard UI.

### Contributing

If you’d like to contribute:

1. Fork the repo
2. Create a new branch (`git checkout -b feature-name`)
3. Commit changes (`git commit -am 'Add new feature'`)
4. Push the branch (`git push origin feature-name`)
5. Open a pull request

### License

This project is licensed under the MIT License.