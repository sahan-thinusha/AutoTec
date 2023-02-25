package user

import (
	"autotec/pkg/dto"
	"autotec/pkg/env"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func Login(user *dto.UserDto) (*dto.UserDto, error) {
	fetchedUser := dto.UserDto{}
	db := env.MongoDBConnection
	coll := db.Collection("Users").FindOne(context.Background(), bson.M{"userName": user.UserName, "password": user.Password})
	err := coll.Decode(&fetchedUser)
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(fetchedUser.UserName, "") {

		return user, nil
	} else {
		return nil, errors.New("invalid user name or password")
	}
}
