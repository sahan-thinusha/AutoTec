package user

import (
	controller "autotec/pkg/controller/user"
	"autotec/pkg/dto"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func Login(c echo.Context) error {
	user := dto.UserDto{}
	if error := c.Bind(&user); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}

	data, err := controller.Login(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
