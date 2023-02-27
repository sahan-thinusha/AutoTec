package vehicle

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetVehiclesByCustomer(index, limit int, userId string) ([]*entity.Vehicle, error) {
	var vehicles []*entity.Vehicle
	vehicles = []*entity.Vehicle{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("Vehicle").Find(context.Background(), bson.M{"userId": userId}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &vehicles); err != nil {
			return nil, err
		}

		return vehicles, nil
	} else {
		cursor, err := db.Collection("Vehicle").Find(context.Background(), bson.M{"userId": userId})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &vehicles); err != nil {
			return nil, err
		}

		return vehicles, nil
	}
}
