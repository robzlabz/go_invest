package service

import (
	"go_invest/app/model"
	"go_invest/app/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}
