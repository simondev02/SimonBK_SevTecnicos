package models

import "gorm.io/gorm"

type TechnicalServiceStatus struct {
	gorm.Model
	Name *string
}
