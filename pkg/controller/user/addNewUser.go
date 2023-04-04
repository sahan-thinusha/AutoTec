package user

import (
	"autotec/pkg/dto"
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"autotec/pkg/util"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddNewUser(user *entity.User) (*dto.APIResponse, error) {
	currentTime := time.Now()
	user.CreatedAt = &currentTime
	user.UpdatedAt = &currentTime
	user.Id = primitive.NewObjectID().Hex()
	var e error
	user.Password, e = util.Encrypt(user.Password)
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("Users").InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	response := dto.APIResponse{}
	response.Id = user.Id
	return &response, nil
}
