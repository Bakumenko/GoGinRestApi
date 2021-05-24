package repository

import (
	"apiserver/pkg/model"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) GetAllUsers() ([]model.User, error) {
	var users []model.User

	query := fmt.Sprintf(
		"SELECT * FROM %s", viper.GetString("db.usertable"))

	err := u.db.Select(&users, query)

	return users, err
}

func (u *UserPostgres) CreateUser(user model.User) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(
		"INSERT INTO %s (id, firstname, lastname, email, age) values ($1, $2, $3, $4, $5) RETURNING id", viper.GetString("db.usertable"))

	row := u.db.QueryRow(query, user.ID, user.Firstname, user.Lastname, user.Email, user.Age)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (u *UserPostgres) GetOneUser(user_id string) (model.User, error) {
	var user model.User

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE id = $1", viper.GetString("db.usertable"))

	err := u.db.Get(&user, query, user_id)

	return user, err
}

func (u *UserPostgres) UpdateUser(user model.UpdateUserInput, user_id string) (model.User, error) {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Firstname != nil {
		setValues = append(setValues, fmt.Sprintf("firstname=$%d", argId))
		args = append(args, *user.Firstname)
		argId++
	}

	if user.Lastname != nil {
		setValues = append(setValues, fmt.Sprintf("lastname=$%d", argId))
		args = append(args, *user.Lastname)
		argId++
	}

	if user.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *user.Email)
		argId++
	}

	if user.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *user.Age)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d RETURNING id, firstname, lastname, email, age, created_at",
		viper.GetString("db.usertable"), setQuery, argId)
	args = append(args, user_id)

	var updatedUser model.User
	row := u.db.QueryRow(query, args...)
	if err := row.Scan(&updatedUser.ID, &updatedUser.Firstname, &updatedUser.Lastname,
		&updatedUser.Email, &updatedUser.Age, &updatedUser.Created); err != nil {
		return model.User{}, err
	}
	return updatedUser, nil
}

func (u *UserPostgres) DeleteUser(user_id string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", viper.GetString("db.usertable"))
	row, err := u.db.Exec(query, user_id)

	count, err := row.RowsAffected()
	return count, err
}
