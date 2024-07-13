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

	log.Print("cooking db...")
	// db, err := postgresql.NewPosgresDB(postgresql.DBConfig{
	// 	Host:     "localhost",
	// 	Port:     "8888",
	// 	Username: "bionic4l",
	// 	Password: "12345",
	// 	DBName:   "fm_db",
	// 	SSLMode:  "disable",
	// })
	// if err != nil {
	// 	log.Fatal("db not cooked")
	// }

	//repo := repository.NewRepository(db)

	// log.Print("cooking logger...")

	return nil
}
