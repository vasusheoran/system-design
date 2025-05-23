package contracts

type Status int

const (
	Available Status = iota
	Busy
)

type Driver struct {
	ID       string   `json:"username"`
	Name     string   `json:"name"`
	Number   int      `json:"number"`
	Age      int      `json:"age"`
	Vehicle  Vehicle  `json:"vehicle"`
	Location Location `json:"location"`
	Status   Status   `json:"status"`
	History  []Ride   `json:"history"`
}
