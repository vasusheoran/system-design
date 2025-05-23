package main

import (
	"cab-booking/contracts"
	"cab-booking/pkg/driver"
	"cab-booking/pkg/repo"
	"cab-booking/pkg/ride"
	"cab-booking/pkg/user"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	userRepo := repo.NewUserRepo()
	userService := user.NewUserService(userRepo)

	driverRepo := repo.NewDriverRepo()
	driverService := driver.NewDriverService(driverRepo)

	rideService := ride.NewRideService(driverService, userService)

	user := contracts.User{
		ID:   uuid.NewString(), // some username
		Name: "Abc",
		Age:  20,
		Location: contracts.Location{
			X: 0,
			Y: 0,
		},
	}

	err := userService.AddUser(uuid.NewString(), user)
	if err != nil {
		panic(err)
	}

	user.Age = 21
	err = userService.UpdateUser(user)
	if err != nil {
		panic(err)
	}

	err = userService.UpdateLocation(user.ID, contracts.Location{5, 5})
	if err != nil {
		panic(err)
	}

	driver := contracts.Driver{
		ID:   uuid.NewString(), // some username
		Name: "Abc",
		Age:  20,
		Location: contracts.Location{
			X: 0,
			Y: 0,
		},
		Status: contracts.Available,
		Vehicle: contracts.Vehicle{
			ID:       uuid.NewString(),
			Type:     contracts.Sedan,
			Name:     "Swift",
			Capacity: 5,
		},
	}
	err = driverService.AddDriver(driver.ID, driver)
	if err != nil {
		panic(err)
	}

	err = driverService.UpdateLocation(driver.ID, contracts.Location{5, 5})
	if err != nil {
		panic(err)
	}

	err = driverService.UpdateStatus(driver.ID, contracts.Busy)
	if err != nil {
		panic(err)
	}

	_, err = rideService.Search(user.ID, user.Location)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = driverService.UpdateLocation(driver.ID, contracts.Location{1, 1})
	if err != nil {
		panic(err)
	}

	driverList, err := rideService.Search(user.ID, user.Location)
	if err != nil {
		fmt.Println(err.Error())
	}

	if driverList != nil && len(driverList) > 0 {
		fmt.Println(len(driverList))
	}

	ride, err := rideService.BookRide(user.ID, driverList[0].ID, user.Location, contracts.Location{10, 10})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ride)

	ride, err = rideService.BookRide(user.ID, driverList[0].ID, user.Location, contracts.Location{18, 10})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ride)

	earnings, err := rideService.Earnings(driverList[0].ID, time.Now().Add(-(time.Minute * 2)))
	fmt.Printf("Total earnings for driver `%s` is `%.2f`\n", driverList[0].ID, earnings)
}
