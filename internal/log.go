package internal

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Logging

func LoadLogging() (err error) {
	var (
		host  string
		level zerolog.Level
	)
	host, _ = os.Hostname()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if Config.Logging.Level != "" {
		if level, err = zerolog.ParseLevel(Config.Logging.Level); err != nil {
			return
		}
		zerolog.SetGlobalLevel(level)
	}

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("host", host).
		Logger()
	if Config.Logging.Pretty == true {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Logger = logger

	return
}
