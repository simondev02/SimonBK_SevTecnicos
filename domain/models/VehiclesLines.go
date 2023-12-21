package models

import "gorm.io/gorm"

type VehicleLine struct {
	gorm.Model
	ID             *uint
	Name           *string
	DeletedBy      *uint
	UpdatedBy      *uint
	FkVehicleBrand *uint
	VehicleBrand   VehicleBrand `gorm:"foreignKey:FkVehicleBrand"`
}
