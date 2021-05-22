package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname" binding:"required"`
	Lastname  string    `json:"lastname" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Age       uint      `json:"age" binding:"required"`
	Created   time.Time `json:"created_at"`
}
