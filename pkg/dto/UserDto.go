package dto

type UserDto struct {
	UserName string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
