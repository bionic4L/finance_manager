package app

import (
	v1 "finance_manager/internal/api/v1"
	"finance_manager/internal/config"
	"finance_manager/internal/db/postgresql"
	"finance_manager/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
)

func Run(configPath string) error {
	log.Print("cooking config...")
	cfg, err := config.GetConfig()
	router := gin.Default()

	if err != nil {
		log.Fatal("config not cooked :(")
	}
	log.Print("config cooked!")
	log.Print("cooking db...")
	db, err := postgresql.OpenPosgresDB(&config.PostgreSQL_DB{
		Host:     cfg.PostgreSQL_DB.Host,
		Port:     cfg.PostgreSQL_DB.Port,
		User:     cfg.PostgreSQL_DB.User,
		Password: cfg.PostgreSQL_DB.Password,
		DBName:   cfg.PostgreSQL_DB.DBName,
		SSLMode:  cfg.PostgreSQL_DB.SSLMode,
	})
	if err != nil {
		log.Print(err)
		log.Fatal("db not cooked :(")
	}
	defer postgresql.CloseConnection(db.DB)
	log.Print("db is active!")

	log.Print("creating repository...")
	rootRepository := repository.NewRepository(db)
	log.Print("repository cooked!")

	log.Print("cooking router...")
	v1.Router(router, rootRepository)
	log.Print("router cooked!")

	log.Print("*****starting*****")

	// log.Print("cooking logger...")

	router.Run(cfg.HTTPServer.Address)

	return nil
}
