package perfecty

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	Config struct {
		Server Server `yaml:"server"`
	}
)

var (
	config Config
)

func LoadConfig(filePath string) (err error) {
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
