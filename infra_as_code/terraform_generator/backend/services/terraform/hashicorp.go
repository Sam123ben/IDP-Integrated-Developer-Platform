// backend/services/terraform/hashicorp.go

package terraform

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProviderTemplate struct {
	Source  string `json:"source"`
	Version string `json:"version"`
}

// FetchProviderTemplate fetches provider template from HashiCorp Terraform Registry
func FetchProviderTemplate(providerName, version string) (ProviderTemplate, error) {
	url := fmt.Sprintf("https://registry.terraform.io/v1/providers/hashicorp/%s/%s/download", providerName, version)
	resp, err := http.Get(url)
	if err != nil {
		return ProviderTemplate{}, fmt.Errorf("failed to fetch provider template: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ProviderTemplate{}, fmt.Errorf("unexpected status code %d while fetching provider template", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProviderTemplate{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var template ProviderTemplate
	if err := json.Unmarshal(body, &template); err != nil {
		return ProviderTemplate{}, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return template, nil
}
