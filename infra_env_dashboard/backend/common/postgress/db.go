package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDB() {
	viper.SetConfigFile("../../common/configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dbConfig := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.sslmode"),
	)

	var err error
	DB, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}
}
