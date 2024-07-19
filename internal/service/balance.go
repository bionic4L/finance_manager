package service

import (
	"finance_manager/internal/models"
	"finance_manager/internal/repository"
)

type BalanceService struct {
	repository repository.BalanceRepository
}

func NewBalanceService(repo repository.BalanceRepository) *BalanceService {
	return &BalanceService{
		repository: repo,
	}
}

func (bs BalanceService) GetBalance(id int) (*models.User, error) {
	userBalance, err := bs.repository.GetUserBalance(id) //прокид с сервисного слоя на слой репозитория
	if err != nil {
		return nil, err
	}
	return userBalance, nil
}
