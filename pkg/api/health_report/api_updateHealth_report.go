package health_report

import (
	controller "autotec/pkg/controller/health_report"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func UpdateHealthReport(c echo.Context) error {
	healthReport := entity.HealthReport{}
	if error := c.Bind(&healthReport); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.UpdateHealthReport(&healthReport)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
