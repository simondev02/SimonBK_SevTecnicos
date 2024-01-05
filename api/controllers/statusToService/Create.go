package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/statusToService"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateStatusToServiceHandler - Controlador para crear una nueva asignacion de estado al servicio.
// @Summary Crea una nueva asignacion de estado en los servicios con los datos proporcionados
// @Description Crea una nueva asignacionde  de estado con los datos proporcionados
// @Tags StatusToService
// @Accept json
// @Produce json
// @Param vb body swagger.TechnicalServiceToStatus true "Datos de la asignacion del estado al servicio a crear"
// @Success 201 {object} swagger.TechnicalServiceToStatus
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalServiceToStatus/ [post]
func CreateStatusToServiceHandler(c *gin.Context) {
	var vb models.TechnicalServiceToStatus
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserId").(uint)

	err := services.CreateStatusToService(db.DBConn, &vb, userID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Asignacion del Estado al servicio creado exitosamente",
	})
}
