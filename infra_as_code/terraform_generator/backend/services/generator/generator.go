// backend/services/generator/generator.go

package generator

import (
	"backend/services/config"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateProviderFile(baseOutputDir string, providerConfig config.ProviderConfig) error {
	providerFile := filepath.Join(baseOutputDir, "provider.tf")
	content := `
terraform {
  required_providers {
    {{ .Name }} = {
      source  = "hashicorp/{{ .Name }}"
      version = "{{ .Version }}"
    }
  }
}

provider "{{ .Name }}" {
  # Configuration options
}
`
	return renderToFile(providerFile, content, providerConfig)
}

func GenerateBackendFile(baseOutputDir string, backendConfig config.BackendConfig) error {
	backendFile := filepath.Join(baseOutputDir, "backend.tf")
	content := `
terraform {
  backend "azurerm" {
    resource_group_name  = "{{ .ResourceGroupName }}"
    storage_account_name = "{{ .StorageAccountName }}"
    container_name       = "{{ .ContainerName }}"
    key                  = "{{ .Key }}"
    {{- if .UseOIDC }}
    use_oidc             = true
    client_id            = "{{ .ClientID }}"
    subscription_id      = "{{ .SubscriptionID }}"
    tenant_id            = "{{ .TenantID }}"
    {{- end }}
  }
}
`
	return renderToFile(backendFile, content, backendConfig)
}

func renderToFile(filePath, content string, data interface{}) error {
	tmpl, err := template.New("file").Parse(content)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", filePath, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to render template to file %s: %v", filePath, err)
	}

	return nil
}
