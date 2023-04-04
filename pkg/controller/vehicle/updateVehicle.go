package vehicle

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateVehicle(vehicle *entity.Vehicle) (*entity.Vehicle, error) {
	currentTime := time.Now()
	vehicle.CreatedAt = &currentTime
	vehicle.UpdatedAt = &currentTime
	db := env.MongoDBConnection
	_, err := db.Collection("Vehicle").UpdateOne(context.Background(), bson.M{"_id": vehicle.Id}, vehicle)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}
