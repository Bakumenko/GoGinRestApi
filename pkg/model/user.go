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
	Created   time.Time `json:"created_at" db:"created_at"`
}

type UpdateUserInput struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	Age       *uint   `json:"age"`
}

func (u UpdateUserInput) Validate() error {
	if u.Firstname == nil && u.Lastname == nil && u.Age == nil {
		errors.New("no values for update")
	}
	return nil
}
