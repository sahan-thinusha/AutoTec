package entity

type User struct {
	Id        string `json:"id" bson:"_id"`
	Base      `bson:",inline"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	UserName  string `json:"userName" bson:"userName"`
	ContactNo string `json:"contactNo" bson:"contactNo"`
	Role      string `json:"role" bson:"role"`
	Password  string `json:"password" bson:"password"`
}
