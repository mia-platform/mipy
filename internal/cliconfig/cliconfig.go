package cliconfig

import (
	"encoding/json"
	"os"
)

type Config struct {
	BasePath  string     `json:"basePath"`
	Templates []Template `json:"templates"`
	LogLevel  string     `json:"logLevel"`
}

type Template struct {
	Type                string `json:"type"`
	Id                  string `json:"id"`
	CICDProvider        string `json:"cicdProvider"`
	CICDBaseUrl         string `json:"cicdProviderBaseUrl"`
	AzureOrganization   string `json:"azureOrganization"`
	AzureProject        string `json:"azureProject"`
	TerraformPipelineID string `json:"terraformPipelineId"`
}

func loadPreferredConfigPath() (string, error) {
	data, err := os.ReadFile("config_path.txt")
	if err != nil {
		return "mipyconfig.json", nil
	}
	return string(data), nil
}

func SavePreferredConfigPath(path string) error {
	return os.WriteFile("config_path.txt", []byte(path), 0644)
}

func ReadConfigFile() (*Config, error) {
	path, err := loadPreferredConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := new(Config)
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config, nil
}
