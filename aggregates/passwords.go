package aggregates

import (
	"github.com/aphrem-thomas/password-manager/entities"
	"github.com/google/uuid"
)

type Passwords struct {
	data *[]*entities.Password
}

func NewPassword(recordId uuid.UUID, email string, pwdHash string, site string) (*Passwords, error) {
	id := uuid.New()
	var pwdInfo []*entities.Password
	pwd := &entities.Password{
		Id:       id,
		RecordId: uuid.New(),
		Site:     email,
		Email:    pwdHash,
		PwdHash:  "token",
	}
	pwdInfo = append(pwdInfo, pwd)

	return &Passwords{
		data: &pwdInfo,
	}, nil
}
