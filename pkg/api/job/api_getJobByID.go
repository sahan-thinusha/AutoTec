package job

import (
	controller "autotec/pkg/controller/job"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetJobByID(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	id := c.QueryParam("id")

	data, err := controller.GetJobByID(index, limit, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
