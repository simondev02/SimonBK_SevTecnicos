package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/technicalServiceType"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// UpdateTechnicalServiceTypedHandler - Controlador para actualizar un Tipo de Servicio Tecnico por ID
// @Summary Actualiza un Tipo de Servicio Tecnico por ID
// @Description Actualiza un Tipo de Servicio tecnico por ID con los datos proporcionados
// @Tags TechnicalServiceType
// @Accept json
// @Produce json
// @Param id path int true "ID del tipo de servicio Tecnico a actualizar"
// @Param vb body swagger.TechnicalServiceType true "Datos del Tipo de Servicio Tecnico a actualizar"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalServiceType/{id} [put]
func UpdateTechnicalServiceTypeHandler(c *gin.Context) {
	// Obtener el ID del técnico desde el URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Vincular el cuerpo de la solicitud a la estructura de TechnicalServiceType
	var tech models.TechnicalServiceType
	if err := c.ShouldBindJSON(&tech); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el UserId desde el contexto
	userIDInterface, exists := c.Get("UserId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el UserId"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserId en formato incorrecto"})
		return
	}

	// Llamar a la función UpdateTechnician pasando el ID del técnico y el UserId
	err = services.UpdateTechnicalServiceTyupe(db.DBConn, uint(id), &userID, &tech)
	if err != nil {
		// Comprobar si el error es debido a una violación de la restricción única de Name
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"idx_technical_service_types_name\"") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Este tipo de servicio técnico ya existe"})
			return
		}
		if strings.Contains(err.Error(), "tipo de servicio técnico no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el tipo de servicio técnico"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tipo de servicio técnico actualizado exitosamente"})
}
