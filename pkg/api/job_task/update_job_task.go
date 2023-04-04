package job_task

import (
	controller "autotec/pkg/controller/job_task"
	"autotec/pkg/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UpdateJobTask(c echo.Context) error {
	job := entity.JobTask{}
	if error := c.Bind(&job); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.UpdateJobTask(&job)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
