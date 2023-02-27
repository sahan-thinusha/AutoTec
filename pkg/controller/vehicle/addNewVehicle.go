package vehicle

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"time"
)

func AddNewVehicle(vehicle *entity.Vehicle) (*entity.Vehicle, error) {
	id := xid.New()
	currentTime := time.Now()
	vehicle.CreatedAt = &currentTime
	vehicle.UpdatedAt = &currentTime
	vehicle.Id = id.String()
	db := env.MongoDBConnection
	_, err := db.Collection("Vehicle").InsertOne(context.Background(), vehicle)
	if err != nil {
		return nil, err
	}
	return vehicle, nil
}
