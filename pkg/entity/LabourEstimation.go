package entity

type LabourEstimateRequest struct {
	CarModel string `json:"car_model"`
	Part     string `json:"part"`
	Make     string `json:"make"`
}

type LabourEstimateResponse struct {
	Estimated_repair_time float64 `json:"estimated_repair_time"`
}
