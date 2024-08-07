package repository

import (
	"context"
	"finance_manager/internal/models"
	dbactions "finance_manager/internal/repository/db_actions"

	"github.com/jmoiron/sqlx"
)

type BalanceRepository interface {
	GetUserBalance(ctx context.Context, id int) (*models.User, error)
}

type DepositRepository interface {
	Deposit(ctx context.Context, id int, amount int) error
}

type UserCreateRepository interface {
	UserCreate(ctx context.Context, name string) error
}

type TransactionRepository interface {
	Transaction(ctx context.Context, fromID, toID, amount int) error
}

type Repository struct {
	BalanceRepository
	DepositRepository
	UserCreateRepository
	TransactionRepository
}

func NewRepository(DB *sqlx.DB) *Repository {
	return &Repository{
		BalanceRepository:     dbactions.NewBalanceRepository(DB),
		DepositRepository:     dbactions.NewDepositRepository(DB),
		UserCreateRepository:  dbactions.NewCreateUserRepository(DB),
		TransactionRepository: dbactions.NewTransactionRepository(DB),
	}
}
