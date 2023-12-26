package Controllers

import (
	services "SimonBK_SevTecnicos/domain/services/Technicians"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Obtiene todos los Tecnicos
// @Description Obtiene una lista de todos los Tecnicos.
// @Tags Technicians
// @Produce json
// @Param page query int false "Número de página para la paginación" default(1)
// @Param pageSize query int false "Tamaño de página para la paginación" default(10)
// @Param technicians query string false "Nombre del Tecnico"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technician/ [get]
func GetAllTechniciansController(c *gin.Context) {
	technician := c.Query("technician")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	techs, err := services.GetAllTechnicians(db.DBConn, &technician, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, techs)
}
