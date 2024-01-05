package views

import "time"

type TechnicalServiceTypeById struct {
	ID        uint       `json:"id"`
	Name      *string    `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}
