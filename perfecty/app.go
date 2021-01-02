package perfecty

import (
	"fmt"
	"github.com/rs/zerolog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const filePath = "config/perfecty.yml"

// Load Setup the server configuration and logging
func Load() (err error) {
	// config
	if err = loadConfig(filePath); err != nil {
		return
	}

	// logging
	if err = loadLogging(); err != nil {
		return
	}

	return
}

// Config

type (
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Logging struct {
		Level string `yaml:"level"`
	}
	Config struct {
		Server  Server  `yaml:"server"`
		Logging Logging `yaml:"logging"`
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
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if config.Logging.Level != "" {
		var level zerolog.Level
		if level, err = zerolog.ParseLevel(config.Logging.Level); err != nil {
			return
		}
		zerolog.SetGlobalLevel(level)
	}

	return
}
