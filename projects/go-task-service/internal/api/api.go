package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-task-service/internal/config"
)

type API struct {
	router http.Handler
	cfg    *config.Config
	logger *zerolog.Logger
}

func NewAPI(cfg *config.Config, handler http.Handler, log *zerolog.Logger) (*API, error) {
	return &API{
		router: handler,
		logger: log,
		cfg:    cfg,
	}, nil
}

// Run starts the api server
func (a *API) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:         a.cfg.Addr(),
		Handler:      a.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		a.logger.Info().Msg("Server is shutting down")

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		a.logger.Info().Msg("Server shutdown completed")

		shutdownError <- nil
	}()

	var err error
	if a.cfg.TLS != nil {
		log.Info().Str("addr", srv.Addr).Bool("tls", true).Msg("Serving HTTPS")
		err = srv.ListenAndServeTLS(a.cfg.TLS.CertFile, a.cfg.TLS.KeyFile)
	} else {
		log.Info().Str("addr", srv.Addr).Bool("tls", false).Msg("Serving HTTP")
		err = srv.ListenAndServe()
	}

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	a.logger.Info().Str("addr", srv.Addr).Msg("Server stopped successfully")

	return nil
}
