package contracts

type VehicleType int

const (
	Hatchback VehicleType = iota
	Sedan
)

type Vehicle struct {
	ID       string      `json:"id"`
	Type     VehicleType `json:"vehicle_type"`
	Name     string      `json:"name"`
	Capacity int         `json:"capacity"`
}
