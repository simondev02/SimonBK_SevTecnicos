package services

import (
	"SimonBK_SevTecnicos/api/views"
	"SimonBK_SevTecnicos/domain/models"
	"fmt"

	"gorm.io/gorm"
)

func GetAllTechnicalServiceTypes(db *gorm.DB, name *string, page int, pageSize int) (views.Return, error) {
	var types []models.TechnicalServiceType
	query := db

	if name != nil && *name != "" {
		query = query.Where("name ILIKE ?", "%"+*name+"%")
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.TechnicalServiceType{}).Count(&total).Error; err != nil {
		return views.Return{}, fmt.Errorf("error al obtener el total de registros: %w", err)
	}

	// Calcular el offset basado en la página y el tamaño de página
	offset := (page - 1) * pageSize

	// Ahora aplicamos el límite y el offset para obtener los registros de la página actual
	query = query.Offset(offset).Limit(pageSize)

	if err := query.Find(&types).Error; err != nil {
		return views.Return{}, err
	}

	var response []interface{}
	for _, serviceType := range types {
		typeResponse := views.TechnicalServiceType{
			ID:   serviceType.ID,
			Name: serviceType.Name,
		}
		response = append(response, typeResponse)
	}

	// Devolver una instancia de views.Return con los valores apropiados
	return views.Return{
		Page:     page,
		PageSize: pageSize,
		Total:    int(total),
		Results:  response,
	}, nil
}
