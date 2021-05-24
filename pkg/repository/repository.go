package repository

import (
	"apiserver/pkg/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(user model.User) (uuid.UUID, error)
	GetOneUser(user_id string) (model.User, error)
	UpdateUser(user model.UpdateUserInput, user_id string) (model.User, error)
	DeleteUser(user_id string) (int64, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
