package config

import (
	"go-rest-api-template/pkg/logging"
	"gopkg.in/yaml.v2"
	"os"
	"sync"
)

type Config struct {
	IsDebug bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}

		f, err := os.Open("config.yml")
		if err != nil {
			logger.Fatal(err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				logger.Fatal(err)
			}
		}(f)

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&instance)

		if err != nil {
			logger.Fatal(err)
		}
	})
	return instance
}
