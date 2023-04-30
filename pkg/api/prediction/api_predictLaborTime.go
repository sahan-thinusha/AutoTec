package prediction

import (
	controller "autotec/pkg/controller/prediction"
	"autotec/pkg/entity"
	"github.com/labstack/echo/v4"
	"net/http"
)

func PredictLabourTime(c echo.Context) error {
	estimation := entity.LabourEstimateRequest{}
	if error := c.Bind(&estimation); error != nil {
		return c.JSON(http.StatusBadRequest, error.Error())
	}
	data, err := controller.GetEstPrediction(estimation)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)

}
