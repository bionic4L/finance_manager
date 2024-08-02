package app

import (
	"context"
	v1 "finance_manager/internal/api/v1"
	"finance_manager/internal/config"
	"finance_manager/internal/db/postgresql"
	"finance_manager/internal/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	// graceful shutdown
	exitSig := make(chan os.Signal, 1)
	signal.Notify(exitSig, syscall.SIGINT, syscall.SIGTERM)
	<-exitSig
	log.Info("ready for break down that shit ~gracefully~")
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	log.Info("waiting for all proccesses done...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("error while shutting down server: %s", err)
		return err
	}
	log.Info("program has gracefully downed")
	// for {
	// 	select {
	// 	case <-ctx.Done():
	//
	// 		srv.Shutdown(ctx)
	// 		return nil
	// 	case <-time.After(1 * time.Second): //я так понимаю эта тема как раз для крона раз в месяц
	// 		log.Info("5 seconds gone")
	// 	}
	// }

	return nil

}
