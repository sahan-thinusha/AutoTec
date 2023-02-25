package user

import (
	"autotec/pkg/dto"
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strings"
	"time"
)

/**
Fetch user data
*/
func Login(user *dto.UserDto) (*dto.TokenDto, error) {
	fetchedUser := entity.User{}
	db := env.MongoDBConnection
	coll := db.Collection("Users").FindOne(context.Background(), bson.M{"userName": user.UserName, "password": user.Password})
	err := coll.Decode(&fetchedUser)
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(fetchedUser.UserName, "") {
		data, err := createToken(fetchedUser.Roles, fetchedUser.Id)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		return nil, errors.New("invalid user name or password")
	}
}

/**
Jwt Token creation Function
*/
func createToken(roles []string, uid string) (*dto.TokenDto, error) {

	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(env.SigningPrivateKey))
	if err != nil {
		log.Println(err)
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)

	var rolesArr []string

	for _, role := range roles {
		rolesArr = append(roles, role)
	}

	userId := ""
	if !strings.EqualFold(uid, "") {
		userId = uid
	}

	claims["roles"] = rolesArr
	claims["uid"] = userId
	exp := time.Now().Add(time.Hour * 24).Unix()
	claims["exp"] = exp

	tokenString, err := token.SignedString(key)

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
