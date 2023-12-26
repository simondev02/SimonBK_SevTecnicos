package services

import (
	"SimonBK_SevTecnicos/api/views"
	"SimonBK_SevTecnicos/domain/models"
	"fmt"

	"gorm.io/gorm"
)

func GetAllTechnicians(db *gorm.DB, technician *string, page int, pageSize int) (views.Return, error) {
	var techs []models.Technicians
	query := db

	if technician != nil && *technician != "" {
		query = query.Where("name ILIKE ?", "%"+*technician+"%")
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.Technicians{}).Count(&total).Error; err != nil {
		return views.Return{}, fmt.Errorf("error al obtener el total de registros: %w", err)
	}

	// Calcular el offset basado en la página y el tamaño de página
	offset := (page - 1) * pageSize

	// Ahora aplicamos el límite y el offset para obtener los registros de la página actual
	query = query.Offset(offset).Limit(pageSize)

	if err := query.Find(&techs).Error; err != nil {
		return views.Return{}, err
	}

	var response []interface{}
	for _, tech := range techs {
		techResponse := views.Technician{
			ID:    tech.ID,
			Name:  tech.Name,
			Phone: tech.Phone,
			Email: tech.Email,
		}
		response = append(response, techResponse)
	}

	// Devolver una instancia de views.Return con los valores apropiados
	return views.Return{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
		Results:  response,
	}, nil
}
