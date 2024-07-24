package dbactions

import (
	"errors"
	"finance_manager/internal/models"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const queryUpdateUserBalance = `UPDATE users SET balance = $1 WHERE id = $2`

type TransactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t TransactionRepository) Transaction(fromID, toID, amount int) error {
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
		log.Error("недостаточно средств")
		return errors.New("отрицательный баланс")
	}

	if amount < 1 {
		log.Error("перевод меньше минимальной суммы")
		return errors.New("перевод меньше минимальной суммы")
	}

	t.DBUpdateBalanceTransaction(fromUser, toUser, amount)

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
	_, err := t.db.Exec(queryUpdateUserBalance, fromUser.ID, fromUser.Balance-amount)
	if err != nil {
		return err
	}

	_, errr := t.db.Exec(queryUpdateUserBalance, toUser.ID, toUser.Balance+amount)
	if errr != nil {
		return errr
	}
	return nil
}
