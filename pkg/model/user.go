package model

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Firstname string    `json:"firstname" db:"firstname" binding:"required"`
	Lastname  string    `json:"lastname" db:"lastname" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Age       uint      `json:"age" db:"age" binding:"required"`
	Role_id   int64     `db:"role_id"`
	Created   time.Time `json:"created_at" db:"created_at"`
}

type UpdateUserInput struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	Age       *uint   `json:"age"`
}

type UserOutput struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	RoleName  string    `json:"role_name" db:"name"`
	Created   time.Time `json:"created_at" db:"created_at"`
}

func (u UpdateUserInput) Validate() error {
	if u.Firstname == nil && u.Lastname == nil && u.Age == nil {
		errors.New("no values for update")
	}
	return nil
}
