package database

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    _ "github.com/lib/pq"
)

var DB *sql.DB

// Connect opens a connection to the PostgreSQL database
func Connect(user, password, dbname, host string, port int) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    log.Println("Connected to the database successfully")
}

// RunMigrations runs migration scripts from the migrations folder
func RunMigrations(migrationsPath string) {
    files, err := ioutil.ReadDir(migrationsPath)
    if err != nil {
        log.Fatalf("Unable to read migrations directory: %v", err)
    }

    for _, file := range files {
        if file.IsDir() || !strings.HasSuffix(file.Name(), ".up.sql") {
            continue
        }

        path := fmt.Sprintf("%s/%s", migrationsPath, file.Name())
        script, err := ioutil.ReadFile(path)
        if err != nil {
            log.Fatalf("Could not read migration file %s: %v", file.Name(), err)
        }

        if _, err := DB.Exec(string(script)); err != nil {
            log.Fatalf("Failed to execute migration %s: %v", file.Name(), err)
        }

        log.Printf("Successfully ran migration: %s", file.Name())
    }
}