package Controllers

import (
	services "SimonBK_SevTecnicos/domain/services/technicians"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Obtiene Tecnico específico
// @Description Obtiene un Tecnico por su ID específico
// @Tags Technicians
// @Produce json
// @Param id path int true "ID del Tecnico"
// @Success 200 {object} map[string]string "Detalles del técnico"
// @Failure 400 {object} map[string]string "Error: ID inválido"
// @Failure 404 {object} map[string]string "Error: tecnico no encontrado"
// @Router /technician/{id} [get]
func GetTechnicianByIDHandler(c *gin.Context) {
	// Obtener el ID del técnico desde el URL
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamar a la función GetTechnicianByID pasando el ID del técnico
	tech, err := services.GetTechnicianByID(db.DBConn, uint(id))
	if err != nil {
		if strings.Contains(err.Error(), "técnico no encontrado") {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el técnico"})
		return
	}

	c.JSON(http.StatusOK, tech)
}
