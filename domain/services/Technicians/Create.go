package services

import (
	"SimonBK_SevTecnicos/domain/models"

	"gorm.io/gorm"
)

// CreateVehicleBrand - Crea una nueva marca de vehículo en la base de datos
func CreateTechnicians(db *gorm.DB, vb *models.Technicians) error {
	return db.Create(vb).Error
}
