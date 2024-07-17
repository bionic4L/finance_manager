package repository

import "finance_manager/internal/models"

type BalanceRepository interface {
	GetUserBalance(id int) (*models.User, error)
}

type DepositRepository interface {
	Deposit(amount int) (*models.User, error)
}
