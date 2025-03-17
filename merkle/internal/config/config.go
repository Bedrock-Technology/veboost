package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// AppConfig represents the entire application configuration
type AppConfig struct {
	Logger   LogConfig         `mapstructure:"logger"`
	DocAuth  map[string]string `mapstructure:"doc_auth"`
	DNS      struct {
		Endpoint string `mapstructure:"endpoint"`
	} `mapstructure:"dns"`
}

// LogConfig represents logger configuration
type LogConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Stdout bool   `mapstructure:"stdout"`
	File   struct {
		Path       string `mapstructure:"path"`
		MaxSize    int    `mapstructure:"max_size"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxAge     int    `mapstructure:"max_age"`
		Compress   bool   `mapstructure:"compress"`
	} `mapstructure:"file"`
}

// LoadConfig loads the application configuration from file
func LoadConfig(configPath string) (*AppConfig, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &appConfig, nil
}