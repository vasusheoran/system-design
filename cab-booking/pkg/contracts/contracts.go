package contracts

type Type string

const (
	UserType   Type = "user"
	DriverType Type = "driver"
)

type Status string

const (
	Unregistered Status = "unregistered"
	Booking      Status = "booking"
	Booked       Status = "booked"
	Inactive     Status = "inactive"
)

type User struct {
	ID     string `json:"-"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Status Status `json:"status"`
	Ride   Ride   `json:"ride"`
}

type Driver struct {
	ID             string   `json:"-"`
	Name           string   `json:"name"`
	Number         string   `json:"number"`
	Subscribe      bool     `json:"subscribe"`
	License        string   `json:"license"`
	Location       Location `json:"location"`
	VehicleDetails Vehicle  `json:"vehicleDetails"`
}

type Vehicle struct {
	ID        string `json:"-"`
	RegNumber string `json:"regNumber"`
	Capacity  int64  `json:"capacity"`
}

type Location string

type Ride struct {
	ID          string   `json:"-"`
	UserID      string   `json:"-"`
	DriverID    string   `json:"-"`
	Source      Location `json:"source"`
	Destination Location `json:"destination"`
	Price       float64  `json:"-"`
	//Route       []Location   `json:"route"`, SurgeMultiplier, Discount etc
}
