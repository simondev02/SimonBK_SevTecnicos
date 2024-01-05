package services

import (
	"SimonBK_SevTecnicos/api/views"
	"SimonBK_SevTecnicos/domain/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func GetTechnicianByID(db *gorm.DB, id uint) (*views.TechnicianById, error) {
	var tech models.Technicians
	query := db.Where("id = ?", id) // Filtra la consulta por el ID proporcionado

	if err := query.First(&tech).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("técnico no encontrado")
		}
		return nil, err
	}

	techView := &views.TechnicianById{
		ID:        tech.ID,
		Name:      tech.Name,
		Phone:     tech.Phone,
		Email:     tech.Email,
		Dni:       tech.Dni,
		CreatedAt: &tech.CreatedAt,
	}

	return techView, nil
}
