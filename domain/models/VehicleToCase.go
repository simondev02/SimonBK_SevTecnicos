package models

import "gorm.io/gorm"

type VehicleToCase struct {
	gorm.Model
	FkVehicle *uint
	FkCase    *uint
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
	Vehicle   Vehicle `gorm:"foreignKey:FkVehicle"`
	Case      Case    `gorm:"foreignKey:FkVehicle"`
}
