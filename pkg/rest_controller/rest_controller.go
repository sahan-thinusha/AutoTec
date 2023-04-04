package rest_controller

import (
	"autotec/pkg/api/health_report"
	"autotec/pkg/api/job"
	"autotec/pkg/api/job_task"
	"autotec/pkg/api/pre_repair_estimate"
	"autotec/pkg/api/user"
	"autotec/pkg/api/vehicle"
	"autotec/pkg/api/work_progress"
	"autotec/pkg/env"
	"crypto/subtle"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
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

	securedRoutes := e.Group("/autotec")

	securedRoutes.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(env.SigningKey),
	}))

	securedRoutes.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, ok := c.Get("user").(*jwt.Token)
			if !ok {
				return errors.New("JWT token missing or invalid")
			}
			_, ok = token.Claims.(jwt.MapClaims)
			if !ok {
				return errors.New("failed to cast claims as jwt.MapClaims")
			}
			return next(c)
		}
	})

	basicSecured := e.Group("/autotec")

	basicSecured.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("autotec")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("autotec@123")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	SecuredRoutes(securedRoutes)
	BasicSecuredRoutes(basicSecured)
	SwaggerAPIDoc(normalRoutes)

}

func BasicSecuredRoutes(g *echo.Group) {
	g.POST("/v1/api/user/login", user.Login)
}

func SecuredRoutes(g *echo.Group) {
	g.GET("/v1/api/user", user.GetAllUsers)
	g.POST("/v1/api/user", user.AddNewUser)

	g.POST("/v1/api/vehicle", vehicle.AddNewVehicle)
	g.PUT("/v1/api/vehicle", vehicle.UpdateVehicle)
	g.GET("/v1/api/vehicle", vehicle.GetVehiclesBtCustomer)

	g.POST("/v1/api/pre_repair_estimate", pre_repair_estimate.AddNewPreRepairEstimate)
	g.PUT("/v1/api/pre_repair_estimate", pre_repair_estimate.UpdatePreRepairEstimate)
	g.GET("/v1/api/pre_repair_estimates", pre_repair_estimate.GetAllPreRepairEstimates)

	g.POST("/v1/api/job", job.AddNewJob)
	g.PUT("/v1/api/job", job.UpdateJob)
	g.GET("/v1/api/job", job.GetAllJob)
	g.GET("/v1/api/job/detail", job.GetJobByID)

	g.POST("/v1/api/health_report", health_report.AddNewHealthReport)
	g.PUT("/v1/api/health_report", health_report.UpdateHealthReport)
	g.GET("/v1/api/health_report", health_report.GetAllHealthReport)

	g.POST("/v1/api/work_progress", work_progress.AddNewWorkProgress)
	g.PUT("/v1/api/work_progress", work_progress.UpdateWorkProgress)
	g.GET("/v1/api/work_progress", work_progress.GetAllWorkProgress)

	g.POST("/v1/api/job_task", job_task.AddNewJobTask)
	g.PUT("/v1/api/job_task", job_task.UpdateJobTask)
}

func NormalRoutes(g *echo.Group) {
}

func SwaggerAPIDoc(g *echo.Group) {
	g.GET("/v1/api/swagger/*any", echoswagger.WrapHandler)
}
