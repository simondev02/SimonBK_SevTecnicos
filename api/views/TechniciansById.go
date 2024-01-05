package views

import "time"

type TechnicianById struct {
	ID        uint       `json:"id"`
	Name      *string    `json:"name"`
	Phone     *string    `json:"phone"`
	Email     *string    `json:"email"`
	Dni       *uint      `json:"dni"`
	CreatedAt *time.Time `json:"created_at"`
}
