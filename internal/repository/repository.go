package repository

import (
	"finance_manager/internal/models"
	dbactions "finance_manager/internal/repository/db_actions"

	"github.com/jmoiron/sqlx"
)

type BalanceRepository interface {
	GetUserBalance(id int) (*models.User, error)
}

type DepositRepository interface {
	Deposit(id int, amount int) error
}

type Repository struct {
	BalanceRepository
	DepositRepository
}

func NewRepository(DB *sqlx.DB) *Repository {
	return &Repository{
		BalanceRepository: dbactions.NewBalanceRepository(DB),
		DepositRepository: dbactions.NewDepositRepository(DB),
	}
}
