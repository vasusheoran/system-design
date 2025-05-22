package repo

import (
	"cab-booking/pkg/contracts"
	"fmt"
)

type Driver interface {
	GetDriverByID(string) (*contracts.Driver, error)
	GetDriverByNumber(string) (*contracts.Driver, error)
	UpdateLocation(contracts.Location) error
	SaveDriver(driver contracts.Driver) error
}

type driverRepo struct {
	db            map[string]contracts.Driver
	numberToIDMap map[string]string
}

func (d *driverRepo) UpdateLocation(location contracts.Location) error {
	//TODO implement me
	panic("implement me")
}

func (d *driverRepo) GetDriverByNumber(number string) (*contracts.Driver, error) {
	id, ok := d.numberToIDMap[number]
	if !ok {
		return nil, fmt.Errorf("driver with number %s not found", number)
	}

	return d.GetDriverByID(id)
}

func (d *driverRepo) GetDriverByID(id string) (*contracts.Driver, error) {
	driver, ok := d.db[id]
	if !ok {
		return nil, fmt.Errorf("driver with id %s not found", id)
	}

	return &driver, nil
}

func (d *driverRepo) SaveDriver(driver contracts.Driver) error {
	d.db[driver.ID] = driver
	d.numberToIDMap[driver.Number] = driver.ID
	return nil
}

func NewDriver() Driver {
	return &driverRepo{
		db:            map[string]contracts.Driver{},
		numberToIDMap: map[string]string{},
	}
}
