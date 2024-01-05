package views

type Technician struct {
	ID    uint    `json:"id"`
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
	Email *string `json:"email"`
	Dni   *uint   `json:"dni"`
}
