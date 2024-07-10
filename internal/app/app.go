package app

import (
	v1 "finance_manager/internal/api/v1"
	"finance_manager/internal/config"
	"log"
)

func Run(configPath string) error {
	log.Print("cooking config...")
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatal("config not cooked :(")
	}
	log.Print("config cooked!")

	log.Print("cooking router...")
	v1.Router(cfg.HTTPServer.Address)
	log.Print("router cooked!")

	// log.Print("cooking logger...")

	return nil
}
