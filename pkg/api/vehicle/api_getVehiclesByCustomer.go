package vehicle

import (
	controller "autotec/pkg/controller/vehicle"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetVehiclesBtCustomer(c echo.Context) error {

	userId := c.QueryParam("user")

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	data, err := controller.GetVehiclesByCustomer(index, limit, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
