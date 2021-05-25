package repository

import (
	"apiserver/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type RolePostgres struct {
	db *sqlx.DB
}

func NewRolePostgres(db *sqlx.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) GetOneRole(name string) (model.Role, error) {
	var role model.Role

	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE name = $1", viper.GetString("db.roletable"))

	err := r.db.Get(&role, query, name)

	return role, err
}

func (r *RolePostgres) CreateRole(role model.Role) (int64, error) {
	var id int64

	query := fmt.Sprintf(
		"INSERT INTO %s name value $1 RETURNING id", viper.GetString("db.roletable"))

	row := r.db.QueryRow(query, role.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
