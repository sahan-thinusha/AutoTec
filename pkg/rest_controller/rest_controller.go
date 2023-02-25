package rest_controller

import (
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

}

func SwaggerAPIDoc(g *echo.Group) {
	g.GET("/v1/api/swagger/*any", echoswagger.WrapHandler)
}
