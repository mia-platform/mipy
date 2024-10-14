package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "View CR list to be launched",
		Run: func(cmd *cobra.Command, args []string) {
			// Logica per visualizzare in anteprima le CRs
			fmt.Println("Previewing the CRs...")
			fmt.Errorf("Command not implemented")
		},
	}

	return cmd
}
