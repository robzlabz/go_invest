package repository

import (
	"go_invest/app/model"
	"go_invest/database"
)

type UserRepository struct{}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func (r *UserRepository) Create(user *model.User) error {
	result := database.DB.Create(user)
	return result.Error
}
