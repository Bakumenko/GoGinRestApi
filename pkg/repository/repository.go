package repository

import (
	"apiserver/pkg/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user model.User) (uuid.UUID, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
