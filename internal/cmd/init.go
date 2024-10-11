package cmd

import "github.com/spf13/cobra"

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "View CR list to be launched",
	}

	return cmd
}
