package models

import "gorm.io/gorm"

type FuelType struct {
	gorm.Model
	Name *string `json:"name" gorm:"type:varchar(20);unique"`
}
