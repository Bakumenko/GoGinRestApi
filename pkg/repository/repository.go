package repository

import (
	"apiserver/pkg/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetAllUsers() ([]model.UserOutput, error)
	CreateUser(user model.User) (uuid.UUID, error)
	GetOneUser(user_id string) (model.UserOutput, error)
	CheckEmailContatinsInDB(emaid string) (bool, error)
	UpdateUser(user model.UpdateUserInput, user_id string) (model.UserOutput, error)
	DeleteUser(user_id string) (int64, error)
}

type Role interface {
	GetOneRole(name string) (model.Role, error)
	CreateRole(role model.Role) (int64, error)
}

type Repository struct {
	User
	Role
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
		Role: NewRolePostgres(db),
	}
}
