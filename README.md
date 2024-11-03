# Infrastructure & Environment Dashboard

![Build Status](https://github.com/sam123ben/infra_env_dashboard/actions/workflows/ci-infra-dashboard.yml/badge.svg)

Welcome to the **Infrastructure & Environment Dashboard**! This lightweight, performant web-based dashboard allows users to **monitor, manage, and optimize their infrastructure**. The application follows a **microservice architecture** and adopts a 3-tier approach comprising a **UI**, **middleware**, and **database**. The system is fully customizable and driven entirely by its database schema, making it adaptable for various infrastructure needs.

## Project Overview

The **Infrastructure & Environment Dashboard** provides real-time insights into various environments such as **Development, Staging, and Production**. Instead of focusing on detailed server metrics, this dashboard provides a high-level view of **release and deployment statistics**, **version checks**, **URLs**, and the ability to **skip deployments** or **add comments**—enabling a real-time understanding of environment test cycles.

The UI and middleware are implemented using **Golang**, and **PostgreSQL** serves as the backend database. The modular nature of the application makes it highly customizable and easily extendable, enabling additional features such as tracking release statuses, adding or updating comments, managing deployment skips, and capturing detailed environment data.

### Key Features

- **Real-time Environment Overview**: View the status of multiple environments including deployment statuses and environment versions.
- **Full Customization Based on Database**: All content is dynamically generated from the database, ensuring easy customization by updating database entries.
- **Multi-theme Support**: Toggle between **light** and **dark** themes for user preference.
- **Lightweight & Scalable**: Built with **Golang**, offering a highly efficient, scalable architecture.
- **Microservice Architecture**: Modular structure for scalability and easy future enhancements.
- **Deployment Management**: Extend the dashboard to manage skips, comments, rollbacks, and more.

## Technology Stack

- **Frontend & Middleware**: [Golang](https://golang.org/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Containerization**: [Docker](https://www.docker.com/)
- **UI Templates**: HTML5, CSS3
- **Styling & UI Libraries**: Font Awesome for icons
- **Deployment**: Multistage Docker builds for compact, production-ready images.

## Project Intentions

The main goal of this dashboard is to offer a **centralized and customizable platform** for DevOps teams and infrastructure engineers to manage their environments efficiently. This project aims to be **lightweight**, **easy to customize**, and **highly scalable** to accommodate a range of infrastructure monitoring needs.

### Customizability

The dashboard content is entirely driven by the **database**. The structure and fields of the database determine the content shown on the dashboard. This means that customization, such as environment names, URLs, or deployment information, can be managed by simply updating the **PostgreSQL** database.

## Prerequisites

Before running the application, ensure the following prerequisites are met:

- [Golang (v1.21 or later)](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [PostgreSQL](https://www.postgresql.org/download/) (Local or remote instance)
- A modern web browser (Chrome, Firefox, Safari)

### Database Setup

Run a **PostgreSQL** container locally for testing purposes:

```bash
docker run --name my_postgres \
  -e POSTGRES_USER=myuser \
  -e POSTGRES_PASSWORD=mypassword \
  -e POSTGRES_DB=mydatabase \
  -p 5432:5432 \
  -v /path/to/local/db/data:/var/lib/postgresql/data \
  -d postgres
```

Ensure that the database credentials match the entries in the `configs/config.yaml` file for the app to connect properly:

```yaml
database:
  host: "<new_postgres_host>"
  port: <new_postgres_port>
  user: "<new_postgres_user>"
  password: "<new_postgres_password>"
  dbname: "<new_postgres_dbname>"
```

Ensure the `internal/db/postgres.go` file reads these configurations correctly to establish a connection to the new PostgreSQL server.

## Running the Application

You can run the application either **locally without Docker** or **using Docker**.

### Running Locally (Without Docker)

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/infra-env-dashboard.git
   cd infra-env-dashboard
   ```

2. **Install Dependencies**:

   Make sure Go modules are enabled.

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

This project includes a **multistage Dockerfile** to streamline building and deployment.

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

```
infra_env_dashboard/
├── cmd/
│   ├── server/
│   │   └── main.go                # Entry point for running the service
├── configs/
│   └── config.yaml                # Configuration files (e.g., database, ports, etc.)
│   ├── config.go                  # Go code to load config
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
│   ├── integration_test.go        # Integration tests for full system testing
│   ├── db_test.go                 # Unit tests for PostgreSQL database connection and functions
│   ├── handler_test.go            # Unit tests for HTTP handlers
│   ├── service_test.go            # Unit tests for business logic in services
│   └── utils_test.go              # Unit tests for common utility functions
├── go.mod                         # Go module file
└── go.sum                         # Go dependencies
```

## Using Themes

The dashboard features a **light theme** and **dark theme** for improved user experience. To change the theme:

1. Click the **gear icon** in the top-right corner of the header.
2. Select either **Light Theme** or **Dark Theme** from the dropdown menu.

The theme change will be immediately reflected in the UI.

## Known Issues & Future Improvements

- **Database Connection**: Double-check database credentials to avoid connection issues.
- **Docker Networking**: If running PostgreSQL in a separate container, ensure Docker networking is set up for communication.

### Future Improvements

- **User Authentication**: Add user roles and secure access for authorized users.
- **Custom Widgets**: Allow users to add/remove widgets based on their preferences.
- **Notifications**: Integrate real-time notifications for deployments and environment alerts.
- **Advanced Metrics**: Include more detailed, customizable metrics with historical views.

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

For more information, please reach out to **[learntogrowmore@gmail.com](mailto:learntogrowmore@gmail.com)**.

---

Thank you for using the **Infrastructure & Environment Dashboard**. We hope this tool helps you efficiently monitor and manage your infrastructure. If you have any suggestions or issues, feel free to open an issue on our GitHub repository!