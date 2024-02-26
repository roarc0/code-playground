package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var fConfigFile string
	const (
		fConfigFileDefault = "server.yaml"
		fConfigFileDesc    = "select the config file"
	)

	fDebug := flag.Bool("debug", false, "sets log level to debug")
	flag.StringVar(&fConfigFile, "config", fConfigFileDefault, fConfigFileDesc)
	flag.StringVar(&fConfigFile, "c", fConfigFileDefault, fConfigFileDesc+" (shorthand)")
	flag.Parse()

	setupLogger(fDebug)

	cfg, err := readConfig(fConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	if err := serve(ctx, cfg); err != nil {
		log.Error().Err(err).Msg("Serve Error")
		os.Exit(1)
	}
}

func setupLogger(fDebug *bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *fDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime})
}
