package contracts

type User struct {
	ID       string   `json:"username"`
	Name     string   `json:"name"`
	Number   int      `json:"number"`
	Age      int      `json:"age"`
	Location Location `json:"location"`
	Ride     Ride     `json:"ride"`
}
