package service

import (
	"apiserver/pkg/model"
	"apiserver/pkg/repository"
	"github.com/google/uuid"
)

type User interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) (uuid.UUID, error)
	GetOneUser(user_id string) (model.User, error)
	UpdateUser(user model.UpdateUserInput, user_id string) (model.User, error)
	DeleteUser(user_id string) (int64, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
