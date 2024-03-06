package main

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/roarc0/go-task-service/task"
)

// Config contains a really basic the application configuration
type Config struct {
	ListenAddress string           `yaml:"listen-address"`
	TLS           *TLSConfig       `yaml:"tls,omitempty"`
	TaskStore     task.StoreConfig `yaml:"task-store,omitempty"`
}

var defaultConfig = Config{
	ListenAddress: "localhost:8080",
	TaskStore: task.StoreConfig{
		Type: "memory",
	},
}

// TLSConfig is used to configure the TLS
type TLSConfig struct {
	CertFile string `yaml:"cert-file"`
	KeyFile  string `yaml:"key-file"`
}

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
	return &cfg, err
}
