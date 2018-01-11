package cmd

import (
	"log"

	"os"

	"bitbucket.org/uthark/yttrium/internal/rest"
	"github.com/spf13/cobra"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

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
		logger.Fatal(err)
	}

}

func startServer(cmd *cobra.Command, args []string) error {
	server := rest.NewServer()
	return server.Start()
}
