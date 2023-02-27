package vehicle

import (
	controller "autotec/pkg/controller/vehicle"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func UpdateVehicle(c echo.Context) error {
	vehicle := entity.Vehicle{}
	if error := c.Bind(&vehicle); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.UpdateVehicle(&vehicle)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
