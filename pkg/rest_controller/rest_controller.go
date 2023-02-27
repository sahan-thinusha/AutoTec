package rest_controller

import (
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
}

func SwaggerAPIDoc(g *echo.Group) {
	g.GET("/v1/api/swagger/*any", echoswagger.WrapHandler)
}
