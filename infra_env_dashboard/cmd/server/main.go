package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"infra_env_dashboard/configs"
	"infra_env_dashboard/pkg/database"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var templates *template.Template

func main() {
	if err := initializeServer(); err != nil {
		log.Fatalf("Server initialization failed: %s", err)
	}
}

// initializeServer initializes configurations, database, templates, and starts the server
func initializeServer() error {
	cfg, err := configs.LoadConfig("./configs")
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}

	if err := setupDatabase(cfg); err != nil {
		return fmt.Errorf("database setup failed: %w", err)
	}
	log.Println("Database connected successfully")

	if err := loadTemplates(); err != nil {
		return fmt.Errorf("failed to load templates: %w", err)
	}

	migrationsPath := "internal/db/migrations"
	if err := database.RunMigrations(migrationsPath); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	setupRoutes()

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	return nil
}

// setupDatabase establishes a connection to the database
func setupDatabase(cfg *configs.Config) error {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
	var err error
	database.DB, err = sql.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	return database.DB.Ping()
}

// loadTemplates loads HTML templates
func loadTemplates() error {
	var err error
	templates, err = template.ParseFiles("templates/layout.html", "templates/dashboard.html")
	return err
}

// setupRoutes defines HTTP routes
func setupRoutes() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/api/latest-data", getLatestDataHandler)
}

// dashboardHandler renders the main dashboard page
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	environments, err := fetchEnvironments()
	if err != nil {
		log.Printf("Error fetching environment data: %s", err)
		http.Error(w, "Error fetching environment data", http.StatusInternalServerError)
		return
	}

	companyName, err := fetchCompanyName()
	if err != nil {
		log.Printf("Error fetching company name: %s", err)
		http.Error(w, "Error fetching company name", http.StatusInternalServerError)
		return
	}

	data := struct {
		CompanyName  string
		Environments []Environment
	}{
		CompanyName:  companyName,
		Environments: environments,
	}

	if err := templates.ExecuteTemplate(w, "layout.html", data); err != nil {
		log.Printf("Template execution error: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}
}

// getLatestDataHandler handles the API request to fetch the latest data
func getLatestDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	environments, err := fetchEnvironments()
	if err != nil {
		log.Printf("Error fetching latest data: %s", err)
		http.Error(w, "Error fetching latest data", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(environments); err != nil {
		log.Printf("Error encoding response data: %s", err)
		http.Error(w, "Error encoding response data", http.StatusInternalServerError)
	}
}

// fetchEnvironments queries the database for environment information
// fetchEnvironments queries the database for environment information
func fetchEnvironments() ([]Environment, error) {
	query := `
        SELECT 
            e.environment_name, 
            e.description, 
            e.environment_type,
            g.name AS group_name, 
            cg.name AS customer_name, 
            e.url, 
            e.status, 
            e.contact, 
            e.app_version, 
            e.db_version, 
            e.comments 
        FROM 
            environments e
        LEFT JOIN 
            groups g ON e.group_id = g.id
        LEFT JOIN 
            groups cg ON e.customer_id = cg.id
        ORDER BY 
            e.updated_at DESC 
        LIMIT 10;
    `

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var environments []Environment
	for rows.Next() {
		var env Environment
		if err := rows.Scan(
			&env.EnvironmentName,
			&env.Description,
			&env.EnvironmentType,
			&env.GroupName,
			&env.CustomerName,
			&env.URL,
			&env.Status,
			&env.Contact,
			&env.AppVersion,
			&env.DBVersion,
			&env.Comments,
		); err != nil {
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

// Environment represents an environment record
type Environment struct {
	EnvironmentName string `json:"environment_name"`
	Description     string `json:"description"`
	EnvironmentType string `json:"environment_type"`
	GroupName       string `json:"group_name,omitempty"`
	CustomerName    string `json:"customer_name,omitempty"`
	URL             string `json:"url,omitempty"`
	Status          string `json:"status,omitempty"`
	Contact         string `json:"contact,omitempty"`
	AppVersion      string `json:"app_version,omitempty"`
	DBVersion       string `json:"db_version,omitempty"`
	Comments        string `json:"comments,omitempty"`
}
