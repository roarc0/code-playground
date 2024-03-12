package api

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/roarc0/go-gin-tmpl/internal/api/handlers"
	"github.com/roarc0/go-gin-tmpl/internal/config"
)

type API struct {
	router *gin.Engine
	cfg    *config.Config
	logger *zerolog.Logger
	wg     *sync.WaitGroup
}

func NewAPI(cfg *config.Config, handlers *handlers.Handlers, log *zerolog.Logger) (*API, error) {
	router := handlers.BlogHandler
	router.Use(gin.Recovery())

	return &API{
		router: router,
		logger: log,
		cfg:    cfg,
		wg:     &sync.WaitGroup{},
	}, nil
}

// Run starts the api server
func (a *API) Run() error {
	srv := &http.Server{
		Addr:         a.cfg.ListenAddress(),
		Handler:      a.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sign := <-quit

		a.logger.Info().Any("signal", sign.String()).Msg("Caught signal")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		a.logger.Info().Any("addr", srv.Addr).Msg("Completing background tasks")

		a.wg.Wait()
		shutdownError <- nil
	}()

	a.logger.Info().Any("addr", srv.Addr).Msg("Starting server")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	a.logger.Info().Any("addr", srv.Addr).Msg("Stopped server")

	return nil
}
