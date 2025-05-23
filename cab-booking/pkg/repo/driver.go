package repo

import (
	"cab-booking/contracts"
	"fmt"
	"math"
)

type DriverRepo interface {
	GetDriverByUsername(username string) (*contracts.Driver, error)
	Save(driver contracts.Driver) error
	GetNearbyDrivers(location contracts.Location, units float64, size int) []contracts.Driver
}

type driver struct {
	db map[string]contracts.Driver
}

func (d *driver) GetNearbyDrivers(location contracts.Location, units float64, size int) []contracts.Driver {
	var result []contracts.Driver
	for _, dr := range d.db {
		if GetDistance(location, dr.Location) <= units {
			result = append(result, dr)
		}
	}

	if len(result) >= size {
		return result[:size]
	}

	return result
}

func (d *driver) Save(driver contracts.Driver) error {
	if len(driver.ID) == 0 {
		return fmt.Errorf("driver username can not be empty")
	}

	d.db[driver.ID] = driver
	return nil
}

func (d *driver) GetDriverByUsername(username string) (*contracts.Driver, error) {
	if u, ok := d.db[username]; ok {
		return &u, nil
	}

	return nil, fmt.Errorf("driver with `%s` username not found")
}

func NewDriverRepo() DriverRepo {
	return &driver{make(map[string]contracts.Driver)}
}

func GetDistance(source, destination contracts.Location) float64 {
	xdiff := source.X - destination.X
	ydiff := source.Y - destination.Y
	return math.Abs(float64(xdiff)) + math.Abs(float64(ydiff))
}
