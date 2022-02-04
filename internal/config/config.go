package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/picunada/http-rest-balance/pkg/logging"
	"sync"
)

type ServerConfig struct {
	IsDebug *bool `yaml:"is_debug" env-default:"true" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bindIP" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8001"`
	} `yaml:"listen"`
	Storage DatabaseConfig `yaml:"storage"`
}

type DatabaseConfig struct {
	Port        string `yaml:"port" env:"PORT" env-default:"5432"`
	Host        string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name        string `yaml:"name" env:"NAME" env-default:"postgres"`
	User        string `env:"USER" env-default:"user"`
	Password    string `env:"PASSWORD"`
	MaxAttempts int    `yaml:"max_attempts" env-default:"5"`
}

var instance *ServerConfig
var once sync.Once

func GetConfig() *ServerConfig {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read application info")
		instance = &ServerConfig{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})

	return instance
}
