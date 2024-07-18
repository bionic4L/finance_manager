package dbactions

import (
	"finance_manager/internal/models"

	"github.com/jmoiron/sqlx"
)

type BalanceRepository struct {
	DB *sqlx.DB
}

func NewBalanceRepository(DB *sqlx.DB) *BalanceRepository {
	return &BalanceRepository{DB: DB}
}

func (br *BalanceRepository) GetUserBalance(id int) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := br.DB.QueryRow(query, id)

	var user models.User

	if err := row.Scan(&user.ID, &user.Name, &user.Balance); err != nil {
		return nil, err
	}

	return &user, nil
}
