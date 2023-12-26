package models

import "gorm.io/gorm"

type Seller struct {
	gorm.Model
	Name      *string
	Telefono  *string
	Email     *string
	FkCompany *uint
	Company   Company `gorm:"foreignKey:FkCompany"`
	CreatedBy *uint
	DeletedBy *uint
	UpdatedBy *uint
}
