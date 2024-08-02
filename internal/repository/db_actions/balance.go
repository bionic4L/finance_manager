package dbactions

import (
	"context"
	"finance_manager/internal/models"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type BalanceRepository struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) *BalanceRepository {
	return &BalanceRepository{db: db}
}

func (br *BalanceRepository) GetUserBalance(ctx context.Context, id int) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := br.db.QueryRowContext(ctx, query, id)

	var user models.User

	if err := row.Scan(&user.ID, &user.Name, &user.Balance); err != nil {
		log.Error("error while scanning user info")
		return nil, err
	}

	return &user, nil
}
