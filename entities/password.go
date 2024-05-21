package entities

import (
	"github.com/google/uuid"
)

type Password struct {
	Id       uuid.UUID
	RecordId uuid.UUID
	Site     string
	Email    string
	Name     string
	PwdHash  string
}
