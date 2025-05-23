package user

import (
	"cab-booking/contracts"
	"cab-booking/pkg/repo"
	"fmt"
)

type Service interface {
	AddUser(username string, user contracts.User) error
	UpdateUser(user contracts.User) error
	UpdateLocation(username string, location contracts.Location) error
	IsValidUser(username string) bool
	GetUserByUsername(username string) (*contracts.User, error)
}

type service struct {
	userRepo repo.UserRepo
}

func (s *service) GetUserByUsername(username string) (*contracts.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (s *service) IsValidUser(username string) bool {
	if _, err := s.userRepo.GetUserByUsername(username); err != nil {
		return false
	}

	return true
}

func (s *service) UpdateLocation(username string, location contracts.Location) error {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Validate location
	user.Location = location
	return s.userRepo.Save(*user)
}

func (s *service) UpdateUser(user contracts.User) error {
	if _, err := s.userRepo.GetUserByUsername(user.ID); err != nil {
		return fmt.Errorf("user not found")
	}
	return s.userRepo.Save(user)
}

func (s *service) AddUser(userName string, user contracts.User) error {
	if _, err := s.userRepo.GetUserByUsername(userName); err == nil {
		return fmt.Errorf("username already exists")
	}

	return s.userRepo.Save(user)
}

func NewUserService(userRepo repo.UserRepo) Service {
	return &service{userRepo: userRepo}
}
