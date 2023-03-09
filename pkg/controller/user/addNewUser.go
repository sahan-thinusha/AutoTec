package user

import (
	"autotec/pkg/dto"
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"autotec/pkg/util"
	"context"
	"github.com/rs/xid"
	"time"
)

func AddNewUser(user *entity.User) (*dto.APIResponse, error) {
	id := xid.New()
	currentTime := time.Now()
	user.CreatedAt = &currentTime
	user.UpdatedAt = &currentTime
	user.Id = id.String()
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
