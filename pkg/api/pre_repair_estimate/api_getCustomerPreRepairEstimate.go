package pre_repair_estimate

import (
	controller "autotec/pkg/controller/pre_repair_estimate"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetAllCustomerPreRepairEstimates(c echo.Context) error {

	index, e := strconv.Atoi(c.QueryParam("index"))
	if e != nil {
		index = -1
	}
	limit, e := strconv.Atoi(c.QueryParam("limit"))
	if e != nil {
		limit = -1
	}

	uid := c.QueryParam("uid")

	data, err := controller.GetAllPreRepairEstimateForCustomer(index, limit, uid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, data)

}
