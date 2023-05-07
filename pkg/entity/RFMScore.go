package entity

import "time"

type CustomerDetails struct {
	Id           string `json:"id" bson:"_id"`
	Base         `bson:",inline"`
	RFMScore     float64   `json:"rfm_score" bson:"rfm_score"`
	CustomerId   string    `json:"customerId" bson:"customerId"`
	CustomerName string    `json:"customerName" bson:"customerName"`
	LastSaleDate time.Time `json:"lastSaleDate" bson:"lastSaleDate"`
	TotalSpends  float64   `json:"totalSpends" bson:"totalSpends"`
	TotalJobs    int64     `json:"totalJobs" bson:"totalJobs"`
	ContactNo    string    `json:"contactNo" bson:"contactNo"`
}
