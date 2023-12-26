package models

import "gorm.io/gorm"

type VehicleToGroup struct {
	gorm.Model
	FkVehicle      *uint `gorm:"uniqueIndex:idx_vehicle_group"`
	FkVehicleGroup *uint `gorm:"uniqueIndex:idx_vehicle_group"`
	CreatedBy      *uint
	DeletedBy      *uint
	UpdatedBy      *uint
	Vehicle        Vehicle      `gorm:"foreignKey:FkVehicle"`
	VehicleGroup   VehicleGroup `gorm:"foreignKey:FkVehicleGroup"`
}
