package promotion

import (
	controller "autotec/pkg/controller/promotion"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllPromotionsByCustomer(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}
	custID := c.QueryParam("limit")

	data, err := controller.GetAllCustomerPromotions(index, limit, custID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
