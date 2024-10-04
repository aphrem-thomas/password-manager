package db

import (
	"errors"
	"sync"

	"github.com/aphrem-thomas/password-manager/aggregates"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbRepository struct {
	db *gorm.DB
	sync.Mutex
}

type User struct {
	gorm.Model
	Name             string
	Email            string
	PwdHash          string
	Token            string
	PasswordRecordId uint
}

func New() *DbRepository {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&User{})
	return &DbRepository{
		db: db,
	}
}

func (mr *DbRepository) GetAccount(id uuid.UUID) (aggregates.Account, error) {
	// if acc, ok := mr.accounts[id]; ok {
	// 	return acc, nil
	// }
	return aggregates.Account{}, errors.New("unable to find account")
}

func (mr *DbRepository) GetAllAccounts() ([]aggregates.Account, error) {
	var results []aggregates.Account
	// for i := range mr.accounts {
	// 	results = append(results, mr.accounts[i])
	// }
	return results, nil
}

func aggregateToDbUser(aUser aggregates.Account) *User {
	return &User{
		Name:             aUser.GetUser().Name,
		Email:            aUser.GetUser().Email,
		PwdHash:          aUser.GetUser().PwdHash,
		Token:            aUser.GetUser().Token,
		PasswordRecordId: 1,
	}
}

func (mr *DbRepository) AddAccount(ac aggregates.Account) error {
	user := aggregateToDbUser(ac)
	existingUser := mr.db.First(user)
	if existingUser.RowsAffected == 0 {
    
		mr.db.Create(user)
		return nil 
	} else {
		return errors.New("account already exist")
	}
}

func (mr *DbRepository) DeleteAccount(ac aggregates.Account) error {
	return nil
}

func (mr *DbRepository) UpdateAccount(ac aggregates.Account) error {
	return nil
}
