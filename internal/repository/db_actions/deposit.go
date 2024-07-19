package dbactions

import (
	"errors"
	"finance_manager/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	queryGetUserBalance    = "SELECT balance FROM users WHERE user_id = $1"
	queryUpdateUserBalance = "UPDATE users SET balance = $1 WHERE user_id = $2"
)

type DepositRepository struct {
	db *sqlx.DB
}

func NewDepositRepository(db *sqlx.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

func (d *DepositRepository) Deposit(user_id int, amount int) error {

	userBalance, err := d.DBSelectBalance(user_id)
	if err != nil {
		log.Print("error while selecting balance")
		return err
	}

	if userBalance+amount < 0 {
		log.Printf("Отрицательный баланс: %d", userBalance+amount)
		return errors.New("отрицательный баланс")
	}

	if err := d.DBUpdateBalance(user_id, userBalance, amount); err != nil {
		log.Print("error while updating balance")
		return err
	}

	return nil
}

func (d *DepositRepository) DBSelectBalance(user_id int) (int, error) {
	var user models.User

	row := d.db.QueryRow(queryGetUserBalance, user_id)

	if err := row.Scan(&user.Balance); err != nil {
		return user.Balance, err
	}

	return user.Balance, nil
}

func (d *DepositRepository) DBUpdateBalance(user_id int, balance int, amount int) error {
	_, err := d.db.Exec(queryUpdateUserBalance, balance+amount, user_id)
	if err != nil {
		return err
	}
	return nil
}
