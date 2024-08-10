package dbactions

import (
	"context"
	"errors"
	"finance_manager/internal/models"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const queryUpdateUserBalancePlus = `UPDATE users SET balance = balance + $1 WHERE id = $2`
const queryGetTransactionInfoByID = `SELECT * FROM transactions WHERE transactions_id = $1`
const queryUpdateStatusNoMorePending = `UPDATE transactions SET is_pending = FALSE WHERE transactions_id = $1 `
const queryUpdateStatusDone = `UPDATE transactions SET is_done = TRUE WHERE transactions_id = $1`
const queryUpdateStatusCanceled = `UPDATE transactions SET is_canceled = TRUE WHERE transactions_id = $1`

type StatusControllerRepository struct {
	db *sqlx.DB
}

func NewStatusControllerRepository(db *sqlx.DB) *StatusControllerRepository {
	return &StatusControllerRepository{db: db}
}

func (scr *StatusControllerRepository) StatusController(ctx context.Context, transactionID int, confirmTransaction bool) error {
	row := scr.db.QueryRowContext(ctx, queryGetTransactionInfoByID, transactionID)

	var transactionModel models.Transaction

	if err := row.Scan(&transactionModel.ID, &transactionModel.FromID, &transactionModel.ToID, &transactionModel.Amount, &transactionModel.Date, &transactionModel.IsPending, &transactionModel.IsDone, &transactionModel.IsCanceled); err != nil {
		log.Error("error while scanning user info")
		return err
	}
	if confirmTransaction {
		if transactionModel.IsPending {
			_, err := scr.db.Exec(queryUpdateStatusNoMorePending, transactionID)
			if err != nil {
				return err
			}

			_, err = scr.db.Exec(queryUpdateUserBalancePlus, transactionModel.Amount, transactionModel.ToID) // пополнение баланса получателя
			if err != nil {
				return err
			}

			_, err = scr.db.Exec(queryUpdateStatusDone, transactionID)
			if err != nil {
				return err
			}
			log.Info("отправили денги получателю")

		} else {
			return errors.New("транзакция уже была выполнена")
		}
	} else {
		if transactionModel.IsPending {
			_, err := scr.db.Exec(queryUpdateStatusNoMorePending, transactionID)
			if err != nil {
				return err
			}

			_, err = scr.db.Exec(queryUpdateUserBalancePlus, transactionModel.Amount, transactionModel.FromID) // просто возврат средств отправителю
			if err != nil {
				return err
			}

			_, err = scr.db.Exec(queryUpdateStatusCanceled, transactionID)
			if err != nil {
				return err
			}
			log.Info("вернули деньги отправителю")
		} else {
			return errors.New("транзакция уже была выполнена")
		}
	}

	return nil
}
