package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/technicians"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// UpdateTechniciansdHandler - Controlador para actualizar un Tecnico por ID
// @Summary Actualiza un Tecnico por ID
// @Description Actualiza un tecnico por ID con los datos proporcionados
// @Tags Technicians
// @Accept json
// @Produce json
// @Param id path int true "ID del Tecnico a actualizar"
// @Param vb body swagger.Technician true "Datos de la marca de vehículo a actualizar"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technician/{id} [put]
func UpdateTechnicianHandler(c *gin.Context) {
	// Obtener el ID del técnico desde el URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Vincular el cuerpo de la solicitud a la estructura de Technicians
	var tech models.Technicians
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
	err = services.UpdateTechnician(db.DBConn, uint(id), &userID, &tech)
	if err != nil {
		if strings.Contains(err.Error(), "técnico no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		// Comprobar si el error es debido a una violación de la restricción única de DNI
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El Numero de documento proporcionado ya existe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el técnico"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Técnico actualizado exitosamente"})
}
