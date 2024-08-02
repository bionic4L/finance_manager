package service

import (
	"context"
	"finance_manager/internal/repository"
)

type DepositService struct {
	repository repository.DepositRepository
}

func NewDepositService(repository repository.DepositRepository) *DepositService {
	return &DepositService{repository: repository}
}

func (ds *DepositService) Deposit(ctx context.Context, id int, amount int) error {
	err := ds.repository.Deposit(ctx, id, amount)
	if err != nil {
		return err
	}
	return nil
}
