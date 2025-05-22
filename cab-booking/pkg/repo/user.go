package repo

import (
	"cab-booking/pkg/contracts"
	"fmt"
	"log"
)

type User interface {
	GetUserByID(string) (*contracts.User, error)
	GetUserByNumber(string) (*contracts.User, error)
	SaveUser(contracts.User) error
}

type userRepo struct {
	logger        *log.Logger
	db            map[string]contracts.User
	numberToIDMap map[string]string
}

func (u *userRepo) GetUserByNumber(number string) (*contracts.User, error) {
	id, ok := u.numberToIDMap[number]
	if !ok {
		return nil, fmt.Errorf("user with number %s not found", number)
	}

	return u.GetUserByID(id)
}

// GetUserByID return err if no user found with this ID
func (u *userRepo) GetUserByID(id string) (*contracts.User, error) {
	user, ok := u.db[id]
	if !ok {
		return nil, fmt.Errorf("user with id %s not found", id)
	}

	return &user, nil
}

func (u *userRepo) SaveUser(user contracts.User) error {
	//if _, ok := u.db[user.ID]; ok {
	//	return fmt.Errorf("user with id %s already exists", user.ID)
	//}

	u.db[user.ID] = user
	u.numberToIDMap[user.Number] = user.ID
	return nil
}

func NewUser(logger *log.Logger) User {
	return &userRepo{
		logger:        logger,
		db:            make(map[string]contracts.User),
		numberToIDMap: make(map[string]string),
	}
}
