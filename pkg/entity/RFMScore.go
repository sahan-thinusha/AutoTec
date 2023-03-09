package entity

type RFMScore struct {
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	RFMScore     float64 `json:"rfm_score" bson:"rfm_score"`
	CustomerId   string  `json:"customerId" bson:"customerId"`
	CustomerName string  `json:"customerName" bson:"customerName"`
}
