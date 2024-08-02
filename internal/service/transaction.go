package service

import (
	"context"
	"finance_manager/internal/repository"
)

type TransactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) *TransactionService {
	return &TransactionService{repository: repository}
}

func (ts *TransactionService) Transaction(ctx context.Context, fromID int, toID int, amount int) error {
	if err := ts.repository.Transaction(ctx, fromID, toID, amount); err != nil {
		return err
	}
	return nil
}
