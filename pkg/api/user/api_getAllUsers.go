package user

import (
	controller "autotec/pkg/controller/user"
	echo "github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllUsers(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	role, _ := strconv.Atoi(c.QueryParam("role"))

	data, err := controller.GetAllUsers(index, limit, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
