package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name      string
	Dni       string `gorm:"uniqueIndex:idx_fkcompany_identification"`
	FkCompany int    `gorm:"uniqueIndex:idx_fkcompany_identification"`
	DeletedBy *uint
	UpdatedBy *uint
	Company   Company `gorm:"foreignKey:FkCompany"`
}
