package rest_controller

import (
	"autotec/pkg/api/health_report"
	"autotec/pkg/api/job"
	"autotec/pkg/api/pre_repair_estimate"
	"autotec/pkg/api/user"
	"autotec/pkg/api/vehicle"
	"autotec/pkg/api/work_progress"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
)

func EchoController(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	normalRoutes := e.Group("/autotec")

	NormalRoutes(normalRoutes)

	SwaggerAPIDoc(normalRoutes)

}

func NormalRoutes(g *echo.Group) {
	g.POST("/v1/api/user", user.AddNewUser)
	g.POST("/v1/api/user/login", user.Login)
	g.GET("/v1/api/user", user.GetAllUsers)

	g.POST("/v1/api/vehicle", vehicle.AddNewVehicle)
	g.PUT("/v1/api/vehicle", vehicle.UpdateVehicle)
	g.GET("/v1/api/vehicle", vehicle.GetVehiclesBtCustomer)

	g.POST("/v1/api/pre_repair_estimate", pre_repair_estimate.AddNewPreRepairEstimate)
	g.PUT("/v1/api/pre_repair_estimate", pre_repair_estimate.UpdatePreRepairEstimate)
	g.GET("/v1/api/pre_repair_estimates", pre_repair_estimate.GetAllPreRepairEstimates)

	g.POST("/v1/api/job", job.AddNewJob)
	g.PUT("/v1/api/job", job.UpdateJob)
	g.GET("/v1/api/job", job.GetAllJob)

	g.POST("/v1/api/health_report", health_report.AddNewHealthReport)
	g.PUT("/v1/api/health_report", health_report.UpdateHealthReport)
	g.GET("/v1/api/health_report", health_report.GetAllHealthReport)

	g.POST("/v1/api/work_progress", work_progress.AddNewWorkProgress)
	g.PUT("/v1/api/work_progress", work_progress.UpdateWorkProgress)
	g.GET("/v1/api/work_progress", work_progress.GetAllWorkProgress)
}

func SwaggerAPIDoc(g *echo.Group) {
	g.GET("/v1/api/swagger/*any", echoswagger.WrapHandler)
}
