package models

import "gorm.io/gorm"

type VehicleBrand struct {
	gorm.Model
	Name *string
}
