package entity

import "time"

type PreRepairEstimate struct {
	Id                       string `json:"id" bson:"_id"`
	Base                     `bson:",inline"`
	Date                     time.Time                  `json:"date" bson:"date"`
	VehicleId                string                     `json:"vehicleId" bson:"vehicleId"`
	VehicleName              string                     `json:"vehicleName" bson:"vehicleName"`
	PreRepairEstimateDetails []PreRepairEstimateDetails `json:"preRepairEstimateDetails" bson:"preRepairEstimateDetails"`
	CustomerId               string                     `json:"customerId" bson:"customerId"`
	CustomerName             string                     `json:"customerName" bson:"customerName"`
	Status                   string                     `json:"status" bson:"status"`
	LabourRate               float64                    `json:"labourRate" bson:"labourRate"`
}

type PreRepairEstimateDetails struct {
	Name     string  `json:"name" bson:"name"`
	Quantity int64   `json:"quantity" bson:"quantity"`
	Price    float64 `json:"price" bson:"price"`
	Note     string  `json:"note" bson:"note"`
}
