package prediction

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func GetEstPrediction(request entity.LabourEstimateRequest) (*entity.LabourEstimateResponse, error) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Accept", "application/json").
		SetBody(request).
		Post(env.BaseURL)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err

	}

	if resp.IsSuccess() {
		ets := entity.LabourEstimateResponse{}
		json.Unmarshal(resp.Body(), &ets)
		return &ets, nil
	} else {
		return nil, errors.New(resp.Status())
	}
	return nil, nil
}
