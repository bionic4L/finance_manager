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
	// router.Run(cfg.Address)
	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: router.Handler(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Print(srv.Addr)
	}()

	// graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("some issues while shutting down server")
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Info("ready for break down that shit (gracefully)...")
			srv.Shutdown(ctx)
			return nil
		case <-time.After(60 * time.Second): //я так понимаю эта тема как раз для крона раз в месяц
			log.Info("60 seconds gone")
		}
	}
	// return nil
}
