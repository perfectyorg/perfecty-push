package perfecty

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rwngallego/perfecty-push/internal/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const filePath = "configs/internal.yml"

// Start Setup and start the push server
func Start() (err error) {
	if err = loadConfig(filePath); err != nil {
		return
	}

	if err = loadLogging(); err != nil {
		return
	}

	if err = server.StartServer(); err != nil {
		return
	}

	return
}

// PerfectyConfig

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
	PerfectyConfig struct {
		Server  Server
		Logging Logging
	}
)

var (
	Config PerfectyConfig
)

func loadConfig(filePath string) (err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		err = fmt.Errorf("cannot find the configs file at %s: %w", filePath, err)
		return
	}

	if err = yaml.Unmarshal(content, &Config); err != nil {
		err = fmt.Errorf("cannot parse the configs file: %w", err)
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
