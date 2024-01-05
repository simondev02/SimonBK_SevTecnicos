package services

import (
	"SimonBK_SevTecnicos/domain/models"

	"gorm.io/gorm"
)

// CreateTechnicianstoService - Crea una nueva asignacion de técnico al servicio  en la base de datos
func CreateTechniciansToService(db *gorm.DB, vb *models.TechniciansToTechnicialService, userID uint) error {

	vb.CreatedBy = &userID
	if err := db.Create(vb).Error; err != nil {
		return err
	}
	return nil
}
