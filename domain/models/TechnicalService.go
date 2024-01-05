package models

import (
	"time"

	"gorm.io/gorm"
)

type TechnicalService struct {
	gorm.Model
	FkVehicle                      *uint
	FkTechnicalServiceType         *uint
	StartDate                      *time.Time
	EndDate                        *time.Time
	CreatedBy                      *uint
	DeletedBy                      *uint
	UpdatedBy                      *uint
	Vehicle                        Vehicle                          `gorm:"foreignKey:FkVehicle"`
	TechnicalServiceType           TechnicalServiceType             `gorm:"foreignKey:FkTechnicalServiceType"`
	TechniciansToTechnicialService []TechniciansToTechnicialService `gorm:"foreignKey:FkTechnicalService"`
}
