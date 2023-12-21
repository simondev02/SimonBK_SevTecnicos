package models

import "gorm.io/gorm"

type VehicleGroup struct {
	gorm.Model
	Name        *string
	Description *string
	DeletedBy   *uint `json:"deletedBy"`
	UpdatedBy   *uint `json:"updatedBy"`
}
