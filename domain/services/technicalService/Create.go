package services

import (
	"SimonBK_SevTecnicos/domain/models"
	"errors"

	"gorm.io/gorm"
)

// CreateTechnicalService - Crea un nuevo servicio técnico en la base de datos
func CreateTechnicalService(db *gorm.DB, vb *models.TechnicalService, userID uint) error {
	if vb.EndDate != nil && vb.StartDate == nil {
		return errors.New("endDate no puede existir si no se registra startDate")
	}

	if vb.EndDate != nil && vb.StartDate != nil && vb.EndDate.Before(*vb.StartDate) {
		return errors.New("endDate no puede ser menor o igual que startDate")
	}

	vb.CreatedBy = &userID
	if err := db.Create(vb).Error; err != nil {
		return err
	}
	return nil
}
