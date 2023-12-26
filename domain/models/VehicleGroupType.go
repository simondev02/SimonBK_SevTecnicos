package models

import "gorm.io/gorm"

type VehicleGroupType struct {
	gorm.Model
	Name      *string
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
}
