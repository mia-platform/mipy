package cmd

import "github.com/spf13/cobra"

func ConfigGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get current config",
		Run: func(cmd *cobra.Command, _ []string) {

		},
	}
	return cmd
}

func ConfigSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "set config given path",
		RunE: func(cmd *cobra.Command, _ []string) error {
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
	)

	return cmd
}
