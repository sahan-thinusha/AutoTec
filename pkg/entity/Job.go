package entity

type Job struct {
<<<<<<< HEAD
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	Subject      string `json:"subject" bson:"subject"`
	Description  string `json:"description" bson:"description"`
	LabourTime   int64  `json:"labour_time" bson:"labour_time"`
	CustomerId   string `json:"customerId" bson:"customerId"`
	CustomerName string `json:"customerName" bson:"customerName"`
	Status       string `json:"status" bson:"status"`
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
=======
	Id          string `json:"id" bson:"_id"`
	Base        `bson:",inline"`
	Subject     string `json:"subject" bson:"subject"`
	Description string `json:"description" bson:"description"`
	LabourTime  int64  `json:"labour_time" bson:"labour_time"`
	CustomerId  string `json:"customerId" bson:"customerId"`
	Customer    string `json:"customer" bson:"customer"`
	VehicleID   string `json:"vehicleID" bson:"vehicleID"`
	Vehicle     string `json:"vehicle" bson:"vehicle"`
>>>>>>> d11b8fd (changes)
}
