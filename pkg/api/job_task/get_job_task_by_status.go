package job_task

import (
	controller "autotec/pkg/controller/job_task"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllTasksByEmployeeAndStatus(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	uid := c.QueryParam("uid")
	status := c.QueryParam("status")

	data, err := controller.GetAllJobTasksByStatusAndEmployee(index, limit, uid, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
