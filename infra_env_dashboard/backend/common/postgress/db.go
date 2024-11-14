package postgress

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the PostgreSQL database connection using GORM
func InitDB() (*gorm.DB, error) {
	// Check if DATABASE_URL environment variable is set
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		log.Println("Attempting to connect to PostgreSQL using DATABASE_URL environment variable.")
		// Try to connect using DATABASE_URL
		db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to PostgreSQL using DATABASE_URL.")
			return db, nil
		}
		// Log the error and fall back to config.yaml if connection fails
		log.Printf("Failed to connect using DATABASE_URL: %v. Falling back to config.yaml.", err)
	}

	// If DATABASE_URL is not set or fails, fall back to config.yaml
	log.Println("Falling back to config.yaml for database connection settings.")

	// Configure viper to find the config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./common/configs") // Specify the relative path to config.yaml

	// Read in the config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Construct the DB connection string from config.yaml
	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.sslmode"),
	)

	log.Println("Attempting to connect to PostgreSQL using settings from config.yaml.")

	// Open a GORM database connection
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL using config.yaml.")
	return db, nil
}