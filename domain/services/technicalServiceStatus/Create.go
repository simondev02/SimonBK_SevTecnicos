package services

import (
	"SimonBK_SevTecnicos/domain/models"

	"gorm.io/gorm"
)

// CreateVehicleBrand - Crea una nueva marca de vehículo en la base de datos
func CreateTechnicalServiceStatus(db *gorm.DB, vb *models.TechnicalServiceStatus, userID uint) error {
	vb.CreatedBy = &userID
	if err := db.Create(vb).Error; err != nil {
		return err
	}
	return nil
}
