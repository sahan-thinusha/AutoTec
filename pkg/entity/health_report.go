package entity

import "time"

type HealthReport struct {
	Id               string `json:"id" bson:"_id"`
	Base             `bson:",inline"`
	Date             time.Time          `json:"date" bson:"date"`
	VehicleId        string             `json:"VehicleId" bson:"VehicleId"`
	CustomerId       string             `json:"customerId" bson:"customerId"`
	HealthReportItem []HealthReportItem `json:"healthReportItem" bson:"healthReportItem"`
}

type HealthReportItem struct {
	Part      string `json:"part" bson:"part"`
	Condition string `json:"condition" bson:"condition"`
}
