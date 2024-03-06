package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.comroarc0/go-gin-todo/controllers"
	"github.comroarc0/go-gin-todo/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	var fConfigFile string
	const (
		fConfigFileDefault = "todo.yaml"
		fConfigFileDesc    = "select the config file"
	)

	flag.StringVar(&fConfigFile, "config", fConfigFileDefault, fConfigFileDesc)
	flag.StringVar(&fConfigFile, "c", fConfigFileDefault, fConfigFileDesc+" (shorthand)")
	flag.Parse()

	cfg, err := readConfig(fConfigFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read config")
	}

	utils.SetupLogger(cfg.Verbose)

	db, err := utils.OpenDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect database")
	}

	router := controllers.NewTodoController(db)
	router.Use(gin.Recovery())

	srv := &http.Server{
		Addr:           cfg.ListenAddress,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Listen error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server shutdown error")
	}
	log.Info().Msg("Server exiting")
}
