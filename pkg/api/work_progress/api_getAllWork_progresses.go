package work_progress

import (
	controller "autotec/pkg/controller/work_progress"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllWorkProgress(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	data, err := controller.GetAllWorkProgress(index, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
