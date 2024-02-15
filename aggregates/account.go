package aggregates

import (
	"github.com/aphrem-thomas/password-manager/entities"
	"github.com/google/uuid"
)

type Account struct {
	user *entities.User
	data []*entities.Passwords
}

// NewAccount is a factory function
func NewAccount(name string, email string, pwdHash string, token string) (*Account, error) {
	userInfo := &entities.User{
		Name:    name,
		Id:      uuid.New(),
		Email:   email,
		PwdHash: pwdHash,
		Token:   token,
	}
	var userPassInfo []*entities.Passwords
	userPassInfo = append(userPassInfo, &entities.Passwords{
		Id:      uuid.New(),
		Site:    "",
		Email:   "",
		Name:    "",
		PwdHash: "",
	})
	return &Account{
		user: userInfo,
		data: userPassInfo,
	}, nil
}

func (ac *Account) GetId() uuid.UUID {
	return ac.user.Id
}

func (ac *Account) SetId(id uuid.UUID) {
	if ac.user == nil {
		ac.user = &entities.User{}
	}
	ac.user.Id = id
}
