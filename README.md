# Infrastructure & Environment Dashboard

Welcome to the **Infrastructure & Environment Dashboard** project! This is a lightweight and performant web-based dashboard that allows users to **monitor, manage, and optimize their infrastructure**. The application is designed with a **microservice architecture** and follows a 3-tier approach, comprising a **UI**, **middleware**, and **database**.

&#x20;

## Project Overview

This dashboard provides real-time insights into various environments such as Development, Staging, and Production. Instead of full monitoring, it focuses on providing release and deployment statistics, version checks, and details like URLs, skipping deployments, and adding comments for a real-time understanding of the environment test cycles.



The UI and middleware are written in Golang, while PostgreSQL is used as the backend database. The app is designed to be highly modular and easy to extend with more features, such as tracking release statuses, adding or updating comments, managing deployment skips, and capturing details to better understand environment test cycles.Features

- **Real-time Environment Monitoring**: Monitor the status of different environments, including server health and performance metrics.
- **Multi-theme Support**: Easily toggle between **light** and **dark** themes to suit user preferences.
- **Lightweight and Efficient**: Designed using Golang, which makes it highly **efficient** and **scalable**.
- **Microservice Architecture**: Built with scalability in mind, using a microservices approach for easy future expansion.
- **Deployment Management**: The dashboard can be extended to skip deployments, update comments, and trigger rollbacks.

## Technology Stack

- **Frontend & Middleware**: [Golang](https://golang.org/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Containerization**: [Docker](https://www.docker.com/)
- **UI Templates**: HTML5, CSS3
- **Styling & UI Libraries**: Font Awesome for icons
- **Deployment**: Multistage Docker builds for small, production-ready images.

## Project Intentions

The primary intention of this dashboard is to provide a **centralized platform** where DevOps teams and infrastructure engineers can monitor and manage their infrastructure environments efficiently. The project is built to be lightweight, **easily customizable**, and highly **scalable**.

Future plans include integrating more deployment capabilities, advanced analytics, and other environment management features.

## Prerequisites

Before you can run the application, ensure that the following prerequisites are installed:

- [Golang (v1.21 or later)](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL (Running on a local or remote instance)](https://www.postgresql.org/download/)
- A modern web browser (Chrome, Firefox, Safari)

### Database Setup

To run a PostgreSQL container locally (for testing purposes), you can use the following command:

```bash
docker run --name my_postgres \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  -v /path/to/local/db/data:/var/lib/postgresql/data \
  -d postgres
```

Ensure that the `POSTGRES_USER`, `POSTGRES_PASSWORD`, and `POSTGRES_DB` are properly updated in the application code to match these values.

If an external team is using their own PostgreSQL database, they can update the database configurations in the `configs/config.yaml` file. This file contains details such as the database host, port, user, and password. Make sure to update the following fields accordingly:

```yaml
database:
  host: "<new_postgres_host>"
  port: <new_postgres_port>
  user: "<new_postgres_user>"
  password: "<new_postgres_password>"
  dbname: "<new_postgres_dbname>"
```

Additionally, ensure that the `internal/db/postgres.go` file correctly reads these configurations to establish a connection to the new PostgreSQL server.

## Running the Application

You can run the application locally in two ways: **directly using Go** or **using Docker**.

### Running Locally (Without Docker)

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/infra-env-dashboard.git
   cd infra-env-dashboard
   ```

2. **Install Dependencies**:

   Ensure you have Go modules enabled.

   ```bash
   go mod tidy
   ```

3. **Run the Application**:

   ```bash
   go run cmd/server/main.go
   ```

4. **Access the Dashboard**:

   Open your web browser and navigate to: [http://localhost:8080](http://localhost:8080)

### Running with Docker

This project includes a **multistage Dockerfile** for easier building and deployment.

1. **Build the Docker Image**:

   ```bash
   docker build -t infra-env-dashboard .
   ```

2. **Run the Docker Container**:

   ```bash
   docker run -p 8080:8080 infra-env-dashboard
   ```

3. **Access the Dashboard**:

   Open your web browser and navigate to: [http://localhost:8080](http://localhost:8080)

## Directory Structure

Below is an overview of the project's directory structure:

```
infra_env_dashboard/
├── cmd/
│   ├── server/
│   │   └── main.go                # Entry point for running the service
├── configs/
│   └── config.yaml                # Configuration files (e.g., database, ports, etc.)
├── internal/
│   ├── db/
│   │   ├── migrations/            # Database migration files
│   │   ├── postgres.go            # PostgreSQL connection and helper methods
│   ├── environments/
│   │   ├── handler.go             # Handlers for environment-related endpoints
│   │   ├── service.go             # Business logic for environment features
│   │   └── repository.go          # Database queries for environments
│   ├── middlewares/
│   │   └── auth.go                # Middleware for authentication and authorization
│   ├── common/
│   │   ├── utils.go               # Utility functions used across the application
│   │   ├── responses.go           # Common response helpers (e.g., JSON serialization)
│   └── models/
│       └── environment.go         # Models for environment objects
├── pkg/
│   ├── database/
│   │   ├── db.go                  # Database connection pool and configuration
│   └── logger/
│       └── logger.go              # Logging implementation for reuse across services
├── static/
│   ├── css/
│   │   └── style.css              # CSS files
│   └── js/
│       └── script.js              # JavaScript files
├── templates/
│   ├── layout.html                # HTML layout for the entire UI
│   └── dashboard.html             # Dashboard HTML template
├── test/
│   ├── environments_test.go       # Unit tests for environment features
│   └── integration_test.go        # Integration tests for full system testing
├── go.mod                         # Go module file
└── go.sum                         # Go dependencies
```

## Environment Variables

To configure the application, the following **environment variables** are used:

- **`PORT`**: Defines the port on which the server runs (default is `8080`).
- **Database Variables**: Update the credentials in the main application code or via environment variables for the database.

## Docker Deployment

### Docker Multi-Stage Build

This project utilizes a **multi-stage Docker build** to keep the final image lightweight.

1. **Build Stage**: Uses `golang:1.21` to build the Go application and compile the binary.
2. **Run Stage**: Uses `alpine:latest` to run the application, minimizing the image size.

**Commands Used**:

- **Build**: `docker build -t infra-env-dashboard .`
- **Run**: `docker run -p 8080:8080 infra-env-dashboard`

## Using Themes

The dashboard provides a **light theme** and **dark theme** toggle. To change themes:

1. Click on the **gear icon** in the top-right corner of the header.
2. Select either **Light Theme** or **Dark Theme** from the dropdown menu.

The theme toggle will immediately update the dashboard's appearance.

## Known Issues & Future Improvements

- **Database Connection**: Ensure that the database connection credentials are set correctly.
- **Docker Networking**: When running PostgreSQL in a separate container, ensure Docker networking is set up so that the Go application can communicate with the database.

### Future Improvements

- **User Authentication**: Add user roles and authentication for dashboard access.
- **Custom Widgets**: Allow users to add or remove widgets on the dashboard based on their needs.
- **Notifications**: Integrate real-time notifications for deployment or environment alerts.
- **Metrics**: Add more detailed and customizable metrics with historical views.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

### Steps to Contribute:

1. **Fork the Repository**.
2. **Create a Feature Branch**: `git checkout -b feature/new-feature`
3. **Commit Changes**: `git commit -m 'Add new feature'`
4. **Push to Branch**: `git push origin feature/new-feature`
5. **Open a Pull Request**.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For more information, please reach out to **[learntogrowmore@gmail.com](mailto\:learntogrowmore@gmail.com)**.

---

Thank you for using the **Infrastructure & Environment Dashboard**. We hope this tool helps you efficiently monitor and manage your infrastructure. If you have any suggestions or issues, feel free to open an issue on our GitHub repository!