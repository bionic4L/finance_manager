package repository

import (
	"finance_manager/internal/models"
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

// func NewRepository(DB *sqlx.DB) *Repository {
// 	return &Repository{
// 		BalanceRepository: BalanceRepository,
// 		DepositRepository: DepositRepository,
// 	}
// }
