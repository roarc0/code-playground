package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	_ "honnef.co/go/tools/config"

	"github.com/roarc0/go-gin-tmpl/internal/api"
	"github.com/roarc0/go-gin-tmpl/internal/api/handlers"
	"github.com/roarc0/go-gin-tmpl/internal/config"
	"github.com/roarc0/go-gin-tmpl/internal/database"
	"github.com/roarc0/go-gin-tmpl/internal/logger"
)

func main() {
	logger.SetupLogger(false)

	godotenv.Load(".env.dev")

	cfg := config.NewConfig()

	err := cfg.ParseFlags()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse command-line flags")
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to the db")
	}

	h := handlers.NewHandlers(db)

	srv, err := api.NewAPI(cfg, h, &log.Logger)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to setup the server")
	}

	err = srv.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
