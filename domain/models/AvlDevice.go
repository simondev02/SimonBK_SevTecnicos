package models

import (
	"gorm.io/gorm"
)

type AvlDevice struct {
	gorm.Model
	IMEI           *uint `gorm:"unique"`
	FkFwVersion    *uint
	FkReference    *uint
	FkManufacturer *uint
	FkCompany      *uint
	FkSimCard      *uint   `gorm:"unique"`
	Company        Company `gorm:"foreignKey:FkCompany"`
	DeletedBy      *uint
	UpdatedBy      *uint
}
