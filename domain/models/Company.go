package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string `gorm:"unique"`
	Nit       string `gorm:"unique"`
	Address   string
	Phone     string
	Email     string
	DeletedBy *uint
	UpdatedBy *uint
}
