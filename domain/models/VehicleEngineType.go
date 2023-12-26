package models

import "gorm.io/gorm"

type EngineType struct {
	gorm.Model
	Name *string `gorm:"type:varchar(20);unique"`
}
