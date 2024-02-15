package services

import (
	"github.com/aphrem-thomas/password-manager/aggregates"
	account "github.com/aphrem-thomas/password-manager/domains/accounts"
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

func (ac *AccountService) GetAccount(id uuid.UUID) (aggregates.Account, error) {
	return ac.accounts.GetAccount(id)
}

func (ac *AccountService) AddAccount() error {
	acn, _ := aggregates.NewAccount("Aphrem Thomas", "aphrem@thomasaphrem.com", "fdadsfafr34ewr34re", "adsfasdfcoijermlkdfj234324")
	return ac.accounts.AddAccount(*acn)
}
