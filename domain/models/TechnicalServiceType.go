package models

import "gorm.io/gorm"

type TechnicalServiceType struct {
	gorm.Model
	Name      *string `gorm:"unique"`
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
}
