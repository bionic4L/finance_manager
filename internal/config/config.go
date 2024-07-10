package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

const PATH = "C:/Users/rozzo/OneDrive/Рабочий стол/LRN GO/finance_manager/config/local.yaml"

type (
	Config struct {
		Environment string `yaml:"env" env-default:"local"`
		// DBPath string `yaml:"db_path" env-required:"true"`
		HTTPServer  `yaml:"http_server"`
		LoggerLevel string `yaml:"logger_level" env-default:"debug"`
	}

	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"localhost:8888"`
		Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}
)

func GetConfig() (*Config, error) {
	configPath := PATH // не грузится!!!!!!!!!!!!!!!!
	cfg := &Config{}
	var err error

	if configPath == "" {
		err = errors.New("can't find CONFIG_PATH in .env")
		return cfg, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = errors.New("config file NOT EXISTS")
		return cfg, err
	}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		log.Println(err)
		err = errors.New("can't read config file")
		return cfg, err
	}

	return cfg, nil

}
