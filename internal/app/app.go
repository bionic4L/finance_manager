package app

import (
	v1 "finance_manager/internal/api/v1"
	"finance_manager/internal/config"
	"finance_manager/internal/db/postgresql"
	"log"
)

func Run(configPath string) error {
	log.Print("cooking config...")
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatal("config not cooked :(")
	}
	log.Print("config cooked!")
	log.Print("cooking db...")
	if err := postgresql.OpenPosgresDB(&config.PostgreSQL_DB{
		Host:     cfg.PostgreSQL_DB.Host,
		Port:     cfg.PostgreSQL_DB.Port,
		User:     cfg.PostgreSQL_DB.User,
		Password: cfg.PostgreSQL_DB.Password,
		DBName:   cfg.PostgreSQL_DB.DBName,
		SSLMode:  cfg.PostgreSQL_DB.SSLMode,
	}); err != nil {
		log.Print(err)
		log.Fatal("db not cooked :(")
	}
	log.Print("db is active!")

	log.Print("cooking router...")
	v1.Router(cfg.HTTPServer.Address)
	log.Print("router cooked!")

	//repo := repository.NewRepository(db)

	// log.Print("cooking logger...")

	return nil
}
