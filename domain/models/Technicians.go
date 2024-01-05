package models

import "gorm.io/gorm"

type Technicians struct {
	gorm.Model
	Name      *string
	Phone     *string
	Email     *string
	Dni       *uint `gorm:"unique;not null"`
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
}
