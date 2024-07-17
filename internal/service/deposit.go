package service

import (
	dbactions "finance_manager/internal/repository/db_actions"
)

type DepositService struct {
	DepositRepository dbactions.DepositRepository
}

func (ds *DepositService) Deposit(id int, amount int) error {
	err := ds.DepositRepository.DepositToUser(id, amount)
	if err != nil {
		return err
	}
	return nil
}
