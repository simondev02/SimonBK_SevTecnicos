package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name      *string
	Dni       *string `gorm:"uniqueIndex:idx_fkcompany_identification"`
	FkCompany *uint   `gorm:"uniqueIndex:idx_fkcompany_identification"`
	Company   Company `gorm:"foreignKey:FkCompany"`
}
