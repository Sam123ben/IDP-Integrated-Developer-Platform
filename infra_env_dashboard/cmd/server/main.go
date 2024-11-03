package main

import (
    "infra_env_dashboard/pkg/database"
    "log"
    "net/http"
    "html/template"
    _ "github.com/lib/pq"
)

var templates *template.Template

func main() {
    // Load templates
    templates = template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/dashboard.html",
    ))

    // Connect to PostgreSQL
    dbUser := "myuser"
    dbPassword := "mypassword"
    dbName := "mydatabase"
    dbHost := "localhost"
    dbPort := 5432

    database.Connect(dbUser, dbPassword, dbName, dbHost, dbPort)

    // Run database migrations
    migrationsPath := "internal/db/migrations"
    database.RunMigrations(migrationsPath)

    // Static file server for CSS, JS
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Route for dashboard
    http.HandleFunc("/", dashboardHandler)

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

// fetchEnvironments queries the database for environment information
func fetchEnvironments() ([]Environment, error) {
    rows, err := database.DB.Query("SELECT name, description FROM environments")
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

// Environment struct
type Environment struct {
    Name        string
    Description string
}