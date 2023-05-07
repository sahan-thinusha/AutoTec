package efficiency

import (
	controller "autotec/pkg/controller/efficiency"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllEmployeeEfficiency(c echo.Context) error {

	data, err := controller.GetTechnicianEfficiency()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
