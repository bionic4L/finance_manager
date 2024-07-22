package service

import "finance_manager/internal/repository"

type CreateUserService struct {
	repository repository.UserCreateRepository
}

func NewCreateUserService(repository repository.UserCreateRepository) *CreateUserService {
	return &CreateUserService{repository: repository}
}

func (cus CreateUserService) CreateUser(name string) error {
	if err := cus.CreateUser(name); err != nil {
		return err
	}

	return nil
}
