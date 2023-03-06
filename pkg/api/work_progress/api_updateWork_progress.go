package work_progress

import (
	controller "autotec/pkg/controller/work_progress"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func UpdateWorkProgress(c echo.Context) error {
	workProgress := entity.WorkProgress{}
	if error := c.Bind(&workProgress); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.UpdateWorkProgress(&workProgress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
