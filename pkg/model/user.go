package model

import (
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
