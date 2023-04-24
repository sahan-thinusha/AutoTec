package efficiency

import (
	controller "autotec/pkg/controller/efficiency"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllEmployeeEfficiency(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	data, err := controller.GetTechnicianEfficiency(index, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
