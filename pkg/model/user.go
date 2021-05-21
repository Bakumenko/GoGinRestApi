package model

import (
	"github.com/google/uuid"
	"time"
)

type myUUID uuid.UUID

type User struct {
	ID        myUUID    `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Age       uint      `json:"age"`
	Created   time.Time `json:"created"`
}
