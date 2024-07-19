package config

import (
	"errors"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const PATH = "D:/LRN GO/finance_manager/config/local.yaml"

type (
	Config struct {
		Environment   string `yaml:"env" env-default:"local"`
		HTTPServer    `yaml:"http_server"`
		PostgreSQL_DB `yaml:"postgres_db"`
		LoggerLevel   string `yaml:"logger_level" env-default:"debug"`
	}

	HTTPServer struct {
		Address string `yaml:"address" env-default:"localhost:8888"`
	}

	PostgreSQL_DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
		SSLMode  string `yaml:"ssl_mode"`
	}
)

func GetConfig() (*Config, error) {
	configPath := PATH
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

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil { //это и возвращает конфиг. остальное проверки
		log.Println(err)
		err = errors.New("can't read config file")
		return cfg, err
	}

	return cfg, nil

}
