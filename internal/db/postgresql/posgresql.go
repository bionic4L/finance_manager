package postgresql

import (
	"database/sql"
	"finance_manager/internal/config"
	_ "finance_manager/internal/db/postgresql/migrations"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	"github.com/jmoiron/sqlx"
)

func OpenPosgresDB(cfg *config.PostgreSQL_DB) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	DB := db.DB
	log.Print("creating migrations...")
	if err := goose.Up(DB, "D:/LRN GO/finance_manager/internal/db/postgresql/migrations"); err != nil {
		log.Print("migrations not applied")
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
