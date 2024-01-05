package Controllers

import (
	services "SimonBK_SevTecnicos/domain/services/technicalServiceStatus"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Obtiene todos los Estados de Servicios Tecnicos
// @Description Obtiene una lista de todos Estados de servicios tecnicos.
// @Tags TechnicalServiceStatus
// @Produce json
// @Param page query int false "Número de página para la paginación" default(1)
// @Param pageSize query int false "Tamaño de página para la paginación" default(10)
// @Param name query string false "Nombre del Estado"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalServiceStatus/ [get]
func GetAllTechnicalServiceStatusHandler(c *gin.Context) {
	name := c.Query("name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	techs, err := services.GetAllTechnicalServiceStatus(db.DBConn, &name, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, techs)
}
