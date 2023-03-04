package entity

type Job struct {
	Id          string `json:"id" bson:"_id"`
	Base        `bson:",inline"`
	Subject     string `json:"subject" bson:"subject"`
	Description string `json:"description" bson:"description"`
	LabourTime  int64  `json:"labour_time" bson:"labour_time"`
	CustomerId  string `json:"customerId" bson:"customerId"`
}
