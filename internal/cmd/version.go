package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of mipy",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("TO IMPLEMENT mipy vX.Y.Z")
		},
	}

	return versionCmd
}
