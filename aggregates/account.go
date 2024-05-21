package aggregates

import (
	"github.com/aphrem-thomas/password-manager/entities"
	"github.com/google/uuid"
)

type Account struct {
	user *entities.User
}

// NewAccount is a factory function
func NewAccount(name string, email string, pwdHash string, token string) (*Account, error) {
	recordId := uuid.New()
	userInfo := &entities.User{
		Name:             name,
		Id:               uuid.New(),
		Email:            email,
		PwdHash:          pwdHash,
		Token:            token,
		PasswordRecordId: recordId,
	}

	return &Account{
		user: userInfo,
	}, nil
}

func (ac *Account) GetId() uuid.UUID {
	return ac.user.Id
}

func (ac *Account) GetUser() entities.User {
	return *ac.user
}

func (ac *Account) SetId(id uuid.UUID) {
	if ac.user == nil {
		ac.user = &entities.User{}
	}
	ac.user.Id = id
}
