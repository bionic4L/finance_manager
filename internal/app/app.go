package app

import (
	v1 "finance_manager/internal/api/v1"
	"finance_manager/internal/config"
	"finance_manager/internal/db/postgresql"
	"finance_manager/internal/repository"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Run(configPath string) error {
	cfg, err := config.GetConfig()

	setLogLevel(cfg.LoggerLevel)
	log.Info("logger cooked!")

	router := gin.Default()

	if err != nil {
		log.Fatal("config not cooked :(")
	}
	log.Info("config cooked!")

	log.Info("cooking db...")
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
	log.Info("db is active!")

	log.Info("creating repository...")
	rootRepository := repository.NewRepository(db)
	log.Info("repository cooked!")

	log.Info("cooking router...")
	v1.Router(router, rootRepository)
	log.Info("router cooked!")

	log.Info("*****starting*****")
	router.Run(cfg.HTTPServer.Address) // go func

	//router shutdown

	// srv := http.Server{
	// 	Addr:    cfg.HTTPServer.Address,
	// 	Handler: router.Handler(),
	// }

	// srv.Shutdown()

	return nil
}
