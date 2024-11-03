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

func LoadConfig(path string) (*Config, error) {
    viper.SetConfigName("config")  // Name of the config file (without extension)
    viper.SetConfigType("yaml")    // Type of the config file
    viper.AddConfigPath(path)      // Path to look for the config file

    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading config file: %w", err)
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("unable to decode into struct: %w", err)
    }

    return &config, nil
}