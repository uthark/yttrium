package cmd

import (
	"strings"

	"github.com/uthark/yttrium/internal/config"
	"github.com/uthark/yttrium/internal/util"
	"github.com/fsnotify/fsnotify"
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
	logger.Println("Config file used: ", viper.ConfigFileUsed())
	logger.Println("Current configuration:", viper.AllSettings())

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Println("Config file changed:", e.Name, e.Op, e)
		c := &config.Configuration{}
		err := viper.Unmarshal(c)
		if err != nil {
			logger.Println("Unable to unmarshal config:", err)
		}
		config.SetDefaultConfiguration(c)
		logger.Println(c)
		restartServer()
	})

	c := &config.Configuration{}
	util.Must(viper.Unmarshal(c))
	config.SetDefaultConfiguration(c)

}
