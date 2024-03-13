package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-task-service/internal/api"
	"github.com/roarc0/go-task-service/internal/api/controllers"
	"github.com/roarc0/go-task-service/internal/config"
	"github.com/roarc0/go-task-service/internal/logger"
)

func main() {
	logger.SetupLogger()

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg, err := config.Load(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	logger.SetDebug(cfg.Debug)

	taskController := controllers.NewDefaultTaskController(ctx, cfg)
	srv, err := api.NewAPI(cfg, taskController.Handler(), &log.Logger)
	if err != nil {
		log.Fatal().Err(err).Msg("API Creation Error")
	}

	if err := srv.Run(ctx); err != nil {
		log.Fatal().Err(err).Msg("API Run Error")
	}
}
