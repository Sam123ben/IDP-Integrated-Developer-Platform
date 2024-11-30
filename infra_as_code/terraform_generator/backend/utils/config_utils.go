// backend/utils/config_utils.go

package utils

import (
	"backend/models"
	"encoding/json"
	"os"
)

// LoadConfig reads the configuration from a JSON file
func LoadConfig(path string) (*models.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config models.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
