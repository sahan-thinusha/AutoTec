package entity

type Job struct {
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	Subject      string  `json:"subject" bson:"subject"`
	Description  string  `json:"description" bson:"description"`
	LabourTime   int64   `json:"labour_time" bson:"labour_time"`
	CustomerId   string  `json:"customerId" bson:"customerId"`
	CustomerName string  `json:"customerName" bson:"customerName"`
	Status       string  `json:"status" bson:"status"`
	LabourRate   float64 `json:"labourRate" bson:"labourRate"`
}

type JobDetails struct {
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	Subject      string     `json:"subject" bson:"subject"`
	Description  string     `json:"description" bson:"description"`
	LabourTime   int64      `json:"labour_time" bson:"labour_time"`
	CustomerId   string     `json:"customerId" bson:"customerId"`
	CustomerName string     `json:"customerName" bson:"customerName"`
	Status       string     `json:"status" bson:"status"`
	JobTask      []*JobTask `json:"jobTask" bson:"jobTask"`
	Customer     string     `json:"customer" bson:"customer"`
	VehicleID    string     `json:"vehicleID" bson:"vehicleID"`
	Vehicle      string     `json:"vehicle" bson:"vehicle"`
	LabourRate   float64    `json:"labourRate" bson:"labourRate"`
}
