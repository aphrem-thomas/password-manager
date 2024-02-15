package memory

import (
	"errors"
	"sync"

	"github.com/aphrem-thomas/password-manager/aggregates"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	accounts map[uuid.UUID]aggregates.Account
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		accounts: make(map[uuid.UUID]aggregates.Account),
	}
}

func (mr *MemoryRepository) GetAccount(id uuid.UUID) (aggregates.Account, error) {
	if acc, ok := mr.accounts[id]; ok {
		return acc, nil
	}
	return aggregates.Account{}, errors.New("unable to find account")
}

func (mr *MemoryRepository) AddAccount(ac aggregates.Account) error {
	if mr.accounts == nil {
		mr.Lock()
		mr.accounts = make(map[uuid.UUID]aggregates.Account)
		mr.Unlock()
	}
	if _, ok := mr.accounts[ac.GetId()]; ok {
		return errors.New("account already exist")
	} else {
		mr.Lock()
		mr.accounts[ac.GetId()] = ac
		mr.Unlock()
		return nil
	}

}

func (mr *MemoryRepository) DeleteAccount(ac aggregates.Account) error {
	return nil
}

func (mr *MemoryRepository) UpdateAccount(ac aggregates.Account) error {
	return nil
}
