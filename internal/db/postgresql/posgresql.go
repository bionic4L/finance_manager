package postgresql

import (
	"database/sql"
	"finance_manager/internal/config"
	_ "finance_manager/internal/db/postgresql/migrations"
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"github.com/jmoiron/sqlx"
)

func OpenPosgresDB(cfg *config.PostgreSQL_DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		log.Error("cannot open the DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Error("pinged DB but there is no answer")
		return nil, err
	}

	//migration down flag
	downFlag := flag.Int64("down", -1, "down 1 migration")
	flag.Parse()

	DB := db.DB
	log.Print("cooking migrations...")
	if *downFlag != -1 {
		if err := goose.DownTo(DB, "./internal/db/postgresql/migrations", *downFlag); err != nil {
			log.Warn("migrations not downed")
			return nil, err
		}
		log.Info("1 migration downed")
	} else {
		if err := goose.Up(DB, "./internal/db/postgresql/migrations"); err != nil {
			log.Warn("migrations not applied")
			return nil, err
		}
	}

	// if err := goose.Down(DB, "D:/LRN GO/finance_manager/internal/db/postgresql/migrations"); err != nil {
	// 	log.Print("migrations not applied")
	// 	return nil, err
	// }

	return db, nil
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
