package dbactions

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
)

type DepositRepository struct {
	DB *sqlx.DB
}

func (d *DepositRepository) DepositToUser(id int, amount int) error {
	var userBalance int
	queryGetUserBalance := "SELECT balance FROM users WHERE id = 1$"
	queryUpdateUserBalance := "UPDATE users SET balance = $1 WHERE id = $2"

	if err := d.DB.Get(&userBalance, queryGetUserBalance, id); err != nil {
		return err
	}

	_, err := d.DB.Exec(queryUpdateUserBalance, userBalance+amount, id)
	if err != nil {
		return err
	}

	if userBalance+amount < 0 {
		log.Printf("Отрицательный баланс: %d", userBalance+amount)
		return errors.New("отрицательный баланс")
	}

	return nil
}
