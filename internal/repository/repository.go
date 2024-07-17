package repository

import (
	dbactions "finance_manager/internal/repository/db_actions"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	dbactions.Balance
	dbactions.Deposit
	dbactions.Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
