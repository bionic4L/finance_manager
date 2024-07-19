package dbactions

import (
	"finance_manager/internal/models"

	"github.com/jmoiron/sqlx"
)

type BalanceRepository struct {
	db *sqlx.DB
}

func NewBalanceRepository(db *sqlx.DB) *BalanceRepository {
	return &BalanceRepository{db: db}
}

func (br *BalanceRepository) GetUserBalance(id int) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := br.db.QueryRow(query, id)

	var user models.User

	if err := row.Scan(&user.ID, &user.Name, &user.Balance); err != nil {
		return nil, err
	}

	return &user, nil
}
