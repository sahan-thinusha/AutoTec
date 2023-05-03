package entity

import "time"

type JobTask struct {
	Id            string `json:"id" bson:"_id"`
	Base          `bson:",inline"`
	JobID         string    `json:"jobID" bson:"jobID"`
	Subject       string    `json:"subject" bson:"subject"`
	Description   string    `json:"description" bson:"description"`
	StartedAt     time.Time `json:"startedAt" bson:"startedAt"`
	EndedAt       time.Time `json:"endedAt" bson:"endedAt"`
	LabourTime    float64   `json:"labour_time" bson:"labour_time"`
	EstimatedTime float64   `json:"estimatedTime" bson:"estimatedTime"`
	CustomerId    string    `json:"customerId" bson:"customerId"`
	CustomerName  string    `json:"customerName" bson:"customerName"`
	LabourID      string    `json:"labourID" bson:"labourID"`
	LabourName    string    `json:"labourName" bson:"labourName"`
	VehicleId     string    `json:"vehicleId" bson:"vehicleId"`
	VehicleName   string    `json:"vehicleName" bson:"vehicleName"`
	Status        string    `json:"status" bson:"status"`
	LabourRate    float64   `json:"labourRate" bson:"labourRate"`
}

type JobTaskUpdate struct {
	Id            string `json:"id" bson:"_id"`
	Base          `bson:",inline"`
	JobID         string  `json:"jobID" bson:"jobID"`
	Subject       string  `json:"subject" bson:"subject"`
	Description   string  `json:"description" bson:"description"`
	LabourTime    float64 `json:"labour_time" bson:"labour_time"`
	EstimatedTime float64 `json:"estimatedTime" bson:"estimatedTime"`
	CustomerId    string  `json:"customerId" bson:"customerId"`
	CustomerName  string  `json:"customerName" bson:"customerName"`
	LabourID      string  `json:"labourID" bson:"labourID"`
	LabourName    string  `json:"labourName" bson:"labourName"`
	VehicleId     string  `json:"vehicleId" bson:"vehicleId"`
	VehicleName   string  `json:"vehicleName" bson:"vehicleName"`
	Status        string  `json:"status" bson:"status"`
	LabourRate    float64 `json:"labourRate" bson:"labourRate"`
}
