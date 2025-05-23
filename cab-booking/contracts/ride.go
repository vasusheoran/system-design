package contracts

import "time"

type Ride struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	DriverID  string    `json:"driver_id"`
	Source    Location  `json:"source"`
	Dest      Location  `json:"dest"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
