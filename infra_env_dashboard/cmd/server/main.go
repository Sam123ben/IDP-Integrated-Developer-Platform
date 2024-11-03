package main

import (
    "database/sql"
    "fmt"
    "infra_env_dashboard/configs" // Importing the configs package correctly
    "infra_env_dashboard/pkg/database"
    "log"
    "net/http"
    "html/template"
    "encoding/json"
    _ "github.com/lib/pq" // PostgreSQL driver
)

var templates *template.Template

func main() {
    // Load configuration from configs/config.yaml
    cfg, err := configs.LoadConfig("./configs")
    if err != nil {
        log.Fatalf("Failed to load configuration: %s", err)
    }

    // Connect to PostgreSQL using config values
    dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

    database.DB, err = sql.Open("postgres", dbInfo)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %s", err)
    }

    // Verify the database connection
    if err = database.DB.Ping(); err != nil {
        log.Fatalf("Failed to ping the database: %s", err)
    }

    log.Println("Connected to the database successfully")

    // Load templates
    templates = template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/dashboard.html",
    ))

    // Run database migrations
    migrationsPath := "internal/db/migrations"
    database.RunMigrations(migrationsPath)

    // Static file server for CSS, JS
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Routes for dashboard and API
    http.HandleFunc("/", dashboardHandler)
    http.HandleFunc("/api/latest-data", getLatestDataHandler) // API endpoint for refreshing data

    // Start the server
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s", err.Error())
    }
}

// dashboardHandler renders the main dashboard page
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    // Fetch environment data from the database
    environments, err := fetchEnvironments()
    if err != nil {
        http.Error(w, "Error fetching environment data", http.StatusInternalServerError)
        return
    }

    // Fetch the company name from the database
    companyName, err := fetchCompanyName()
    if err != nil {
        http.Error(w, "Error fetching company name", http.StatusInternalServerError)
        return
    }

    // Passing data to the template
    data := struct {
        CompanyName  string
        Environments []Environment
    }{
        CompanyName:  companyName,
        Environments: environments,
    }

    if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// getLatestDataHandler handles the API request to fetch the latest data
func getLatestDataHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Fetch the latest environment data from the database
    environments, err := fetchEnvironments()
    if err != nil {
        http.Error(w, "Error fetching latest data", http.StatusInternalServerError)
        return
    }

    // Encode the data as JSON and send it as the response
    if err := json.NewEncoder(w).Encode(environments); err != nil {
        http.Error(w, "Error encoding response data", http.StatusInternalServerError)
    }
}

// fetchEnvironments queries the database for environment information
func fetchEnvironments() ([]Environment, error) {
    rows, err := database.DB.Query("SELECT name, description FROM environments ORDER BY updated_at DESC LIMIT 10")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var environments []Environment
    for rows.Next() {
        var env Environment
        if err := rows.Scan(&env.Name, &env.Description); err != nil {
            return nil, err
        }
        environments = append(environments, env)
    }

    return environments, nil
}

// fetchCompanyName queries the database for the company name
func fetchCompanyName() (string, error) {
    var companyName string
    row := database.DB.QueryRow("SELECT name FROM company LIMIT 1")
    if err := row.Scan(&companyName); err != nil {
        return "", err
    }

    return companyName, nil
}

// Environment struct to hold information about each environment
type Environment struct {
    Name        string
    Description string
}