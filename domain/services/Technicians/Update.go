package services

import (
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func UpdateTechnician(db *gorm.DB, id uint, userID *uint, tech *models.Technicians) error {
	var existingTech models.Technicians

	// Primero, encuentra el técnico por ID
	if err := db.First(&existingTech, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("técnico no encontrado")
		}
		return err
	}

	// Configura el usuario que actualizó el registro
	tech.UpdatedBy = userID

	// Luego, actualiza el registro
	if err := db.Model(&existingTech).Updates(tech).Error; err != nil {
		return fmt.Errorf("error al actualizar el técnico: %v", err)
	}

	return nil
}
