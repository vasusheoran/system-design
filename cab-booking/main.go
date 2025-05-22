package main

import (
	"cab-booking/pkg/contracts"
	"cab-booking/pkg/repo"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	userRepo := repo.NewUser(logger)
	driverRepo := repo.NewDriver()

	userNumbers := []string{}
	driverNumbers := []string{}

	for i := 0; i < 10; i++ {
		dr := contracts.Driver{
			ID:     uuid.NewString(),
			Name:   fmt.Sprintf("test-%d", rand.Intn(10)),
			Number: fmt.Sprintf("01234567%d%d", rand.Intn(10), rand.Intn(10)),
			VehicleDetails: contracts.Vehicle{
				ID:        uuid.NewString(),
				RegNumber: uuid.NewString(),
				Capacity:  rand.Int63n(8),
			},
		}

		driverNumbers = append(driverNumbers, dr.Number)
		driverRepo.SaveDriver(dr)
	}
	for i := 0; i < 10; i++ {
		user := contracts.User{
			ID:     uuid.NewString(),
			Name:   fmt.Sprintf("test-%d", rand.Intn(10)),
			Number: fmt.Sprintf("01234567%d%d", rand.Intn(10), rand.Intn(10)),
		}

		userNumbers = append(userNumbers, user.Number)
		userRepo.SaveUser(user)
	}

	user, err := userRepo.GetUserByNumber(userNumbers[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	driver, err := driverRepo.GetDriverByNumber(driverNumbers[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(driver)

}
