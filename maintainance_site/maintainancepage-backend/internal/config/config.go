// backend/internal/config/config.go
package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost         string
    DBPort         int
    DBUser         string
    DBPassword     string
    DBName         string
    LoggingEnabled bool
    ServerBaseURL  string
    Port           string
}

func LoadConfig() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        return nil, fmt.Errorf("error loading .env file: %w", err)
    }

    dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
    if err != nil {
        return nil, fmt.Errorf("invalid DB_PORT: %w", err)
    }

    loggingEnabled, _ := strconv.ParseBool(os.Getenv("LOGGING_ENABLED"))

    return &Config{
        DBHost:         os.Getenv("DB_HOST"),
        DBPort:         dbPort,
        DBUser:         os.Getenv("DB_USER"),
        DBPassword:     os.Getenv("DB_PASSWORD"),
        DBName:         os.Getenv("DB_NAME"),
        LoggingEnabled: loggingEnabled,
        ServerBaseURL:  os.Getenv("SERVER_BASE_URL"),
        Port:           os.Getenv("PORT"),
    }, nil
}