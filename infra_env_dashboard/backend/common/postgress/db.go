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
		// If DATABASE_URL is available, use it to connect to PostgreSQL
		return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	}

	// If DATABASE_URL is not set, fallback to reading from config.yaml
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./common/configs")

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

	// Open a GORM database connection
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	log.Println("Database connected successfully.")
	return db, nil
}
