package memory

import (
	"errors"
	"fmt"
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

func (mr *MemoryRepository) GetAllAccounts() ([]aggregates.Account, error) {
	var results []aggregates.Account
	for i := range mr.accounts {
		results = append(results, mr.accounts[i])
	}
	return results, nil
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
		fmt.Println("in mr addaccount", mr.accounts)
		return nil
	}

}

func (mr *MemoryRepository) DeleteAccount(ac aggregates.Account) error {
	return nil
}

func (mr *MemoryRepository) UpdateAccount(ac aggregates.Account) error {
	return nil
}
