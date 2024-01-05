package models

import "gorm.io/gorm"

type TechnicalServiceStatus struct {
	gorm.Model
	Name      *string `gorm:"unique"`
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
}
