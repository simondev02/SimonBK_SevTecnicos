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
// @Summary Obtiene Tipo de Servicio Tecnico específico
// @Description Obtiene un Tipo de Servicio Tecnico por su ID específico
// @Tags TechnicalServiceType
// @Produce json
// @Param id path int true "ID del Tipo de servicio Tecnico"
// @Success 200 {object} map[string]string "Detalles del Tipo de servicio técnico"
// @Failure 400 {object} map[string]string "Error: ID inválido"
// @Failure 404 {object} map[string]string "Error: Tipo de servicio tecnico no encontrado"
// @Router /technicalServiceType/{id} [get]
func GetTechnicalServiceTypeByIDHandler(c *gin.Context) {
	// Obtener el ID del técnico desde el URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar a la función GetTechnicalServiceTypeByID pasando el ID del técnico
	tech, err := services.GetTechnicalServiceTypeByID(db.DBConn, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "tipo de servicio técnico no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el tipo de servicio técnico"})
		return
	}

	c.JSON(http.StatusOK, tech)
}
