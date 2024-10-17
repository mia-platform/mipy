package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mia-platform/mipy/internal/cliconfig"
	"github.com/spf13/cobra"
)

func ConfigGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get current config",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println("Get current configuration:")
			config, err := cliconfig.ReadConfigFile()
			jsonToPrint, err := json.MarshalIndent(config, "", "  ")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(string(jsonToPrint))
		},
	}
	return cmd
}

func ConfigSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set [PATH]",
		Short: "set config given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			fmt.Printf("Setting configuration from file: %s\n", path)
			if err := cliconfig.SavePreferredConfigPath(path); err != nil {
				return errors.New("Error while retrieving new configuration")
			}
			return nil
		},
	}

	return cmd
}

func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Set or get config",
	}

	cmd.AddCommand(
		ConfigGetCmd(),
		ConfigSetCmd(),
	)

	return cmd
}
