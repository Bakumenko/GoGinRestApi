package repository

import (
	"apiserver/pkg/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CreateUser(user model.User) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(
		"INSERT INTO %s (id, firstname, lastname, email, age) values ($1, $2, $3, $4, $5) RETURNING id", viper.GetString("db.usertable"))

	row := u.db.QueryRow(query, user.ID, user.Firstname, user.Lastname, user.Email, user.Age)
	if err :=row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}