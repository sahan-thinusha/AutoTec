package entity

type PreRepairEstimate struct {
	Id                       string `json:"id" bson:"_id"`
	Base                     `bson:",inline"`
	Subject                  string                     `json:"subject" bson:"subject"`
	Number                   string                     `json:"number" bson:"number"`
	Description              string                     `json:"description" bson:"description"`
	CustomerId               string                     `json:"customerId" bson:"customerId"`
	PreRepairEstimateDetails []PreRepairEstimateDetails `json:"preRepairEstimateDetails" bson:"preRepairEstimateDetails"`
}

type PreRepairEstimateDetails struct {
	Quantity float64 `json:"quantity" bson:"quantity"`
	Price    float64 `json:"price" bson:"price"`
}
