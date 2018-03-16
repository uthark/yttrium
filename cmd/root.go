package cmd

import (
	"os"

	"github.com/uthark/yttrium/internal/rest"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "yttrium",
	Short:        "Microservice for...",
	Long:         `Scaffold for microservice.`,
	RunE:         startServer,
	SilenceUsage: true,
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

var server = rest.NewServer()

func startServer(_ *cobra.Command, _ []string) error {
	stop := make(chan int)
	server.Init(stop)

	server.Start()

	<-stop

	return nil
}

func restartServer() {
	server.Restart()
}
