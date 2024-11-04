package common

import (
	"context"
	"fmt"
	"log"

	"io/ioutil"

	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v2"
)

var DB *pgx.Conn

// Config struct to hold database configuration
type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"database"`
}

// LoadConfig loads configuration from config.yaml
func LoadConfig(configFile string) (*Config, error) {
	config := &Config{}

	// Read YAML file
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Unmarshal YAML file into Config struct
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config file: %v", err)
	}

	return config, nil
}

// Connect initializes the database connection using config.yaml
func Connect() {
	config, err := LoadConfig("backend/common/config.yaml") // Path to your config.yaml
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
	)

	DB, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}
