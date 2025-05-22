package rides

import (
	"cab-booking/pkg/contracts"
	"cab-booking/pkg/repo"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
)

type Rides interface {
	Create(string, contracts.Location, contracts.Location) (string, error)
	// UserID
	End(string, contracts.Location) (string, error)
	GetRideByID(string) (contracts.Ride, error)
	Subscribe(driverID string) error
}

type rides struct {
	userRepo   repo.User
	driverRepo repo.Driver
	rideMap    map[string]contracts.Ride
}

func (r *rides) End(userID string, location contracts.Location) (string, error) {

}

func (r *rides) Subscribe(driverID string) error {
	//TODO implement me
	driver, err := r.driverRepo.GetDriverByID(driverID)
	if err != nil {
		return err
	}

	driver.Subscribe = true
	return r.driverRepo.SaveDriver(*driver)
}

func (r *rides) GetRideByID(s string) (contracts.Ride, error) {
	//TODO implement me
	panic("implement me")
}

func (r *rides) getPrice(source, destination contracts.Location) float64 {
	return rand.Float64() + 100.0
}

// Only get valid drivers
func (r *rides) getDriversByLocation(location contracts.Location) ([]contracts.Driver, error) {
	var drivers []contracts.Driver

	for i := 0; i < 10; i++ {
		dr := contracts.Driver{
			ID:       uuid.NewString(),
			Name:     fmt.Sprintf("test-%d", rand.Intn(10)),
			Location: location,
			VehicleDetails: contracts.Vehicle{
				ID:        uuid.NewString(),
				RegNumber: uuid.NewString(),
				Capacity:  rand.Int63n(8),
			},
		}
		drivers = append(drivers, dr)
	}
	return drivers, nil
}

func (r *rides) Create(userID string, source, destination contracts.Location) (string, error) {
	user, err := r.userRepo.GetUserByID(userID)
	if err != nil {
		return "", err
	}

	user.Ride = contracts.Ride{
		ID:          uuid.NewString(),
		UserID:      userID,
		Source:      source,
		Destination: destination,
		Price:       r.getPrice(source, destination),
	}
	drivers, _ := r.getDriversByLocation(source)
	user.Ride.DriverID = drivers[0].ID

	err = r.userRepo.SaveUser(*user)
	if err != nil {
		return "", err
	}

	r.rideMap[user.Ride.ID] = user.Ride

	return user.ID, nil
}

func NewRides(user repo.User, driver repo.Driver) Rides {
	// Iterate users and populate all rides
	return &rides{
		userRepo:   user,
		driverRepo: driver,
		rideMap:    make(map[string]contracts.Ride),
	}
}
