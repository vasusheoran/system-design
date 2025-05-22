package registration

import (
	"cab-booking/pkg/contracts"
	"cab-booking/pkg/repo"
	"fmt"
	"github.com/google/uuid"
)

type Registration interface {
	SaveUser(name, number string) error
	SaveDriver(name, number string, vehicle contracts.Vehicle) error
	GetUserDetails(string) (*contracts.User, error)
	GetDriverDetails(string) (*contracts.Driver, error)
}

type registration struct {
	driverRepo repo.Driver
	userRepo   repo.User
}

func (r registration) SaveUser(name, number string) error {
	if user, err := r.userRepo.GetUserByNumber(number); user != nil || err == nil {
		return fmt.Errorf("user with number %s already exists", number)
	}

	user := contracts.User{
		ID:     uuid.NewString(),
		Name:   name,
		Number: number,
		Status: contracts.Inactive,
		Ride:   contracts.Ride{},
	}

	return r.userRepo.SaveUser(user)
}

func (r registration) SaveDriver(name, number string, vehicle contracts.Vehicle) error {
	if driver, err := r.driverRepo.GetDriverByNumber(number); driver != nil || err == nil {
		return fmt.Errorf("driver with number %s already exists", number)
	}

	driver := contracts.Driver{
		ID:             uuid.NewString(),
		Name:           name,
		Number:         number,
		License:        "",
		Location:       "",
		VehicleDetails: vehicle,
	}

	return r.driverRepo.SaveDriver(driver)
}

func (r registration) GetUserDetails(number string) (*contracts.User, error) {
	return r.userRepo.GetUserByNumber(number)
}

func (r registration) GetDriverDetails(number string) (*contracts.Driver, error) {
	return r.driverRepo.GetDriverByNumber(number)
}

func New(driverRepo repo.Driver, userRepo repo.User) Registration {
	return &registration{
		driverRepo: driverRepo,
		userRepo:   userRepo,
	}
}
