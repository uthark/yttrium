package cmd

import (
	"os"

	"bitbucket.org/uthark/yttrium/internal/rest"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yttrium",
	Short: "Microservice for...",
	Long:  `Scaffold for microservice.`,
	RunE:  startServer,
}

// Execute executes the root command.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.SetOutput(os.Stderr)
		logger.Fatal(err)
		os.Exit(-1)
	}
	os.Exit(0)

}

func startServer(_ *cobra.Command, _ []string) error {
	server := rest.NewServer()
	return server.Start()
}
