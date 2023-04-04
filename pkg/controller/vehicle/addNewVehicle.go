package vehicle

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddNewVehicle(vehicle *entity.Vehicle) (*entity.Vehicle, error) {
	currentTime := time.Now()
	vehicle.CreatedAt = &currentTime
	vehicle.UpdatedAt = &currentTime
	vehicle.Id = primitive.NewObjectID().Hex()
	db := env.MongoDBConnection
	_, err := db.Collection("Vehicle").InsertOne(context.Background(), vehicle)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}
