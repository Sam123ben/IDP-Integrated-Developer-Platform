package database

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// Config holds database connection parameters
type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// Connect initializes the database connection and assigns it to the global DB variable
func Connect(cfg Config) error {
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
    )

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        return fmt.Errorf("failed to open database connection: %w", err)
    }

    // Set connection pool parameters
    DB.SetMaxOpenConns(25)
    DB.SetMaxIdleConns(25)
    DB.SetConnMaxLifetime(5 * time.Minute)

    // Verify the database connection
    if err = DB.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %w", err)
    }

    log.Println("Database connection established successfully")
    return nil
}

// Close safely closes the database connection
func Close() error {
    if DB != nil {
        return DB.Close()
    }
    return nil
}