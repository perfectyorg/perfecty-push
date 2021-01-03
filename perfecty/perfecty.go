package perfecty

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const filePath = "config/perfecty.yml"

// Start Setup and start the push server
func Start() (err error) {
	if err = loadConfig(filePath); err != nil {
		return
	}

	if err = loadLogging(); err != nil {
		return
	}

	if err = startServer(); err != nil {
		return
	}

	return
}

// Config

type (
	Ssl struct {
		Enabled  bool   `yaml:"enabled"`
		CertFile string `yaml:"cert_file"`
		KeyFile  string `yaml:"key_file"`
	}
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Ssl  Ssl
	}
	Logging struct {
		Level  string `yaml:"level"`
		Pretty bool   `yaml:"pretty"`
	}
	Config struct {
		Server  Server
		Logging Logging
	}
)

var (
	config Config
)

func loadConfig(filePath string) (err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("cannot find the config file at %s: %w", filePath, err)
		return
	}

	if err = yaml.Unmarshal(content, &config); err != nil {
		err = fmt.Errorf("cannot parse the config file: %w", err)
		return
	}

	return
}

// Logging

func loadLogging() (err error) {
	var (
		host  string
		level zerolog.Level
	)
	host, _ = os.Hostname()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Logging.Level != "" {
		if level, err = zerolog.ParseLevel(config.Logging.Level); err != nil {
			return
		}
		zerolog.SetGlobalLevel(level)
	}

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("host", host).
		Logger()
	if config.Logging.Pretty == true {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	log.Logger = logger

	return
}
