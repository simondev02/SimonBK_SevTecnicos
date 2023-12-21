package models

import (
	"gorm.io/gorm"
)

type VehicleType struct {
	gorm.Model
	DeletedBy   *uint
	UpdatedBy   *uint
	Name        *string
	Description *string
}
