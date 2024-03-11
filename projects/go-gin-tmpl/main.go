package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/roarc0/go-gin-tmpl/database"
)

func main() {
	cfg := NewConfig()
	articleService, err := database.NewArticleService(context.Background(), cfg.ConnectionURI, cfg.DBName, cfg.CollectionName)
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to the database")
	}

	router := getBlogRouter(articleService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Listen")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server Shutdown")
	}

	<-ctx.Done()
	log.Info().Msg("Server exiting")
}
