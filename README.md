# Infra Environment Dashboard

| Status Badge | Description |
|--------------|-------------|
| [![Dashboard Infra Build](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/manual:%20Build-Dashboard-Infra.yaml/badge.svg)](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/manual:%20Build-Dashboard-Infra.yaml) | Dashboard Infra Build Status |
| [![Backend Services CI](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Backend-Services.yaml/badge.svg)](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Backend-Services.yaml) | Backend Services CI Status |
| [![Frontend Services CI](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Frontend-Services.yaml/badge.svg)](https://github.com/Sam123ben/IDP-Integrated-Developer-Platform/actions/workflows/ci:%20Build-Frontend-Services.yaml) | Frontend Services CI Status |

Infra Environment Dashboard is a web application that helps teams manage and visualize infrastructure environments across different infra types like **INTERNAL** (for internal development) and **CUSTOMER** (for customer-specific environments). The dashboard is highly configurable using the database, meaning if the correct data is entered into the database, it will automatically flow into the dashboard without additional configuration. The dashboard makes it easy to see each environment’s status, grouped by products or customers.

## Features

![Infra Environment Dashboard Screenshot](./docs/images/Devops-Dashboard.png)

- **Company Service**: Provides company information.
- **Environment Service**: Lists environment types, sections, and groups under INTERNAL and CUSTOMER infra types.
- **Frontend Dashboard**: User-friendly interface to view and manage environments in one place.

## Project Structure

The project is divided into multiple services and components:

```
# Updated Project Folder Structure

.
├── README.md                     # Main project documentation
├── docs                          # Documentation assets
│   └── images                    # Images used in documentation
│       ├── Deployment-Version.png
│       ├── Devops-Dashboard.png
│       └── dashboard.png
└── infra_env_dashboard           # Primary application directory
    ├── README.md                 # README for infra_env_dashboard specifics
    ├── backend                   # Backend services and configurations
    │   ├── Dockerfile            # Dockerfile for backend containerization
    │   ├── README.md             # Backend-specific documentation
    │   ├── common                # Common utilities and services used across backend
    │   │   ├── configs           # Configuration files
    │   │   │   └── config.yaml
    │   │   ├── httpservice       # HTTP service handling
    │   │   │   └── http_service.go
    │   │   ├── logger            # Logger utility for backend services
    │   │   │   └── logger.go
    │   │   ├── postgress         # PostgreSQL database connection handler
    │   │   │   └── db.go
    │   │   └── utils             # Utility functions for error handling, etc.
    │   │       └── error_handler.go
    │   ├── docs                  # Swagger and other backend documentation
    │   │   ├── docs.go
    │   │   ├── swagger.json
    │   │   └── swagger.yaml
    │   ├── go.mod                # Go module file
    │   ├── go.sum                # Go dependencies
    │   ├── main.go               # Entry point for backend services
    │   └── services              # Backend services for specific features
    │       ├── fetch_company_details # Service to fetch company details
    │       │   ├── handlers
    │       │   │   └── company_handler.go
    │       │   ├── models
    │       │   │   ├── company.go
    │       │   │   └── response.go
    │       │   ├── repository
    │       │   │   └── company_repository.go
    │       │   └── router
    │       │       └── router.go
    │       ├── fetch_customer_env_details # Service to fetch customer environment details
    │       │   ├── handlers
    │       │   │   └── handler.go
    │       │   ├── models
    │       │   │   ├── models.go
    │       │   │   └── response.go
    │       │   ├── repository
    │       │   │   └── repository.go
    │       │   └── router
    │       │       └── router.go
    │       └── fetch_internal_env_details # Service to fetch internal environment details
    │           ├── handlers
    │           │   └── handler.go
    │           ├── models
    │           │   ├── models.go
    │           │   └── response.go
    │           ├── repository
    │           │   └── repository.go
    │           └── router
    │               └── router.go
    ├── database                   # Database setup scripts
    │   └── 000_create_database_schema.sql # SQL script to create initial schema
    ├── docker-compose.yml         # Docker Compose for the entire application stack
    ├── frontend                   # Frontend application (TypeScript)
    │   ├── Dockerfile             # Dockerfile for frontend containerization
    │   ├── assets                 # Static assets (images, fonts, etc.)
    │   ├── dist                   # Compiled frontend assets
    │   │   ├── bundle.js
    │   │   ├── bundle.js.LICENSE.txt
    │   │   └── index.html
    │   ├── package.json           # Node package configuration
    │   ├── src                    # Source code for the frontend (converted to TypeScript)
    │   │   ├── App.tsx            # Main application component
    │   │   ├── components         # Reusable React components in TypeScript
    │   │   │   ├── AppVersionModal.tsx
    │   │   │   ├── Card.tsx
    │   │   │   ├── CardMenu.tsx
    │   │   │   ├── Footer.tsx
    │   │   │   ├── Header.tsx
    │   │   │   ├── MainContent.tsx
    │   │   │   ├── Modal.tsx
    │   │   │   ├── PrivacyModal.tsx
    │   │   │   ├── Sidebar.tsx
    │   │   │   └── TileContainer.tsx
    │   │   ├── config.ts          # TypeScript configuration file for frontend settings
    │   │   ├── index.css          # Main CSS file
    │   │   ├── index.html         # HTML template
    │   │   ├── index.tsx          # Entry point for React (TypeScript)
    │   │   ├── services           # API services for frontend (converted to TypeScript)
    │   │   │   ├── api.tsx
    │   │   │   └── fetchData.tsx
    │   │   └── styles             # CSS files for styling components
    │   │       ├── App.css
    │   │       ├── AppVersionModal.css
    │   │       ├── Card.css
    │   │       ├── CardMenu.css
    │   │       ├── Footer.css
    │   │       ├── Header.css
    │   │       ├── MainContent.css
    │   │       ├── Modal.css
    │   │       ├── PrivacyModal.css
    │   │       ├── Sidebar.css
    │   │       └── TileContainer.css
    │   ├── tsconfig.json          # TypeScript configuration for frontend
    │   └── webpack.config.js      # Webpack configuration
    └── infra-automation           # Terraform scripts for infrastructure management
        ├── README.md              # Documentation for infrastructure automation
        ├── backend.tf             # Backend configuration for Terraform
        ├── environments           # Environment-specific configurations (dev/prod)
        │   ├── dev
        │   │   └── main.tf
        │   └── prod
        │       └── main.tf
        ├── main.tf                # Main Terraform configuration
        ├── modules                # Modular Terraform code
        │   ├── app
        │   │   ├── main.tf
        │   │   ├── outputs.tf
        │   │   └── variables.tf
        │   ├── database
        │   │   ├── outputs.tf
        │   │   ├── postgress.tf
        │   │   ├── runsql.tf
        │   │   ├── scripts
        │   │   │   └── 000_create_database_schema.sql
        │   │   └── variables.tf
        │   ├── network
        │   │   ├── bastion.tf
        │   │   ├── nsg.tf
        │   │   ├── outputs.tf
        │   │   ├── subnets.tf
        │   │   ├── variables.tf
        │   │   └── vnet.tf
        │   ├── openvpn
        │   │   ├── main.tf
        │   │   ├── outputs.tf
        │   │   ├── scripts
        │   │   │   └── install_openvpn.sh
        │   │   └── variables.tf
        │   └── resource_group
        │       ├── main.tf
        │       ├── outputs.tf
        │       └── variables.tf
        ├── providers.tf           # Provider configurations for Terraform
        └── variables.tf           # Terraform variables
```

## Getting Started

### Prerequisites
Ensure the following are installed on your system:
- **Docker** and **Docker Compose**
- **Go (Golang)** for running backend services
- **Node.js** and **npm** for frontend
- **TypeScript** for the updated TypeScript frontend components

---

## Option 1: Manually Start Each Service Locally

### 1. Clone the Repository

```bash
git clone https://github.com/Sam123ben/IDP-Integrated-Developer-Platform.git
cd infra_env_dashboard
```

### 2. Set Up and Start PostgreSQL Database
Run the following Docker command to start a PostgreSQL container locally:

```bash
docker run --name my_postgres \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  -d postgres
```

> **Note**: This command starts the PostgreSQL process locally without a persistent volume. To retain data, use the `-v` option to map a local directory to the Docker container (e.g., `-v /your/local/dir:/var/lib/postgresql/data`).

### 3. Initialize Database Schema and Test Data
To enter the running PostgreSQL service and interact with the database, execute:

```bash
docker exec -it my_postgres psql -U myuser -d mydatabase
```

Initialize the database schema and add some quick test data by running the SQL commands in the `database/000_create_database_schema.sql` file.

### 4. Start Backend Services
Navigate to the backend directory:

```bash
cd infra_env_dashboard/backend
```

Start each backend service individually using Go:

```bash
go run main.go
```

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

## Option 2: Using Docker Compose [Not fully tested]
Use Docker Compose to simplify starting up all services:

1. Adjust any environment variables in `docker-compose.yml` as needed.
2. Run:

   ```bash
   docker-compose up --build
   ```

Access the dashboard at [http://localhost:3000](http://localhost:3000).

## Future Enhancements

- **Unit Testing** using Go's `testing` package and `Jest` with `React Testing Library` for the TypeScript components.
- **Integration Testing** with tools like `Cypress`.
- **Documentation and Code Quality Improvements**:
  - **Backend**: Use `golint` and `go vet`.
  - **Frontend**: Implement `ESLint` and `Prettier`.

## Contributing

If you’d like to contribute:

1. Clone the repo.
2. Create a new branch (`git checkout -b feature-name`).
3. Commit changes (`git commit -am 'Add new feature'`).
4. Push the branch (`git push origin feature-name`).
5. Open a pull request.

## License

More Details will come in here.
