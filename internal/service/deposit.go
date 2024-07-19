package service

import (
	dbactions "finance_manager/internal/repository/db_actions"
)

type DepositService struct {
	repository dbactions.DepositRepository
}

func (ds *DepositService) Deposit(id int, amount int) error {
	err := ds.repository.Deposit(id, amount)
	if err != nil {
		return err
	}
	return nil
}
