package driver

import (
	"cab-booking/contracts"
	"cab-booking/pkg/repo"
	"fmt"
	"time"
)

type Service interface {
	AddDriver(username string, driver contracts.Driver) error
	UpdateLocation(username string, location contracts.Location) error
	UpdateStatus(username string, status contracts.Status) error
	GetNearbyDrivers(location contracts.Location, units float64, size int) []contracts.Driver
	IsValidUser(username string) bool
	GetDriverByUsername(username string) (contracts.Driver, error)
	UpdateHistory(username string, ride contracts.Ride) error
	GetEarningsByUsername(username string, since time.Time) (float64, error)
}

type service struct {
	driverRepo repo.DriverRepo
}

func (s *service) GetEarningsByUsername(username string, since time.Time) (float64, error) {
	driver, err := s.driverRepo.GetDriverByUsername(username)
	if err != nil {
		return 0.0, fmt.Errorf("driver not found")
	}

	var total float64
	for i := len(driver.History) - 1; i >= 0; i-- {
		ride := driver.History[i]
		if ride.CreatedAt.Sub(since) >= 0 {
			total += ride.Price
		} else {
			break
		}
	}
	return total, nil
}

func (s *service) UpdateHistory(username string, ride contracts.Ride) error {
	driver, err := s.driverRepo.GetDriverByUsername(username)
	if err != nil {
		return fmt.Errorf("driver not found")
	}
	driver.History = append(driver.History, ride)
	return s.driverRepo.Save(*driver)
}

func (s *service) GetDriverByUsername(username string) (contracts.Driver, error) {
	dr, err := s.driverRepo.GetDriverByUsername(username)
	if err != nil {
		return contracts.Driver{}, fmt.Errorf("driver not found")
	}

	return *dr, nil
}

func (s *service) IsValidUser(username string) bool {
	if _, err := s.driverRepo.GetDriverByUsername(username); err != nil {
		return false
	}

	return true
}

func (s *service) GetNearbyDrivers(location contracts.Location, units float64, size int) []contracts.Driver {
	drivers := s.driverRepo.GetNearbyDrivers(location, units, size)

	if len(drivers) == 0 {
		return nil
	}

	return drivers
}

func (s *service) UpdateStatus(username string, status contracts.Status) error {
	driver, err := s.driverRepo.GetDriverByUsername(username)
	if err != nil {
		return fmt.Errorf("driver not found")
	}
	driver.Status = status
	return s.driverRepo.Save(*driver)
}

func (s *service) UpdateLocation(username string, location contracts.Location) error {
	driver, err := s.driverRepo.GetDriverByUsername(username)
	if err != nil {
		return fmt.Errorf("driver not found")
	}

	// Validate location
	driver.Location = location
	return s.driverRepo.Save(*driver)
}

func (s *service) AddDriver(username string, driver contracts.Driver) error {
	if _, err := s.driverRepo.GetDriverByUsername(username); err == nil {
		return fmt.Errorf("username already exists")
	}

	return s.driverRepo.Save(driver)
}

func NewDriverService(driverRepo repo.DriverRepo) Service {
	return &service{driverRepo: driverRepo}
}
