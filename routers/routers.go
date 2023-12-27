package routers

import (
	technician "SimonBK_SevTecnicos/api/controllers/technician"
	"SimonBK_SevTecnicos/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {

	// Validacion de acces Token
	r.Use(middleware.ValidateTokenMiddleware())

	// Grupo de rutas para vehiculos
	technicians := r.Group("/technician")
	{
		technicians.GET("/", technician.GetAllTechniciansController)
		technicians.POST("/", technician.CreateTechniciansHandler)
		technicians.PUT("/:id", technician.UpdateTechnicianHandler)
		technicians.DELETE("/:id", technician.DeleteTechnicianHandler)
	}

}
