package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.yttrium.yaml)")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Search config in current directory with name ".yttrium" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigName(".yttrium")
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.Println("No config file found: ", err)
	}
}
