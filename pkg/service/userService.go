package service

import (
	"apiserver/pkg/model"
	"apiserver/pkg/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user model.User) (uuid.UUID, error) {
	user.ID = uuid.New()
	return u.repo.CreateUser(user)
}