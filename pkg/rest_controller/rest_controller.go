package rest_controller

import (
	job "autotec/pkg/api/job"
	pre_repair_estimate "autotec/pkg/api/pre_repair_estimate"
	user "autotec/pkg/api/user"
	vehicle "autotec/pkg/api/vehicle"

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
}

func SwaggerAPIDoc(g *echo.Group) {
	g.GET("/v1/api/swagger/*any", echoswagger.WrapHandler)
}
