package cmd

import (
	"fmt"

	"bitbucket.org/uthark/yttrium/internal/version"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %+v\n", version.GetVersionInfo())

		},
	}

	rootCmd.AddCommand(versionCmd)
}
