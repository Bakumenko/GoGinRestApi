package service

import (
	"apiserver/pkg/model"
	"apiserver/pkg/repository"
	"github.com/google/uuid"
)

type User interface {
	CreateUser(user model.User) (uuid.UUID, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}