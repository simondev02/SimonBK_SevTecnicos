package Controllers

import (
	services "SimonBK_SevTecnicos/domain/services/technicalServiceType"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Obtiene todos los Tipos de Servicios Tecnicos
// @Description Obtiene una lista de todos los Tipos de Servicios Tecnicos.
// @Tags TechnicalServiceType
// @Produce json
// @Param page query int false "Número de página para la paginación" default(1)
// @Param pageSize query int false "Tamaño de página para la paginación" default(10)
// @Param name query string false "Nombre del Tipo de servicio Tecnico"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalServiceType/ [get]
func GetAllTechnicalServiceTypeController(c *gin.Context) {
	name := c.Query("name")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	techs, err := services.GetAllTechnicalServiceTypes(db.DBConn, &name, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, techs)
}
