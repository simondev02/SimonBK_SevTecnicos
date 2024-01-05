package services

import (
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func DeleteTechnicalServiceStatus(db *gorm.DB, id uint, userID *uint) error {
	var tech models.TechnicalServiceStatus

	// Primero, encuentra el técnico por ID
	if err := db.First(&tech, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("estado no encontrado no encontrado")
		}
		return err
	}

	// Establece el usuario que eliminó el registro
	tech.DeletedBy = userID

	// Luego, utiliza Update para establecer el campo DeletedByUserID antes de eliminar
	if err := db.Model(&tech).Updates(models.TechnicalServiceStatus{DeletedBy: userID}).Error; err != nil {
		return fmt.Errorf("error al configurar el usuario que eliminó: %v", err)
	}

	// Finalmente, usa Delete para que GORM llene automáticamente el campo DeletedAt
	if err := db.Delete(&tech).Error; err != nil {
		return fmt.Errorf("error al eliminar el técnico: %v", err)
	}

	return nil
}
