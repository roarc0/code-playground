package main

import (
	"context"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-task-service/controllers"
	"github.com/roarc0/go-task-service/task"
)

func serve(ctx context.Context, cfg *Config) error {
	taskRunner := task.NewRunner()
	taskStore, err := task.StoreFactory(cfg.TaskStore, taskRunner)
	if err != nil {
		return err
	}
	taskStore.Start(ctx)
	taskController := controllers.NewTaskController(taskStore)

	srv := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: taskController.Handler(),
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		log.Info().Msg("Server shutting down...")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer shutdownCancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("Failed to shutdown the server")
		}

		log.Info().Msg("Server shutdown completed")
	}()

	if cfg.TLS != nil {
		log.Info().Str("listenAddress", cfg.ListenAddress).Bool("tls", true).Msg("Serving HTTPS")
		err = srv.ListenAndServeTLS(cfg.TLS.CertFile, cfg.TLS.KeyFile)
	} else {
		log.Info().Str("listenAddress", cfg.ListenAddress).Bool("tls", false).Msg("Serving HTTP")
		err = srv.ListenAndServe()
	}

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("ListenAndServe returned an error")
		cancel()
	}

	wg.Wait()
	return nil
}
