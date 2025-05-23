package repo

import (
	"cab-booking/contracts"
	"fmt"
)

type UserRepo interface {
	GetUserByUsername(username string) (*contracts.User, error)
	Save(user contracts.User) error
}

type user struct {
	db map[string]contracts.User
}

func (u *user) Save(user contracts.User) error {
	if len(user.ID) == 0 {
		return fmt.Errorf("user username can not be empty")
	}

	u.db[user.ID] = user
	return nil
}

func (u *user) GetUserByUsername(username string) (*contracts.User, error) {
	if u, ok := u.db[username]; ok {
		return &u, nil
	}

	return nil, fmt.Errorf("user with `%s` username not found")
}

func NewUserRepo() UserRepo {
	return &user{make(map[string]contracts.User)}
}
