package services

import (
	"SimonBK_SevTecnicos/api/views"
	"SimonBK_SevTecnicos/domain/models"
	"fmt"

	"gorm.io/gorm"
)

func GetAllTechnicalService(db *gorm.DB, plate *string, serviceType *string, page int, pageSize int) (views.Return, error) {
	var techs []models.TechnicalService
	query := db.
		Preload("Vehicle").
		Preload("TechnicalServiceType").
		Preload("Vehicle.Company").
		Preload("Vehicle.Customer").
		Joins("JOIN (SELECT DISTINCT ON (fk_technical_service) * FROM technicians_to_technicial_services ORDER BY fk_technical_service, created_at DESC) AS latest_technician ON latest_technician.fk_technical_Service = technical_services.id").
		Preload("TechniciansToTechnicialService.Technicians")

	if plate != nil && *plate != "" {
		query = query.Where("Vehicle.plate ILIKE ?", "%"+*plate+"%")
	}

	if serviceType != nil && *serviceType != "" {
		query = query.Where("TechnicalServiceType.name ILIKE ?", "%"+*serviceType+"%")
	}

	// Calcular el total de registros
	var total int64
	if err := query.Model(&models.TechnicalService{}).Count(&total).Error; err != nil {
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
		techResponse := views.TechnicalService{
			Id:          tech.ID,
			Plate:       tech.Vehicle.Plate,
			Vin:         tech.Vehicle.Vin,
			ServiceType: tech.TechnicalServiceType.Name,
			Technician:  tech.TechniciansToTechnicialService[0].Technicians.Name,
			StartDate:   tech.StartDate,
			EndDate:     tech.EndDate,
			Company:     tech.Vehicle.Company.Name,
			Customer:    tech.Vehicle.Customer.Name,
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
