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

func (u *UserPostgres) GetAllUsers() ([]model.UserOutput, error) {
	var users []model.UserOutput

	query := fmt.Sprintf(
		"SELECT u.id, u.firstname, u.lastname, u.email, u.age, r.name, u.created_at "+
			"FROM %s as u "+
			"LEFT JOIN %s as r "+
			"ON u.role_id = r.id",
		viper.GetString("db.usertable"), viper.GetString("db.roletable"))

	err := u.db.Select(&users, query)

	return users, err
}

func (u *UserPostgres) CreateUser(user model.User) (uuid.UUID, error) {
	var id uuid.UUID

	query := fmt.Sprintf(
		"INSERT INTO %s (id, firstname, lastname, email, age, role_id) "+
			"VALUES ($1, $2, $3, $4, $5, $6) "+
			"RETURNING id",
		viper.GetString("db.usertable"))

	row := u.db.QueryRow(query, user.ID, user.Firstname, user.Lastname, user.Email, user.Age, user.Role_id)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (u *UserPostgres) GetOneUser(user_id string) (model.UserOutput, error) {
	var user model.UserOutput

	query := fmt.Sprintf(
		"SELECT u.id, u.firstname, u.lastname, u.email, u.age, r.name, u.created_at "+
			"FROM %s as u "+
			"LEFT JOIN %s as r "+
			"ON u.role_id = r.id "+
			"WHERE u.id = $1",
		viper.GetString("db.usertable"),
		viper.GetString("db.roletable"))

	err := u.db.Get(&user, query, user_id)

	return user, err
}

func (u *UserPostgres) CheckEmailContatinsInDB(email string) (bool, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE email = $1", viper.GetString("db.usertable"))

	row, err := u.db.Exec(query, email)
	if err != nil {
		return true, err
	}

	if count, err := row.RowsAffected(); err != nil || count != 0 {
		return true, err
	} else {
		return false, err
	}
}

func (u *UserPostgres) UpdateUser(user model.UpdateUserInput, user_id string) (model.UserOutput, error) {
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

	query := fmt.Sprintf(
		"UPDATE %s as u "+
			"SET %s "+
			"FROM %s as r "+
			"WHERE u.role_id = r.id "+
			"AND u.id = $%d "+
			"RETURNING u.id, u.firstname, u.lastname, u.email, u.age, r.name, u.created_at",
		viper.GetString("db.usertable"),
		setQuery,
		viper.GetString("db.roletable"),
		argId)
	args = append(args, user_id)

	var updatedUser model.UserOutput
	row := u.db.QueryRow(query, args...)
	if err := row.Scan(&updatedUser.ID, &updatedUser.Firstname, &updatedUser.Lastname,
		&updatedUser.Email, &updatedUser.Age, &updatedUser.RoleName, &updatedUser.Created); err != nil {
		return model.UserOutput{}, err
	}
	return updatedUser, nil
}

func (u *UserPostgres) DeleteUser(user_id string) (int64, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", viper.GetString("db.usertable"))
	row, err := u.db.Exec(query, user_id)

	count, err := row.RowsAffected()
	return count, err
}
