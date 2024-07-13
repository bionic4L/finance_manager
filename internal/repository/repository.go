package repository

import "github.com/jmoiron/sqlx"

type (
	CheckBalance interface {
		GetUserBalance()
	}

	Deposit interface {
		DepositToUser()
	}

	Transaction interface {
		Transaction()
	}
)

type Repository struct {
	CheckBalance
	Deposit
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
