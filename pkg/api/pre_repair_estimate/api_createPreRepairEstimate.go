package pre_repair_estimate

import (
	controller "autotec/pkg/controller/pre_repair_estimate"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func AddNewPreRepairEstimate(c echo.Context) error {
	preRepairEstimate := entity.PreRepairEstimate{}
	if error := c.Bind(&preRepairEstimate); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.AddrNewPreRepairEstimate(&preRepairEstimate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}
