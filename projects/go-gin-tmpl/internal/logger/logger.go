package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

// SetupLogger the logger
func SetupLogger(debug bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime})
}
