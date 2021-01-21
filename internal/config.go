package internal

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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

func LoadConfig(filePath string) (err error) {
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
