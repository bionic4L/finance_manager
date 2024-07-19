package dbactions

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	queryGetUserBalance    = "SELECT balance FROM users WHERE id = 1$"
	queryUpdateUserBalance = "UPDATE users SET balance = $1 WHERE id = $2"
)

type DepositRepository struct {
	DB *sqlx.DB
}

func NewDepositRepository(DB *sqlx.DB) *DepositRepository {
	return &DepositRepository{DB: DB}
}

func (d *DepositRepository) Deposit(id int, amount int) error {

	userBalance, err := d.DBSelectBalance(id)
	if err != nil {
		log.Print("error while selecting balance")
		return err
	}

	if userBalance+amount < 0 {
		log.Printf("Отрицательный баланс: %d", userBalance+amount)
		return errors.New("отрицательный баланс")
	}

	if err := d.DBUpdateBalance(id, userBalance, amount); err != nil {
		log.Print("error while updating balance")
		return err
	}

	return nil
}

func (d *DepositRepository) DBSelectBalance(id int) (int, error) {
	var userBalance int

	if err := d.DB.Get(&userBalance, queryGetUserBalance, id); err != nil {
		return -1, err
	}

	return userBalance, nil
}

func (d *DepositRepository) DBUpdateBalance(id int, balance int, amount int) error {
	_, err := d.DB.Exec(queryUpdateUserBalance, balance+amount, id)
	if err != nil {
		return err
	}
	return nil
}
