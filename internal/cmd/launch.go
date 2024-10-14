package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var crList []string
var parallel bool
var errorCode int
var debug bool
var dryRun bool

func LaunchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "launch",
		Short: "Launch",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return errors.New("command not implemented")
		},
	}

	cmd.Flags().StringSliceVar(&crList, "cr-list", []string{}, "List of CRs to launch")
	cmd.Flags().BoolVar(&parallel, "parallel", false, "Launch CRs in parallel")
	cmd.Flags().IntVar(&errorCode, "error-code", 500, "Error code to trigger on failure")
	cmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Preview CRs without executing")
	return cmd
}
