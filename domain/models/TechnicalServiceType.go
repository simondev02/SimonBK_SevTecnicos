package models

import "gorm.io/gorm"

type TechnicalServiceType struct {
	gorm.Model
	Name      *string
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
}
