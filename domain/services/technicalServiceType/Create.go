package services

import (
	"SimonBK_SevTecnicos/domain/models"

	"gorm.io/gorm"
)

// CreateTechnicalServiceType - Crea una nueva Tipo de servicio tecnici en la base de datos
func CreateTechnicalServiceType(db *gorm.DB, vb *models.TechnicalServiceType, userID uint) error {
	vb.CreatedBy = &userID
	if err := db.Create(vb).Error; err != nil {
		return err
	}
	return nil
}
