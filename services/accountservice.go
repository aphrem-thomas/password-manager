package services

import (
	"github.com/aphrem-thomas/password-manager/aggregates"
	account "github.com/aphrem-thomas/password-manager/domains/accounts"
	"github.com/aphrem-thomas/password-manager/domains/accounts/db"
	"github.com/aphrem-thomas/password-manager/domains/accounts/memory"
	"github.com/google/uuid"
)

type AccountConfiguration func(as *AccountService) error

type AccountService struct {
	accounts account.AccountRepository
}

func NewAccountService(cfgs ...AccountConfiguration) (*AccountService, error) {
	as := &AccountService{}
	for _, cfg := range cfgs {
		err := cfg(as)
		if err != nil {
			return nil, err
		}
	}
	return as, nil
}

func WithAccountRepository(ac account.AccountRepository) AccountConfiguration {
	return func(as *AccountService) error {
		as.accounts = ac
		return nil
	}
}

func WithMemoryAccountRepository() AccountConfiguration {
	ac := memory.New()
	return WithAccountRepository(ac)
}
func WithDBAccountRepository() AccountConfiguration {
	ac := db.New()
	return WithAccountRepository(ac)
}

func (ac *AccountService) GetAccount(id uuid.UUID) (aggregates.Account, error) {
	return ac.accounts.GetAccount(id)
}

func (ac *AccountService) GetAllAccounts() ([]aggregates.Account, error) {
	return ac.accounts.GetAllAccounts()
}

func (ac *AccountService) AddAccount(name string, email string, pwd string, hash string) error {
	acn, _ := aggregates.NewAccount(name, email, pwd, hash)
	return ac.accounts.AddAccount(*acn)
}
