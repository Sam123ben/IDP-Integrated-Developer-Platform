// backend/services/config/loader.go

package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Providers map[string]ProviderConfig `yaml:"providers"`
}

type ProviderConfig struct {
	Name          string              `yaml:"name"`
	Version       string              `yaml:"version"`
	Backend       BackendConfig       `yaml:"backend"`
	ResourceGroup ResourceGroupConfig `yaml:"resource_group"`
}

type BackendConfig struct {
	ResourceGroupName  string `yaml:"resource_group_name"`
	StorageAccountName string `yaml:"storage_account_name"`
	ContainerName      string `yaml:"container_name"`
	Key                string `yaml:"key"`
	UseOIDC            bool   `yaml:"use_oidc"`
	ClientID           string `yaml:"client_id"`
	SubscriptionID     string `yaml:"subscription_id"`
	TenantID           string `yaml:"tenant_id"`
}

type ResourceGroupConfig struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
}

func LoadConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("failed to decode config file: %v", err)
	}

	return config, nil
}
