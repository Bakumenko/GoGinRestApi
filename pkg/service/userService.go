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

func (u *UserService) GetAllUsers() ([]model.User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserService) CreateUser(user model.User) (uuid.UUID, error) {
	user.ID = uuid.New()
	return u.repo.CreateUser(user)
}

func (u *UserService) GetOneUser(user_id string) (model.User, error) {
	return u.repo.GetOneUser(user_id)
}

func (u *UserService) UpdateUser(user model.UpdateUserInput, user_id string) (model.User, error) {
	if err := user.Validate(); err != nil {
		return model.User{}, err
	}
	return u.repo.UpdateUser(user, user_id)
}

func (u *UserService) DeleteUser(user_id string) (int64, error) {
	return u.repo.DeleteUser(user_id)
}
