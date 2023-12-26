package models

import "gorm.io/gorm"

type VehicleLine struct {
	gorm.Model
	Name           *string
	FkVehicleBrand *uint
	VehicleBrand   VehicleBrand `gorm:"foreignKey:FkVehicleBrand"`
}
