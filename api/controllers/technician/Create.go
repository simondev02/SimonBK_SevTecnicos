package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/technicians"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strings"

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

	// Comprobar si el DNI es nulo
	if vb.Dni == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El número de documento es obligatorio"})
		return
	}

	userID := c.MustGet("UserId").(uint)

	err := services.CreateTechnicians(db.DBConn, &vb, userID)
	if err != nil {
		// Comprobar si el error es debido a una violación de la restricción única de DNI
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El número de documento ya existe"})
			return
		}
		// Comprobar si el error es debido a un valor nulo en el DNI
		if strings.Contains(err.Error(), "null value in column \"dni\" of relation \"technicians\" violates not-null constraint") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El número de documento es obligatorio"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Técnico creado exitosamente",
	})
}
