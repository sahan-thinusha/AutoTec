package entity

type Vehicle struct {
	Id              string `json:"id" bson:"_id"`
	Base            `bson:",inline"`
	NumberPlate     string `json:"numberPlate" bson:"numberPlate"`
	Make            string `json:"make" bson:"make"`
	Model           string `json:"model" bson:"model"`
	OdometerReading int64  `json:"odometerReading" bson:"odometerReading"`
	UserId          string `json:"userId" bson:"userId"`
	FirstName       string `json:"firstName" bson:"firstName"`
	LastName        string `json:"lastName" bson:"lastName"`
}
