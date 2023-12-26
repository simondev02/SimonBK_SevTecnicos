package models

import "gorm.io/gorm"

type VehicleToCredit struct {
	gorm.Model
	FkVehicle *uint
	FkCredit  *uint
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
	Vehicle   Vehicle `gorm:"foreignKey:FkVehicle"`
	Credit    Credit  `gorm:"foreignKey:FkCredit"`
}
