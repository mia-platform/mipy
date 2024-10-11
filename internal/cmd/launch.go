package cmd

import "github.com/spf13/cobra"

func LaunchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch",
		Short: "Launch",
	}

	return cmd
}
