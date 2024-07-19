package service

import "finance_manager/internal/repository"

type DepositService struct {
	repository repository.DepositRepository
}

func NewDepositService(repository repository.DepositRepository) *DepositService {
	return &DepositService{repository: repository}
}

func (ds *DepositService) Deposit(id int, amount int) error {
	err := ds.repository.Deposit(id, amount)
	if err != nil {
		return err
	}
	return nil
}
