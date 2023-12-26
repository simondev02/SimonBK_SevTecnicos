package models

import "gorm.io/gorm"

type TechnicalServiceToStatus struct {
	gorm.Model
	FkTechnicalService *uint
	FkTechnicalStatus  *uint
}
