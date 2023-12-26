package models

import (
	"gorm.io/gorm"
)

type VehicleType struct {
	gorm.Model
	Name        *string
	Description *string
}
