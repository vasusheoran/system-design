package ride

import (
	"cab-booking/contracts"
	"cab-booking/pkg/driver"
	"cab-booking/pkg/repo"
	"cab-booking/pkg/user"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const (
	SearchDriverLimit = 5
	SearchUnits       = 5.0
	PricePerUnit      = 100.0
)

type SearchResponse struct {
	// Limit view here. Use response wrapper.
	Driver   contracts.Driver
	Distance int
}

type Service interface {
	// Search returns list of drivers nearby user location by `n` units
	// TODO: create a wrapper for SearchResponse instead of returning `contracts.Driver`
	Search(userID string, location contracts.Location) ([]contracts.Driver, error)
	BookRide(userID string, driverID string, source, destination contracts.Location) (contracts.Ride, error)
	CalculateBill(source contracts.Location, destination contracts.Location) float64
	Earnings(driverID string, since time.Time) (float64, error)
}

type service struct {
	driverService driver.Service
	userService   user.Service
}

func (s *service) Earnings(driverID string, since time.Time) (float64, error) {
	return s.driverService.GetEarningsByUsername(driverID, since)
}

func (s *service) BookRide(userID string, driverID string, source, destination contracts.Location) (contracts.Ride, error) {
	if !s.driverService.IsValidUser(driverID) {
		return contracts.Ride{}, fmt.Errorf("invalid driver")
	}

	user, err := s.userService.GetUserByUsername(userID)
	if err != nil {
		return contracts.Ride{}, err
	}

	user.Location = source
	user.Ride = contracts.Ride{
		ID:        uuid.NewString(),
		UserID:    userID,
		DriverID:  driverID,
		Source:    source,
		Dest:      destination,
		Price:     s.CalculateBill(source, destination),
		CreatedAt: time.Now(),
	}

	err = s.userService.UpdateUser(*user)
	if err != nil {
		return contracts.Ride{}, err
	}

	err = s.driverService.UpdateHistory(driverID, user.Ride)
	if err != nil {
		// Handle driver history updates async
		return user.Ride, nil
	}

	return user.Ride, nil
}

func (s *service) Search(userID string, location contracts.Location) ([]contracts.Driver, error) {
	if !s.userService.IsValidUser(userID) {
		return nil, fmt.Errorf("user `%s` not found or is a invalid user", userID)
	}

	drivers := s.getNearbyDriversWithinUnits(location, SearchUnits)
	if drivers == nil {
		return nil, fmt.Errorf("no nearby drivers found for location `%s` within `%.0f` units", location, SearchUnits)
	}

	return drivers, nil
}

func (s *service) getNearbyDriversWithinUnits(location contracts.Location, units float64) []contracts.Driver {
	return s.driverService.GetNearbyDrivers(location, units, SearchDriverLimit)
}

func (s *service) CalculateBill(source contracts.Location, destination contracts.Location) float64 {
	return repo.GetDistance(source, destination) * PricePerUnit
}

func NewRideService(driverService driver.Service, userService user.Service) Service {
	return &service{
		driverService: driverService,
		userService:   userService,
	}
}
