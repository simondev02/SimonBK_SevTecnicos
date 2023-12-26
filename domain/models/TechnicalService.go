package models

import "gorm.io/gorm"

type TechnicalService struct {
	gorm.Model
	FkVehicle              *uint
	FkTechnicalServiceType *uint
	CreatedBy              *uint
	DeletedBy              *uint
	UpdatedBy              *uint
	Vehicle                Vehicle              `gorm:"foreignKey:FkVehicle"`
	TechnicalServiceType   TechnicalServiceType `gorm:"foreignKey:FkTechnicalServiceType"`
}
