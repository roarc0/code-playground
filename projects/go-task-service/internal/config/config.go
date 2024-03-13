package config

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

// Config contains a really basic the application configuration
type Config struct {
	Debug         bool        `env:"DEBUG, default=true"`
	ListenAddress string      `env:"LISTEN_ADDRESS, default=localhost"`
	Port          int         `env:"PORT, default=8080"`
	TLS           *TLSConfig  `env:", noinit, prefix=TLS_"`
	TaskStore     StoreConfig `env:", prefix=STORE_"`
}

func (cfg *Config) Addr() string {
	return fmt.Sprintf("%s:%d", cfg.ListenAddress, cfg.Port)
}

// StoreConfig is used to configure the
type StoreConfig struct {
	Type   string `env:"TYPE, default=memory"`
	Params any    `env:"PARAMS"`
}

// TLSConfig is used to configure the TLS
type TLSConfig struct {
	CertFile string `env:"TLS_CERT_FILE"`
	KeyFile  string `env:"TLS_KEY_FILE"`
}

// Load creates a new configuration reading the properties from the environment
func Load(ctx context.Context) (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := envconfig.Process(ctx, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
