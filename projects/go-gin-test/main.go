package main

import (
	"flag"

	"github.com/rs/zerolog/log"
)

func main() {
	var fConfigFile string
	const (
		fConfigFileDefault = "server.yaml"
		fConfigFileDesc    = "select the config file"
	)

	flag.StringVar(&fConfigFile, "config", fConfigFileDefault, fConfigFileDesc)
	flag.StringVar(&fConfigFile, "c", fConfigFileDefault, fConfigFileDesc+" (shorthand)")
	flag.Parse()

	cfg, err := readConfig(fConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}

	setupLogger(cfg.Verbose)

	router := GetMainRouter()
	router.Run(cfg.ListenAddress)
}
