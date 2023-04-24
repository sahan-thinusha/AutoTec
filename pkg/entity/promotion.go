package entity

import "time"

type Promotion struct {
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	Subject      string    `json:"subject" bson:"subject"`
	Description  string    `json:"description" bson:"description"`
	CustomerID   string    `json:"customerID" bson:"customerID"`
	CustomerName string    `json:"customerName" bson:"customerName"`
	Amount       float64   `json:"amount" bson:"amount"`
	CouponCode   string    `json:"couponCode" bson:"couponCode"`
	StartAt      time.Time `json:"startAt" bson:"startAt"`
	EndsAt       time.Time `json:"endsAt" bson:"endsAt"`
}
