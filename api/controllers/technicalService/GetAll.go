package controllers

import (
	services "SimonBK_SevTecnicos/domain/services/technicalService"
	"SimonBK_SevTecnicos/infra/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Obtiene todos los  Servicios Tecnicos
// @Description Obtiene una lista de todos servicios tecnicos.
// @Tags TechnicalService
// @Produce json
// @Param page query int false "Número de página para la paginación" default(1)
// @Param pageSize query int false "Tamaño de página para la paginación" default(10)
// @Param plate query string false "Placa del Vehiculo"
// @Param serviceType query string false "Nombre del tipo de servicio"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /technicalService/ [get]
func GetAllTechnicalServiceHandler(c *gin.Context) {
	plate := c.Query("plate")
	serviceType := c.Query("serviceType")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	techs, err := services.GetAllTechnicalService(db.DBConn, &plate, &serviceType, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, techs)
}
