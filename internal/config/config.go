package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Environment string `yaml:"env" env-default:"local"`
		// DBPath string `yaml:"db_path" env-required:"true"`
		HTTPServer
		LoggerLevel string `yaml:"logger_level" env-default:"debug"`
	}

	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"localhost:8888" env-required:"true"`
		Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}
)

func GetConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH") // не грузится!!!!!!!!!!!!!!!!

	if configPath == "" {
		log.Fatal("can't find CONFIG_PATH in .env...")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file NOT EXISTS!...")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Fatal("can't read config file...")
	}

	return &cfg

}
