package pkg

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func Init() {
	log.Logger = logger
}

func Error(err error) {
	log.Logger.Error().Err(err)
}

func Info(msg string) {
	log.Logger.Info().Msg(msg)
}

func Panic(err error) {
	log.Logger.Panic().Err(err)
}
