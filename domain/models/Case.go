package models

import "gorm.io/gorm"

type Case struct {
	gorm.Model
	Number     *uint
	FkCompany  *uint
	FkCustomer *uint
	CreatedBy  *uint
	DeletedBy  *uint
	UpdatedBy  *uint
	Company    Company  `gorm:"foreignKey:FkCompany"`
	Customer   Customer `gorm:"foreignKey:FkCustomer"`
}
