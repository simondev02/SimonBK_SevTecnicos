package models

import "gorm.io/gorm"

type TechniciansToTechnicialService struct {
	gorm.Model
	FkTechnician       *uint
	FkTechnicalService *uint
	CreatedBy          *uint
	DeletedBy          *uint
	UpdatedBy          *uint
	Technicians        Technicians      `gorm:"foreignKey:FkTechnician"`
	TechnicalService   TechnicalService `gorm:"foreignKey:FkTechnicalService"`
}
