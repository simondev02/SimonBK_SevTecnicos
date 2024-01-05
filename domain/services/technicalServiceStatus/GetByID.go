package services

import (
	"SimonBK_SevTecnicos/api/views"
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetTechnicalServiceStatusByID(db *gorm.DB, id uint) (*views.TechnicalServiceStatusById, error) {
	var tech models.TechnicalServiceStatus
	query := db.Where("id = ?", id) // Filtra la consulta por el ID proporcionado

	if err := query.First(&tech).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("estado no encontrado")
		}
		return nil, err
	}

	techView := &views.TechnicalServiceStatusById{
		ID:        tech.ID,
		Name:      tech.Name,
		CreatedAt: &tech.CreatedAt,
	}

	return techView, nil
}
