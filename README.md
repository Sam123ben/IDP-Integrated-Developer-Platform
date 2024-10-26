# IDP-Integrated-Developer-Platform

## Unified Developer Hub

## Overview

Welcome to the **IDP - Integrated Developer Platform** project! This project is a microservice-based application with **React** as the frontend, **Go** for backend microservices, **Kubernetes** as the orchestrator, and **PostgreSQL** as the primary database. The IDP platform acts as a unified layer on top of cloud providers, enabling development teams to deploy and manage multiple environments effortlessly. With **Crossplane** for infrastructure-as-code (IaC) automation, this platform allows for both **persistent** and **ephemeral** feature environments.

## Project Goals

The main objectives of the IDP are:

1. **Simplify Environment Creation**: Developers can spin up production-like and feature-specific environments with minimal friction.
2. **Automate Infrastructure Management**: With Crossplane, infrastructure resources are provisioned and managed automatically, ensuring consistency and reliability.
3. **Enhance Deployment Flexibility**: By leveraging Kubernetes and Crossplane, developers can deploy across different cloud providers in a consistent manner.
4. **Optimize for CI/CD Workflows**: The IDP supports ephemeral environments, making it easy to test features in isolation within CI/CD pipelines.

## Key Features

- **Environment Management**: Provision both true (persistent) and ephemeral environments.
- **Cloud-Agnostic Orchestration**: Kubernetes and Crossplane facilitate a consistent cloud-agnostic experience.
- **Self-Service Interface**: A React-based UI that allows teams to manage their environments, deploy services, and oversee infrastructure resources.
- **Automated IaC**: Crossplane is integrated to provide IaC automation, reducing the need for manual resource creation.

## Architecture

The IDP follows a microservices architecture with the following primary components:

1. **Frontend (React)**: A React-based UI that developers use to interact with the platform, view environment statuses, and trigger deployments.
2. **Backend (Go)**: Multiple Go microservices handle core functionalities, such as environment provisioning, resource management, and API interactions.
3. **Database (PostgreSQL)**: Stores data about environments, configurations, deployments, and other metadata.
4. **Orchestration (Kubernetes)**: Kubernetes manages the deployment, scaling, and networking of services across the platform.
5. **Infrastructure Automation (Crossplane)**: Crossplane is configured to manage IaC workflows, automatically provisioning and deprovisioning infrastructure resources based on platform requirements.

## Folder Structure

To maintain a clean and modular project, here’s a recommended folder structure:

```
IDP-Integrated-Developer-Platform/
├── backend/                     # Backend services and microservices in Go
│   ├── api/                     # API endpoints and handlers
│   ├── models/                  # Database models and ORM structures
│   ├── services/                # Business logic and service layer
│   ├── rbac/                    # Role-based access control code
│   ├── config/                  # Configuration files (e.g., DB, environment)
│   └── main.go                  # Main entry point for the backend application
│
├── frontend/                    # React-based frontend for the platform
│   ├── src/
│       ├── components/          # Reusable UI components
│       ├── pages/               # Main pages and views
│       ├── services/            # Frontend services to interact with backend APIs
│       ├── hooks/               # Custom React hooks
│       └── App.js               # Main entry file for the React app
│
├── infra/                       # Infrastructure as code resources
│   ├── kubernetes/              # Kubernetes configuration and manifests
│   ├── crossplane/              # Crossplane configuration for cloud infrastructure
│   ├── terraform/               # Terraform modules for additional IaC needs
│   └── secrets/                 # Encrypted secrets management (e.g., using Vault)
│
├── workflows/                   # CI/CD and GitHub Actions workflows
│   ├── deploy.yaml              # Deployment workflow file
│   ├── test.yaml                # Testing workflow file
│   └── build.yaml               # Build workflow file
│
├── docs/                        # Documentation for the platform
│   ├── architecture.md          # Detailed architecture document
│   ├── api-reference.md         # API reference for backend services
│   ├── rbac.md                  # Documentation on RBAC configurations
│   └── contributing.md          # Contribution guidelines
│
├── README.md                    # Project overview and setup guide
├── LICENSE                      # License for the project
└── .gitignore                   # Git ignore file
```

### Folder Descriptions

- **\`backend/\`**: Contains Go microservices, API endpoints, and RBAC implementation.
- **\`frontend/\`**: Holds the React application with UI components, pages, and services.
- **\`infra/\`**: Infrastructure code, split into Kubernetes manifests, Crossplane configurations, and Terraform modules.
- **\`workflows/\`**: CI/CD workflows, including files for testing, building, and deploying the project.
- **\`docs/\`**: Project documentation, architecture notes, API references, and contribution guidelines.

## Technology Stack

| Component                | Technology    |
|--------------------------|---------------|
| Frontend                 | React         |
| Backend                  | Go (Golang)   |
| Database                 | PostgreSQL    |
| Orchestrator             | Kubernetes    |
| IaC Automation           | Crossplane, Terraform |
| Containerization         | Docker        |

## Getting Started

### Prerequisites

To set up and run the IDP locally or in a development environment, you’ll need the following:

- **Docker**: For containerization
- **Kubernetes Cluster**: Can be a local Kubernetes (e.g., Minikube or Kind) or a cloud-managed cluster
- **Crossplane**: For IaC automation (install and configure in the Kubernetes cluster)
- **PostgreSQL**: Database instance, either local or cloud-hosted
- **Go** (1.16 or newer): Required for backend development
- **Node.js** (14.x or newer): Required for frontend development
- **kubectl**: Kubernetes CLI for deployment management
- **Cloud Provider Credentials**: Set up as needed for Crossplane to manage infrastructure on specific cloud providers

### Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Sam123ben/IDP-Integrated-Developer-Platform.git
   cd IDP-Integrated-Developer-Platform
   ```

2. **Setup Kubernetes and Crossplane**:
   - Ensure your Kubernetes cluster is running.
   - Install Crossplane in your Kubernetes cluster:
     ```bash
     kubectl apply -f https://raw.githubusercontent.com/crossplane/crossplane/release-1.0/install.yaml
     ```
   - Configure Crossplane with the necessary provider (e.g., AWS, GCP, Azure).

3. **Database Setup (PostgreSQL)**:
   - Install and start PostgreSQL, or use a managed instance.
   - Create a database for the IDP:
     ```sql
     CREATE DATABASE idp_db;
     ```
   - Add a user with access to this database and update the backend configuration with the database credentials.

4. **Frontend Setup**:
   - Navigate to the frontend directory:
     ```bash
     cd frontend
     npm install
     npm start
     ```
   - This will start the React application on \`http://localhost:3000\`.

5. **Backend Setup**:
   - Navigate to the backend directory:
     ```bash
     cd ../backend
     go mod download
     ```
   - Configure the backend with the PostgreSQL connection string:
     ```bash
     export DATABASE_URL="postgres://username:password@localhost:5432/idp_db"
     ```
   - Run the backend:
     ```bash
     go run main.go
     ```

6. **Deploy to Kubernetes**:
   - Once both services are ready, deploy them to Kubernetes using the manifests provided in the \`/infra/kubernetes\` directory.
   - Apply the configuration:
     ```bash
     kubectl apply -f infra/kubernetes/
     ```

## Usage

- **True Environments**: Developers can create persistent environments via the IDP, intended for long-term or production-like deployments.
- **Ephemeral Feature Environments**: Triggered for each feature branch, allowing developers to test features in isolated, short-lived environments.
- **Infrastructure as Code**: Through Crossplane, manage infrastructure resources across cloud providers, facilitating a consistent multi-cloud environment.

## Contributing

Contributions are welcome! To contribute:

1. Fork the repository.
2. Create a feature branch (\`git checkout -b feature-name\`).
3. Commit your changes (\`git commit -m 'Add feature'\`).
4. Push to the branch (\`git push origin feature-name\`).
5. Open a pull request.

Please refer to our [CONTRIBUTING.md](CONTRIBUTING.md) file for more details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

## Contact

For any questions, please feel free to reach out to the project maintainer:

- **Email**: [your-email@example.com](mailto:your-email@example.com)
- **Slack**: IDP Development Channel (link to Slack workspace if applicable)

---
