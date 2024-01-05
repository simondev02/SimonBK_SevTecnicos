package controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/technicalService"
	"strings"

	"SimonBK_SevTecnicos/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateTechnicalServiceHandler - Controlador para crear un nuevo  servicio tecnico
// @Summary Crea una nuevo  servicio tecnico con los datos proporcionados
// @Description Crea un nuevio  servicio tecnico con los datos proporcionados
// @Tags TechnicalService
// @Accept json
// @Produce json
// @Param vb body swagger.TechnicalService true "Datos de la servicio a crear"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalService/ [post]
func CreateTechnicalserviceHandler(c *gin.Context) {
	var vb models.TechnicalService
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserId").(uint)

	err := services.CreateTechnicalService(db.DBConn, &vb, userID)
	if err != nil {
		if strings.Contains(err.Error(), "violates foreign key constraint \"fk_technical_services_vehicle\"") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Este vehículo no existe"})
			return
		}

		if strings.Contains(err.Error(), "endDate no puede ser menor o igual que startDate") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha de finalización no puede ser menor o igual que la fecha de inicio"})
			return
		}

		if strings.Contains(err.Error(), "endDate no puede existir si no se registra startDate") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha de finalización no puede existir si no se registra la fecha de inicio"})
			return
		}

		if strings.Contains(err.Error(), "violates unique constraint \"idx_vehicle_service_on_going\"") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ya existe un servicio activo para este vehículo y servicio técnico"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "servicio técnico creado exitosamente",
	})
}
