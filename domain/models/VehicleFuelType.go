package models

import "gorm.io/gorm"

type FuelType struct {
	gorm.Model
	ID        *uint   `json:"id"`
	Name      *string `json:"name" gorm:"type:varchar(20);unique"`
	DeletedBy *uint   `json:"deletedBy"`
	UpdatedBy *uint   `json:"updatedBy"`
}
