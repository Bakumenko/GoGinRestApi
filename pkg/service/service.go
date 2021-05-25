package service

import (
	"apiserver/pkg/model"
	"apiserver/pkg/repository"
	"github.com/google/uuid"
)

type User interface {
	GetAllUsers() ([]model.UserOutput, error)
	CreateUser(user model.User, roleName string) (uuid.UUID, error)
	GetOneUser(user_id string) (model.UserOutput, error)
	UpdateUser(user model.UpdateUserInput, user_id string) (model.UserOutput, error)
	DeleteUser(user_id string) (int64, error)
}

type Role interface {
	CreateRole(role model.Role) (int64, error)
}

type Service struct {
	User
	Role
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User, repos.Role),
		Role: NewRoleService(repos.Role),
	}
}
