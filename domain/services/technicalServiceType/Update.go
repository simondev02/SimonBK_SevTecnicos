package services

import (
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func UpdateTechnicalServiceTyupe(db *gorm.DB, id uint, userID *uint, tech *models.TechnicalServiceType) error {
	var existingTech models.TechnicalServiceType

	// Primero, encuentra el tipo de servicio técnico por ID
	if err := db.First(&existingTech, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("tipo de servicio técnico no encontrado")
		}
		return err
	}

	// Configura el usuario que actualizó el registro
	tech.UpdatedBy = userID

	// Luego, actualiza el registro
	if err := db.Model(&existingTech).Updates(tech).Error; err != nil {
		return fmt.Errorf("error al actualizar el tipo de servicio técnico: %v", err)
	}

	return nil
}
