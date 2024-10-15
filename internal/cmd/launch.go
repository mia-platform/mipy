package cmd

import (
	"errors"
	"fmt"
	"io/fs"
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

func getCRsDirectories(basePath string, templateId string, environment string) ([]string, error) {
	var crPaths []string
	searchPath := filepath.Join(basePath, templateId, "environment", environment)

	err := filepath.WalkDir(searchPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && path != searchPath {
			crPaths = append(crPaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return crPaths, nil
}

func getCRsFullPath(config cliconfig.Config, environment string) ([]string, error) {
	var crs []string
	var err error
	for _, template := range config.Templates {

		crDirs, err := getCRsDirectories(config.BasePath, template.Id, environment)
		if err != nil {
			fmt.Println(err)
			crDirs = []string{}
		}
		crs = append(crs, crDirs...)
	}
	return crs, err
}

func LaunchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch",
		Short: "Launch",
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Println(environment)

			// get configuration
			config, err := cliconfig.ReadConfigFile()
			if err != nil {
				fmt.Errorf("Failed get templates from configuration")
			}
			crs, err := getCRsFullPath(*config, environment)
			fmt.Println(crs)

			if dryRun == true {
				return nil
			}

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
