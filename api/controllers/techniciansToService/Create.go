package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/techniciansToService"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateTechniciansToServiceHandler - Controlador para crear una nueva asignacion de tecnico al servicio.
// @Summary Crea una nueva asignacion técnico en los servicios con los datos proporcionados
// @Description Crea una nueva asignacionde  técnico con los datos proporcionados
// @Tags TechniciansToService
// @Accept json
// @Produce json
// @Param vb body swagger.TechniciansToService true "Datos de la asignacion de técnico al servicio a crear"
// @Success 201 {object} swagger.TechniciansToService
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /techniciansToService/ [post]
func CreateTechniciansToServiceHandler(c *gin.Context) {
	var vb models.TechniciansToTechnicialService
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserId").(uint)

	err := services.CreateTechniciansToService(db.DBConn, &vb, userID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Asignacion del Técnico al servicio creado exitosamente",
	})
}
