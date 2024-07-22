package dbactions

import (
	"errors"
	"finance_manager/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	queryGetUserInfoByID     = `SELECT * FROM users WHERE id = $1`
	queryUpdateuserInfo      = `UPDATE users SET balance = $1 WHERE id = $2`
	queryInsertInDepositList = `INSERT INTO deposits (user_id, amount) VALUES ($1, $2)`
)

type DepositRepository struct {
	db *sqlx.DB
}

func NewDepositRepository(db *sqlx.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

func (d *DepositRepository) Deposit(user_id int, depAmount int) error {
	userInfo, err := d.DBSelectUserInfoByID(user_id)
	if err != nil {
		log.Print("error while selecting balance")
		return err
	}

	if userInfo.Balance+depAmount < 0 {
		log.Printf("Отрицательный баланс: %d", userInfo.Balance+depAmount)
		return errors.New("отрицательный баланс")
	}

	if err := d.DBUpdateBalance(userInfo.ID, userInfo.Balance, depAmount); err != nil {
		log.Print("error while updating balance")
		return err
	}

	if err := d.DBInsertDepInfo(userInfo.ID, depAmount); err != nil {
		log.Print("error while inserting deposit info")
		return err
	}

	return nil
}

func (d *DepositRepository) DBSelectUserInfoByID(user_id int) (*models.User, error) {
	var user models.User

	row := d.db.QueryRow(queryGetUserInfoByID, user_id)

	if err := row.Scan(&user.ID, &user.Name, &user.Balance); err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *DepositRepository) DBUpdateBalance(user_id int, balance int, depAmount int) error {
	_, err := d.db.Exec(queryUpdateuserInfo, balance+depAmount, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (d *DepositRepository) DBInsertDepInfo(user_id int, depAmount int) error {
	_, err := d.db.Exec(queryInsertInDepositList, user_id, depAmount)
	if err != nil {
		return err
	}

	return nil
}
