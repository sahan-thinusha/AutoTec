package dto

type UserDto struct {
	UserName string `json:"userName" bson:"userName"`
	Password string `json:"password" bson:"password"`
}
