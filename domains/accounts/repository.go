package account

import (
	"github.com/aphrem-thomas/password-manager/aggregates"
	"github.com/google/uuid"
)

type AccountRepository interface {
	GetAccount(uuid.UUID) (aggregates.Account, error)
	AddAccount(aggregates.Account) error
	DeleteAccount(aggregates.Account) error
	UpdateAccount(aggregates.Account) error
}
