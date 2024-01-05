package Controllers

import (
	"SimonBK_SevTecnicos/domain/models"
	services "SimonBK_SevTecnicos/domain/services/technicalServiceStatus"

	"SimonBK_SevTecnicos/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateTechnicalServiceStatusHandler - Controlador para crear un nuevo estado de servicio tecnico
// @Summary Crea una nuevo estado de servicio tecnico con los datos proporcionados
// @Description Crea un nuevio estado de servicio tecnico con los datos proporcionados
// @Tags TechnicalServiceStatus
// @Accept json
// @Produce json
// @Param vb body swagger.TechnicalServiceStatus true "Datos de la estado a crear"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalServiceStatus/ [post]
func CreateTechnicalserviceStatusHandler(c *gin.Context) {
	var vb models.TechnicalServiceStatus
	if err := c.ShouldBindJSON(&vb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("UserId").(uint)

	err := services.CreateTechnicalServiceStatus(db.DBConn, &vb, userID)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "estado de servicio técnico creado exitosamente",
	})
}
