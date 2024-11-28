// backend/models/config.go

package models

type Config struct {
	TerraformVersion string              `json:"terraform_version"`
	Providers        []Provider          `json:"providers"`
	Backend          Backend             `json:"backend"`
	Modules          []Module            `json:"modules"`
	Variables        map[string]Variable `json:"variables"`
	Region           string              `json:"region"`      // Add this field
	Environment      string              `json:"environment"` // Add this field
}

type Provider struct {
	Name          string            `json:"name"`
	Source        string            `json:"source"`
	Version       string            `json:"version"`
	AuthVariables map[string]string `json:"auth_variables"`
}

type Backend struct {
	ResourceGroupName  string `json:"resource_group_name"`
	StorageAccountName string `json:"storage_account_name"`
	ContainerName      string `json:"container_name"`
	Key                string `json:"key"`
	SubscriptionId     string `json:"subscription_id"`
	TenantId           string `json:"tenant_id"`
	ClientId           string `json:"client_id"`
	AccessKey          string `json:"AccessKey"`
}

type Module struct {
	ModuleName string            `json:"module_name"`
	Source     string            `json:"source"`
	Variables  map[string]string `json:"variables"`
}

type Variable struct {
	Value       string `json:"value"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Sensitive   bool   `json:"sensitive"`
}
