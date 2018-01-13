package cmd

import (
	"strings"

	"bitbucket.org/uthark/yttrium/internal/config"
	"bitbucket.org/uthark/yttrium/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file to use.")
	rootCmd.PersistentFlags().Uint16("http-port", 8080, "HTTP Port to use")

	util.Must(viper.BindPFlags(rootCmd.PersistentFlags()))
}

func initConfig() {
	viper.SetEnvPrefix("YTTRIUM")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	// read configuration.
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
	logger.Println("Current configuration:", viper.AllSettings())
	c := &config.Configuration{}
	util.Must(viper.Unmarshal(c))
	config.SetDefaultConfiguration(c)

}
