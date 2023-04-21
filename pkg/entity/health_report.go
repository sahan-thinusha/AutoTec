package entity

import "time"

type HealthReport struct {
	Id               string `json:"id" bson:"_id"`
	Base             `bson:",inline"`
	Date             time.Time          `json:"date" bson:"date"`
	VehicleId        string             `json:"vehicleId" bson:"vehicleId"`
	VehicleName      string             `json:"vehicleName" bson:"vehicleName"`
	CustomerId       string             `json:"customerId" bson:"customerId"`
	CustomerName     string             `json:"customerName" bson:"customerName"`
	HealthReportItem []HealthReportItem `json:"healthReportItem" bson:"healthReportItem"`
}

type HealthReportItem struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Condition int    `json:"condition" bson:"condition"`
}
