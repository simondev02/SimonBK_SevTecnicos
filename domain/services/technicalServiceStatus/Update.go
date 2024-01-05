package services

import (
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func UpdateTechnicalServiceStatus(db *gorm.DB, id uint, userID *uint, tech *models.TechnicalServiceStatus) error {
	var existingTech models.TechnicalServiceStatus

	// Primero, encuentra el técnico por ID
	if err := db.First(&existingTech, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("estado no encontrado no encontrado")
		}
		return err
	}

	// Configura el usuario que actualizó el registro
	tech.UpdatedBy = userID

	// Luego, actualiza el registro
	if err := db.Model(&existingTech).Updates(tech).Error; err != nil {
		return fmt.Errorf("error al actualizar el estado: %v", err)
	}

	return nil
}
