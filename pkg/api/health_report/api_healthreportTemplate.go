package health_report

import (
	controller "autotec/pkg/controller/health_report"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllHealthReportTemplate(c echo.Context) error {

	data, err := controller.GetAllHealthReportTemplates()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
