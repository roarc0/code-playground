package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

// Config contains the application configuration
type Config struct {
	ListenAddress string `yaml:"listen-address"`
	Verbose       bool   `yaml:"verbose"`
}

var defaultConfig = Config{
	ListenAddress: ":8080",
	Verbose:       true,
}

// readConfig reads the configuration from the given file
func readConfig(cfgFile string) (*Config, error) {
	bytes, err := os.ReadFile(cfgFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &defaultConfig, nil
		}
		return nil, err
	}

	bytes = []byte(os.ExpandEnv(string(bytes)))

	var cfg Config
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// loadConfig loads the configuration, falling back to defaults if necessary
func loadConfig(cfgFile string) (*Config, error) {
	if cfgFile == "" {
		return nil, errors.New("no config file provided")
	}

	return readConfig(cfgFile)
}
