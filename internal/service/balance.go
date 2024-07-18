package service

import (
	"finance_manager/internal/models"
	dbactions "finance_manager/internal/repository/db_actions"
)

type BalanceService struct {
	repository dbactions.BalanceRepository
}

func NewBalanceService(repo dbactions.BalanceRepository) *BalanceService {
	return &BalanceService{
		repository: repo,
	}
}

func (bs BalanceService) GetBalance(id int) (*models.User, error) {
	userBalance, err := bs.repository.GetUserBalance(id)
	if err != nil {
		return nil, err
	}
	return userBalance, nil
}
