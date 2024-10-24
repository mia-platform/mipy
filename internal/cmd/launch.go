package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mia-platform/mipy/internal/cliconfig"
	"github.com/spf13/cobra"
)

var crList []string
var parallel bool
var errorCode int
var debug bool
var dryRun bool
var environment string
var username string
var password string

type CRInfo struct {
	Path                string
	TemplateType        string
	CICDProvider        string
	CICDBaseUrl         string
	CICDOrganization    string
	CICDProject         string
	TerraformPipelineID string
	filesInsideFolder   bool // are cr file inside a dedicated folder {cr_name}/{filename}.ext or are multiple files {cr_name}.{filename}.ext
}

type TerraformRequestBody struct {
	Resources struct {
		Repositories struct {
			Self struct {
				RefName string `json:"refName"`
			} `json:"self"`
		} `json:"repositories"`
	} `json:"resources"`
	TemplateParameters struct {
		DebugMode            bool   `json:"DEBUG_MODE"`
		TerraformAutoApprove string `json:"TERRAFORM_AUTO_APPROVE"`
		TerraformAction      string `json:"TERRAFORM_ACTION"`
	} `json:"templateParameters"`
	Variables struct {
		AzureSubscriptionID struct {
			IsSecret bool   `json:"isSecret"`
			Value    string `json:"value"`
		} `json:"AZURE_SUBSCRIPTION_ID"`
		AzureTenantID struct {
			IsSecret bool   `json:"isSecret"`
			Value    string `json:"value"`
		} `json:"AZURE_TENANT_ID"`
		TerraformVariables struct {
			IsSecret bool   `json:"isSecret"`
			Value    string `json:"value"`
		} `json:"TERRAFORM_VARIABLES"`
	} `json:"variables"`
}

func azureTerraformCR(cr CRInfo, user string, password string, variablesContent string, envVars map[string]string) error {
	azSubscriptionID := envVars["AZURE_SUBSCRIPTION_ID"]
	azTenantID := envVars["AZURE_TENANT_ID"]
	terraformAction := envVars["ACTION"]
	terraformAutoApprove := envVars["AUTO_APPROVE"]
	terraformProjectId := envVars["TERRAFORM_PROJECT_ID"]

	requestBody := TerraformRequestBody{}
	requestBody.Resources.Repositories.Self.RefName = "refs/heads/{branch}" // TODO default 'main'
	requestBody.TemplateParameters.DebugMode = true
	requestBody.TemplateParameters.TerraformAutoApprove = terraformAutoApprove
	requestBody.TemplateParameters.TerraformAction = terraformAction
	requestBody.Variables.AzureSubscriptionID.IsSecret = false
	requestBody.Variables.AzureSubscriptionID.Value = azSubscriptionID
	requestBody.Variables.AzureTenantID.IsSecret = false
	requestBody.Variables.AzureTenantID.Value = azTenantID
	requestBody.Variables.TerraformVariables.IsSecret = false
	requestBody.Variables.TerraformVariables.Value = variablesContent

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshal error")
		return err
	}
	auth := base64.StdEncoding.EncodeToString([]byte(user + ":" + password))

	url := fmt.Sprintf("%s/%s/%s/_apis/pipelines/%s/runs?api-version=7.1", cr.CICDBaseUrl, cr.CICDOrganization, cr.CICDProject, terraformProjectId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+auth)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Print the status code
	fmt.Printf("Response Status Code: %d\n", resp.StatusCode)

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	fmt.Println("Response Body:", string(responseBody))
	return nil
}

func handleTerraformCR(cr CRInfo, user string, password string) error {
	fmt.Printf("Handle template %s\n", cr.Path)
	configFilePath := filepath.Join(cr.Path, "configs.env")
	variablesFilePath := filepath.Join(cr.Path, "variables.tf")

	variablesContent, err := os.ReadFile(variablesFilePath)
	if err != nil {
		return fmt.Errorf("error reading variables.tf: %v", err)
	}

	encodedvariablesContent := base64.StdEncoding.EncodeToString(variablesContent)

	// Read the configs.env file
	envVars, err := godotenv.Read(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading configs.env: %v", err)
	}

	if cr.CICDProvider != "azure" {
		return fmt.Errorf("CICD provider not supported")
	}
	return azureTerraformCR(cr, user, password, encodedvariablesContent, envVars)
}

func launchCR(cr CRInfo, user string, password string) error {
	if cr.TemplateType == "terraform" {
		err := handleTerraformCR(cr, user, password)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Template type %s not implemented", cr.TemplateType)
		return fmt.Errorf("Template type not implemented")
	}

	return nil
}

func getCRInfos(basePath string, template cliconfig.Template, environment string) ([]CRInfo, error) {
	var crInfos []CRInfo
	seenCRs := make(map[string]bool)
	searchPath := filepath.Join(basePath, template.Id, "environment", environment)

	baseDepth := len(strings.Split(searchPath, string(os.PathSeparator))) // Conta il livello di profondità del percorso base

	err := filepath.WalkDir(searchPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calcola il livello di profondità corrente
		currentDepth := len(strings.Split(path, string(os.PathSeparator)))

		// Limita la navigazione solo al primo livello
		if currentDepth > baseDepth+1 {
			return fs.SkipDir // Salta questa directory e tutto il suo contenuto
		}

		if d.IsDir() && path != searchPath {
			crInfo := CRInfo{
				Path:              path,
				TemplateType:      template.Type,
				CICDProvider:      template.CICDProvider,
				CICDBaseUrl:       template.CICDBaseUrl,
				filesInsideFolder: true,
			}

			if template.CICDProvider == "azure" {
				crInfo.CICDOrganization = template.AzureOrganization
				crInfo.CICDProject = template.AzureProject

			}
			crInfos = append(crInfos, crInfo)
		}

		if !d.IsDir() {
			filename := d.Name()
			crName := strings.Split(filename, ".")[0]
			crPath := filepath.Join(filepath.Dir(path), crName)
			if !seenCRs[crPath] {
				crInfo := CRInfo{
					Path:              crPath,
					TemplateType:      template.Type,
					CICDProvider:      template.CICDProvider,
					CICDBaseUrl:       template.CICDBaseUrl,
					filesInsideFolder: false,
				}

				if template.CICDProvider == "azure" {
					crInfo.CICDOrganization = template.AzureOrganization
					crInfo.CICDProject = template.AzureProject

				}
				crInfos = append(crInfos, crInfo)
				seenCRs[crPath] = true
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return crInfos, nil
}

func getCRsToLaunch(config cliconfig.Config, environment string) ([]CRInfo, error) {
	var crs []CRInfo
	var err error
	for _, template := range config.Templates {
		crInfos, err := getCRInfos(config.BasePath, template, environment)
		if err != nil {
			fmt.Println(err)
		}
		crs = append(crs, crInfos...)
	}
	return crs, err
}

func LaunchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch",
		Short: "Launch",
		RunE: func(cmd *cobra.Command, _ []string) error {
			// get configuration
			config, err := cliconfig.ReadConfigFile()
			if err != nil {
				fmt.Errorf("Failed get templates from configuration")
			}
			crInfos, err := getCRsToLaunch(*config, environment)
			for _, crInfo := range crInfos {
				fmt.Println(crInfo.Path)
			}

			if dryRun == true {
				return nil
			}

			for _, crInfo := range crInfos {
				err = launchCR(crInfo, username, password)
				if err != nil {
					fmt.Println(err)
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "Username for auth to cicd provider")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Password for auth to cicd provider")
	cmd.Flags().StringSliceVar(&crList, "cr-list", []string{}, "NOT IMPLEMENTED YET: List of CRs to launch")
	cmd.Flags().BoolVar(&parallel, "parallel", false, "NOT IMPLEMENTED YET: Launch CRs in parallel")
	cmd.Flags().IntVar(&errorCode, "error-code", 500, "Error code to trigger on failure")
	cmd.Flags().BoolVar(&debug, "debug", false, "NOT IMPLEMENTED YET: Enable debug mode")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview CRs without executing")
	cmd.Flags().StringVarP(&environment, "environment", "e", "", "Environment to deploy")
	cmd.MarkFlagRequired("environment")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	return cmd
}
