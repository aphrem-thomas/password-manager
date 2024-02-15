package entities

import (
	"github.com/google/uuid"
)

type Passwords struct {
	Id      uuid.UUID
	Site    string
	Email   string
	Name    string
	PwdHash string
}
