package routers

import (
	statustoservice "SimonBK_SevTecnicos/api/controllers/statusToService"
	technicalservice "SimonBK_SevTecnicos/api/controllers/technicalService"
	technicalservicestatus "SimonBK_SevTecnicos/api/controllers/technicalServiceStatus"
	technicalservicetype "SimonBK_SevTecnicos/api/controllers/technicalServiceType"
	technician "SimonBK_SevTecnicos/api/controllers/technician"
	techniciantoservice "SimonBK_SevTecnicos/api/controllers/techniciansToService"
	"SimonBK_SevTecnicos/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {

	// Validacion de acces Token
	r.Use(middleware.ValidateTokenMiddleware())

	// Grupo de rutas para servicios tecnicos
	technicians := r.Group("/technician")
	{
		technicians.GET("/", technician.GetAllTechniciansController)
		technicians.GET("/:id", technician.GetTechnicianByIDHandler)
		technicians.POST("/", technician.CreateTechniciansHandler)
		technicians.PUT("/:id", technician.UpdateTechnicianHandler)
		technicians.DELETE("/:id", technician.DeleteTechnicianHandler)
	}

	// Grupo de rutas para tipos de servicio tecnico
	technicalServiceType := r.Group("/technicalServiceType")
	{
		technicalServiceType.POST("/", technicalservicetype.CreateTechnicalserviceTypeHandler)
		technicalServiceType.DELETE("/:id", technicalservicetype.DeleteTechnicalServiceTypeHandler)
		technicalServiceType.GET("/", technicalservicetype.GetAllTechnicalServiceTypeController)
		technicalServiceType.GET("/:id", technicalservicetype.GetTechnicalServiceTypeByIDHandler)
		technicalServiceType.PUT("/:id", technicalservicetype.UpdateTechnicalServiceTypeHandler)
	}

	// Grupo de rutas para estados de servicio tecnico
	technicalServiceStatus := r.Group("/technicalServiceStatus")
	{
		technicalServiceStatus.POST("/", technicalservicestatus.CreateTechnicalserviceStatusHandler)
		technicalServiceStatus.DELETE("/:id", technicalservicestatus.DeleteTechnicalServiceStatusHandler)
		technicalServiceStatus.GET("/", technicalservicestatus.GetAllTechnicalServiceStatusHandler)
		technicalServiceStatus.GET("/:id", technicalservicestatus.GetTechnicalServiceStatusByIDHandler)
		technicalServiceStatus.PUT("/:id", technicalservicestatus.UpdateTechnicalServiceStatusHandler)
	}
	// Grupo de rutas para servicios tecnicos
	technicalService := r.Group("/technicalService")
	{
		technicalService.GET("/", technicalservice.GetAllTechnicalServiceHandler)
		technicalService.POST("/", technicalservice.CreateTechnicalserviceHandler)
	}
	// Grupo de asignacion de tecnico a servicio
	techniciansToService := r.Group("/techniciansToService")
	{
		techniciansToService.POST("/", techniciantoservice.CreateTechniciansToServiceHandler)
	}
	// Grupo de asignacion de estados a servicios
	statusToService := r.Group("/technicalServiceToStatus")
	{

		statusToService.POST("/", statustoservice.CreateStatusToServiceHandler)
	}
}
