package user

import (
	controller "autotec/pkg/controller/user"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func AddNewUser(c echo.Context) error {

	user := entity.User{}
	if error := c.Bind(&user); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.AddNewUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
