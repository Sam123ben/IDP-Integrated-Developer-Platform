package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    _ "github.com/lib/pq"
)

var DB *sql.DB

// Config holds database connection parameters
type Config struct {
    User     string
    Password string
    DBName   string
    Host     string
    Port     int
}

// Connect opens a connection to the PostgreSQL database using configuration parameters
func Connect(cfg Config) error {
    dsn := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
    )

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        return fmt.Errorf("error opening database: %w", err)
    }

    if err = DB.Ping(); err != nil {
        return fmt.Errorf("could not connect to the database: %w", err)
    }

    log.Println("Connected to the database successfully")
    return nil
}

// RunMigrations runs all .up.sql migration scripts in the specified migrations folder
func RunMigrations(migrationsPath string) error {
    files, err := os.ReadDir(migrationsPath)
    if err != nil {
        return fmt.Errorf("unable to read migrations directory %s: %w", migrationsPath, err)
    }

    for _, file := range files {
        if file.IsDir() || !strings.HasSuffix(file.Name(), ".up.sql") {
            continue
        }

        path := filepath.Join(migrationsPath, file.Name())
        script, err := os.ReadFile(path)
        if err != nil {
            return fmt.Errorf("could not read migration file %s: %w", file.Name(), err)
        }

        log.Printf("Running migration: %s", file.Name())
        if err := executeMigration(script); err != nil {
            return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
        }

        log.Printf("Successfully ran migration: %s", file.Name())
    }

    return nil
}

// executeMigration runs a single migration script within a transaction
func executeMigration(script []byte) error {
    tx, err := DB.Begin()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %w", err)
    }

    if _, err := tx.Exec(string(script)); err != nil {
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("failed to execute migration and rollback: %w; rollback error: %v", err, rbErr)
        }
        return fmt.Errorf("failed to execute migration: %w", err)
    }

    if err := tx.Commit(); err != nil {
        return fmt.Errorf("failed to commit migration: %w", err)
    }

    return nil
}