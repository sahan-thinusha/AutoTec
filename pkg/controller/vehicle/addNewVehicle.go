package vehicle

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddNewVehicle(vehicle *entity.Vehicle) (*entity.Vehicle, error) {
	currentTime := time.Now()
	vehicle.CreatedAt = &currentTime
	vehicle.UpdatedAt = &currentTime
	user := entity.User{}

	vehicle.Id = primitive.NewObjectID().Hex()
	db := env.MongoDBConnection

	data := db.Collection("Users").FindOne(context.Background(), bson.M{"_id": vehicle.UserId})
	data.Decode(user)
	vehicle.FirstName = user.FirstName
	vehicle.LastName = user.LastName
	_, err := db.Collection("Vehicle").InsertOne(context.Background(), vehicle)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}
