package service

import (
	"finance_manager/internal/models"
	dbactions "finance_manager/internal/repository/db_actions"
)

type BalanceService struct {
	Balance dbactions.BalanceRepository
}

func (b BalanceService) GetBalance(id int) (*models.User, error) {
	userBalance, err := b.Balance.GetUserBalance(id)
	if err != nil {
		return nil, err
	}
	return userBalance, nil
}
