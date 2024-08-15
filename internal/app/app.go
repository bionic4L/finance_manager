package app

import (
	"context"
	v1 "finance_manager/internal/api/v1"
	"finance_manager/internal/config"
	"finance_manager/internal/db/postgresql"
	"finance_manager/internal/repository"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Run(configPath string) error {
	cfg, err := config.GetConfig()

	setLogLevel(cfg.LoggerLevel)
	log.Info("logger cooked!")

	router := gin.Default()

	if err != nil {
		log.Error("config not cooked :(")
		return err
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
		log.Error("db not cooked :(")
		return err
	}
	defer postgresql.CloseConnection(db.DB)
	log.Info("db is active!")

	log.Info("cooking repository...")
	rootRepository := repository.NewRepository(db)
	log.Info("repository cooked!")

	log.Info("cooking router...")
	v1.Router(router, rootRepository)
	log.Info("router cooked!")

	log.Info("*****starting*****")
	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: router.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("cannot start srv listening cause: %s\n", err)
		}
	}()

	//cron
	c := cron.New()
	c.AddFunc("@every 00h00m10s", func() {
		Report(context.Background(), db)
	})
	go func() {
		c.Start()
		log.Info("запустили крон в отдельной горутине")
	}()

	// graceful shutdown
	// exitSig := make(chan os.Signal, 1)
	// signal.Notify(exitSig, syscall.SIGINT, syscall.SIGTERM)
	// <-exitSig

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			log.Info("ready for break down that shit ~gracefully~")

			if err := srv.Shutdown(ctx); err != nil {
				log.Errorf("error while shutting down server: %s", err)
				return err
			}
			c.Stop() //stopping cron
			log.Info("program has gracefully downed")
			return nil
		case <-time.After(15 * time.Second): //а нужно ли? если в кроне и так ставится период повтора
			log.Info("5 seconds gone")
			// if err := Report(ctx, db); err != nil {
			// 	log.Error(err)
			// }
		}
	}

	//return nil

}
