package Controllers

import (
	services "SimonBK_SevTecnicos/domain/services/technicalServiceType"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// DeleteTechnicalServiceTypeHandler - Controlador para eliminar un Tipo de Servicio Tecnico ID
// @Summary Elimina un Tipo de Servicio Tecnico por ID
// @Description Elimina un Tipo de Servicio Tecnico  de la base de datos basado en el ID proporcionado
// @Tags TechnicalServiceType
// @Accept json
// @Produce json
// @Param id path int true "ID del Tipo de Servicio tecnico a eliminar"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string "Mensaje de éxito al eliminar"
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 404 {object} map[string]string "Tipo de  Servicio Tecnico no encontrado"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /technicalServiceType/{id} [delete]
func DeleteTechnicalServiceTypeHandler(c *gin.Context) {
	// Obtener el ID del técnico desde el URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
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

	// Llamar a la función DeleteTechnician pasando el ID del técnico y el UserId
	err = services.DeleteTechnicalServiceType(db.DBConn, uint(id), &userID)
	if err != nil {
		if strings.Contains(err.Error(), "Tipo de Servicio técnico no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el técnico"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tipo de Servicio Técnico eliminado exitosamente"})
}
