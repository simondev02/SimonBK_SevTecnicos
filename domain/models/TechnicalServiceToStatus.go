package models

import "gorm.io/gorm"

type TechnicalServiceToStatus struct {
	gorm.Model
	FkTechnicalService *uint
	FkTechnicalStatus  *uint
	CreatedBy          *uint
	DeletedBy          *uint
	UpdatedBy          *uint
}
