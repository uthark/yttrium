package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yttrium",
	Short: "Microservice for...",
	Long:  `Scaffold for microservice.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World.")
	},
}

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}
