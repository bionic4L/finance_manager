package service

import (
	"context"
	"finance_manager/internal/repository"
)

type StatusControllerService struct {
	repository repository.StatusControllerRepository
}

func NewStatusControllerService(repository repository.StatusControllerRepository) *StatusControllerService {
	return &StatusControllerService{
		repository: repository,
	}
}

func (scs StatusControllerService) StatusController(ctx context.Context, id int, confirm bool) error {
	err := scs.repository.StatusController(ctx, id, confirm)
	if err != nil {
		return err
	}
	return nil
}
