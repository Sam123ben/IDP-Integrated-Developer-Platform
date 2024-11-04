package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
	} `mapstructure:"database"`
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file
	viper.AddConfigPath(path)     // Path to look for the config file

	// Setting default values
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.dbname", "infra_env_dashboard")

	// Enable reading from environment variables
	viper.AutomaticEnv()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Validate required fields (optional, but useful for ensuring critical fields are set)
	if config.Database.Host == "" || config.Database.User == "" || config.Database.DBName == "" {
		return nil, fmt.Errorf("missing required database configuration fields")
	}

	return &config, nil
}
