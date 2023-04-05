package user

import (
	"autotec/pkg/dto"
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"autotec/pkg/util"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
)

/*
*
Fetch user data
*/
func Login(user *dto.UserDto) (*dto.TokenDto, error) {
	fetchedUser := entity.User{}
	db := env.MongoDBConnection
	coll := db.Collection("Users").FindOne(context.Background(), bson.M{"userName": user.UserName})
	err := coll.Decode(&fetchedUser)
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(fetchedUser.UserName, "") {
		pw, e := util.Decrypt(fetchedUser.Password)
		if e != nil {
			return nil, e
		}
		if strings.EqualFold(pw, user.Password) {
			data, err := createToken(fetchedUser.Role, fetchedUser.Id)
			if err != nil {
				return nil, err
			}
			return data, nil
		} else {
			return nil, errors.New("invalid user name or password")
		}
	} else {
		return nil, errors.New("invalid user name or password")
	}
}

/*
*
Jwt Token creation Function
*/
func createToken(role string, uid string) (*dto.TokenDto, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	userId := ""
	if !strings.EqualFold(uid, "") {
		userId = uid
	}

	claims["role"] = role
	claims["uid"] = userId
	exp := time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = exp

	tokenString, err := token.SignedString([]byte(env.SigningKey))

	if err != nil {
		fmt.Errorf("something Went Wrong: %s", err.Error())
		return nil, err
	}
	response := dto.TokenDto{}
	response.Token = tokenString
	response.Token_type = "Bearer"
	response.Expiresin = exp

	return &response, nil
}
