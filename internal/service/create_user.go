package service

import (
	"context"
	"finance_manager/internal/repository"
)

type CreateUserService struct {
	repository repository.UserCreateRepository
}

func NewCreateUserService(repository repository.UserCreateRepository) *CreateUserService {
	return &CreateUserService{repository: repository}
}

func (cus *CreateUserService) UserCreate(ctx context.Context, name string) error {
	if err := cus.repository.UserCreate(ctx, name); err != nil {
		return err
	}

	return nil
}
