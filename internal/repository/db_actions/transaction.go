package dbactions

import (
	"context"
	"errors"
	"finance_manager/internal/models"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const queryUpdateUserBalanceMinus = `UPDATE users SET balance = balance - $1 WHERE id = $2`
const queryUpdateUserBalancePlus = `UPDATE users SET balance = balance + $1 WHERE id = $2`
const queryInsertTransactionInfo = `INSERT INTO transactions (from_id, to_id, amount) VALUES ($1, $2, $3)`

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t TransactionRepository) Transaction(ctx context.Context, fromID, toID, amount int) error {
	fromUser, err := t.DBSelectUserInfoByID(fromID)
	if err != nil {
		return err
	}

	toUser, err := t.DBSelectUserInfoByID(toID)
	if err != nil {
		return err
	}

	if fromID == toID {
		log.Warn("нельзя перевести средства самому себе")
		return errors.New("нельзя перевести средства самому себе")
	}

	if fromUser.Balance-amount < 0 {
		return errors.New("недостаточно средств")
	}

	if amount < 1 {
		return errors.New("перевод меньше минимальной суммы")
	}

	if err := t.DBUpdateBalanceTransaction(fromUser, toUser, amount); err != nil {
		return err
	}
	if err := t.DBInsertTransactionInfo(fromID, toID, amount); err != nil {
		return err
	}

	return nil
}

func (t *TransactionRepository) DBSelectUserInfoByID(user_id int) (*models.User, error) {
	var user models.User
	row := t.db.QueryRow(queryGetUserInfoByID, user_id)

	if err := row.Scan(&user.ID, &user.Name, &user.Balance); err != nil {
		return nil, err
	}

	return &user, nil
}

func (t *TransactionRepository) DBUpdateBalanceTransaction(fromUser, toUser *models.User, amount int) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(queryUpdateUserBalanceMinus, amount, fromUser.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	_, err = tx.Exec(queryUpdateUserBalancePlus, amount, toUser.ID)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil
}

func (d *TransactionRepository) DBInsertTransactionInfo(fromID int, toID int, amount int) error {
	_, err := d.db.Exec(queryInsertTransactionInfo, fromID, toID, amount)
	if err != nil {
		return err
	}

	return nil
}
