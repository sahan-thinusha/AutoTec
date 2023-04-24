package promotion

import (
	controller "autotec/pkg/controller/promotion"
	"autotec/pkg/entity"
	echo "github.com/labstack/echo/v4"
	"net/http"
)

func AddNewPromotion(c echo.Context) error {
	promo := entity.Promotion{}
	if error := c.Bind(&promo); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.AddNewPromotion(&promo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
