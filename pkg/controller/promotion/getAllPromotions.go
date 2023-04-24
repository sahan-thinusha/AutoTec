package promotion

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllPromotions(index, limit int) ([]*entity.Promotion, error) {
	var promotion []*entity.Promotion
	promotion = []*entity.Promotion{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("Promotion").Find(context.Background(), bson.M{}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &promotion); err != nil {
			return nil, err
		}

		return promotion, nil
	} else {
		cursor, err := db.Collection("Promotion").Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &promotion); err != nil {
			return nil, err
		}

		return promotion, nil
	}
}
