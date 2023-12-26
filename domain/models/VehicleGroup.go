package models

import "gorm.io/gorm"

type VehicleGroup struct {
	gorm.Model
	Name               *string `gorm:"unique"`
	Description        *string
	FkVehicleGroupType *uint
	VehicleGroupType   VehicleGroupType `gorm:"foreignKey:FkVehicleGroupType"`
	CreatedBy          *uint
	DeletedBy          *uint
	UpdatedBy          *uint
}
