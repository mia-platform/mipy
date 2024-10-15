package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/mia-platform/mipy/internal/cliconfig"
	"github.com/spf13/cobra"
)

var crList []string
var parallel bool
var errorCode int
var debug bool
var dryRun bool
var environment string

type CRInfo struct {
	Path         string
	TemplateType string
}

func handleTerraformCR(cr CRInfo) error {
	fmt.Println("Handle template %s", cr.Path)
	configFilePath := filepath.Join(cr.Path, "config.tf")
	variablesFilePath := filepath.Join(cr.Path, "variables.env")

	configContent, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading config.tf: %v", err)
	}

	// Read the variables.env file
	variablesContent, err := os.ReadFile(variablesFilePath)
	if err != nil {
		return fmt.Errorf("error reading variables.env: %v", err)
	}
	fmt.Println(configContent, variablesContent)
	return nil
}

func launchCR(cr CRInfo) error {
	if cr.TemplateType == "terraform" {
		handleTerraformCR(cr)
	} else {
		fmt.Println("Template type not implemented")
		return fmt.Errorf("Template type not implemented")
	}

	return nil
}

func getCRInfos(basePath string, template cliconfig.Template, environment string) ([]CRInfo, error) {
	var crInfos []CRInfo
	searchPath := filepath.Join(basePath, template.Id, "environment", environment)

	err := filepath.WalkDir(searchPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != searchPath {
			crInfo := CRInfo{
				Path:         path,
				TemplateType: template.Type,
			}
			crInfos = append(crInfos, crInfo)
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
			fmt.Println(crInfos)

			if dryRun == true {
				return nil
			}

			launchCR(crInfos[1])

			// get all cr for each template id in config
			return errors.New("command not implemented")
		},
	}

	cmd.Flags().StringSliceVar(&crList, "cr-list", []string{}, "List of CRs to launch")
	cmd.Flags().BoolVar(&parallel, "parallel", false, "Launch CRs in parallel")
	cmd.Flags().IntVar(&errorCode, "error-code", 500, "Error code to trigger on failure")
	cmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview CRs without executing")
	cmd.Flags().StringVarP(&environment, "environment", "e", "", "Environment to deploy")
	cmd.MarkFlagRequired("environment")
	return cmd
}
