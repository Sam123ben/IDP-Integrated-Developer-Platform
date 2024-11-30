// backend/models/config.go

package models

type Config struct {
	TerraformVersion string              `json:"terraform_version"`
	Providers        []Provider          `json:"providers"`
	Backend          Backend             `json:"backend"`
	Modules          []Module            `json:"modules"`
	Variables        map[string]Variable `json:"variables"`
	Region           string              `json:"region"`
	Environment      string              `json:"environment"`
}

type Provider struct {
	Name          string            `json:"name"`
	Source        string            `json:"source"`
	Version       string            `json:"version"`
	AuthVariables map[string]string `json:"auth_variables"`
}

type Backend struct {
	Type               string            `json:"type"`
	Parameters         map[string]string `json:"parameters"`
	ResourceGroupName  string            `json:"resource_group_name"`
	StorageAccountName string            `json:"storage_account_name"`
	ContainerName      string            `json:"container_name"`
	Key                string            `json:"key"`
	SubscriptionId     string            `json:"subscription_id"`
	TenantID           string            `json:"tenant_id"`
	ClientID           string            `json:"client_id"`
	AccessKey          string            `json:"access_key"`
}

type Module struct {
	ModuleName string                    `json:"module_name"`
	Source     string                    `json:"source"`
	Variables  map[string]ModuleVariable `json:"variables"`
	Outputs    map[string]ModuleOutput   `json:"outputs,omitempty"`
	DependsOn  []string                  `json:"depends_on,omitempty"`
}

// Embed Variable within ModuleVariable
type ModuleVariable struct {
	Variable
	// Add any module-specific fields here if necessary
}

type ModuleOutput struct {
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
}

type Variable struct {
	Type        string                 `json:"type"`
	Description string                 `json:"description"`
	Default     interface{}            `json:"default,omitempty"`
	Sensitive   bool                   `json:"sensitive,omitempty"`
	Value       interface{}            `json:"value,omitempty"`
	Attributes  map[string]interface{} `json:"attributes,omitempty"` // Add attributes for object/tuple types
	Validation  *Validation            `json:"validation,omitempty"`
}

type Validation struct {
	Condition    string `json:"condition"`
	ErrorMessage string `json:"error_message"`
}
