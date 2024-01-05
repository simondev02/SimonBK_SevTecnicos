package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/technicalServiceType"
	"strings"

	"SimonBK_SevTecnicos/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateTechnicalServiceTypeHandler - Controlador para crear un nuevo tipo de servicio tecnico
// @Summary Crea una nuevo tipo de servicio tecnico con los datos proporcionados
// @Description Crea un nuevio tipo de servicio tecnico con los datos proporcionados
// @Tags TechnicalServiceType
// @Accept json
// @Produce json
// @Param vb body swagger.TechnicalServiceType true "Datos de la marca de vehículo a crear"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalServiceType/ [post]
func CreateTechnicalserviceTypeHandler(c *gin.Context) {
	var vb models.TechnicalServiceType
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserId").(uint)

	err := services.CreateTechnicalServiceType(db.DBConn, &vb, userID)
	if err != nil {
		// Comprobar si el error es debido a una violación de la restricción única de Name
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"idx_technical_service_types_name\"") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Este tipo de servicio técnico ya existe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tipo de servicio técnico creado exitosamente",
	})
}
