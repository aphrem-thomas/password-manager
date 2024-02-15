package entities

import "github.com/google/uuid"

type User struct {
	Id      uuid.UUID
	Email   string
	Name    string
	PwdHash string
	Token   string
}
