package config

// Configuration is an app configuration
type Configuration struct {
	// HTTPPort is a port used for listening incoming HTTP traffic
	HTTPPort uint16 `mapstructure:"http-port"`
}

var defaultConfiguration = &Configuration{}

// SetDefaultConfiguration sets default configuration.
func SetDefaultConfiguration(c *Configuration) {
	defaultConfiguration = c
}

// DefaultConfiguration returns current configuration.
func DefaultConfiguration() *Configuration {
	return defaultConfiguration
}