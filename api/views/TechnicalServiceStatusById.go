package views

import "time"

type TechnicalServiceStatusById struct {
	ID        uint
	Name      *string `gorm:"unique"`
	CreatedAt *time.Time
}
