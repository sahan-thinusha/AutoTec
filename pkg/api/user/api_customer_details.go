package user

import (
	controller "autotec/pkg/controller/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetCustomerDetails(c echo.Context) error {

	uid := c.QueryParam("uid")

	data, err := controller.GetCustomerDetails(uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
