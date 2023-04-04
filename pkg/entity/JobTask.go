package entity

import "time"

type JobTask struct {
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	JobID        string    `json:"jobID" bson:"jobID"`
	Subject      string    `json:"subject" bson:"subject"`
	Description  string    `json:"description" bson:"description"`
	StartedAt    time.Time `json:"startedAt" bson:"startedAt"`
	EndedAt      time.Time `json:"endedAt" bson:"endedAt"`
	LabourTime   int64     `json:"labour_time" bson:"labour_time"`
	CustomerId   string    `json:"customerId" bson:"customerId"`
	CustomerName string    `json:"customerName" bson:"customerName"`
	LabourID     string    `json:"labourID" bson:"labourID"`
	LabourName   string    `json:"labourName" bson:"labourName"`
	Status       string    `json:"status" bson:"status"`
}
