package services

import (
	"SimonBK_SevTecnicos/api/views"
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetTechnicalServiceTypeByID(db *gorm.DB, id uint) (*views.TechnicalServiceTypeById, error) {
	var tech models.TechnicalServiceType
	query := db.Where("id = ?", id) // Filtra la consulta por el ID proporcionado

	if err := query.First(&tech).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("tipo de servicio técnico no encontrado")
		}
		return nil, err
	}

	techView := &views.TechnicalServiceTypeById{
		ID:        tech.ID,
		Name:      tech.Name,
		CreatedAt: &tech.CreatedAt,
	}

	return techView, nil
}
