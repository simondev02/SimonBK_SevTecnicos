package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/Technicians"
	"SimonBK_SevTecnicos/infra/db"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateTechniciansHandler - Controlador para crear un nuevo técnico
// @Summary Crea una nuevo técnico con los datos proporcionados
// @Description Crea un nuevio técnico con los datos proporcionados
// @Tags Technicians
// @Accept json
// @Produce json
// @Param vb body swagger.Technician true "Datos de la marca de vehículo a crear"
// @Success 201 {object} swagger.Technician
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technician/ [post]
func CreateTechniciansHandler(c *gin.Context) {
	var vb models.Technicians
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserId").(uint)

	if err := services.CreateTechnicians(db.DBConn, &vb, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("Técnico creado exitosamente con ID: %d", vb.ID),
		"id":      vb.ID,
	})
}
