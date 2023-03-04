package job

import (
	controller "autotec/pkg/controller/job"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func AddNewJob(c echo.Context) error {
	job := entity.Job{}
	if error := c.Bind(&job); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.AddrNewJob(&job)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
