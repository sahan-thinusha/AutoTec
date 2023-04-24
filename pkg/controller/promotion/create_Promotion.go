package promotion

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddNewPromotion(promotion *entity.Promotion) (*entity.Promotion, error) {
	currentTime := time.Now()
	promotion.CreatedAt = &currentTime
	promotion.UpdatedAt = &currentTime
	promotion.Id = primitive.NewObjectID().Hex()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("Promotion").InsertOne(context.Background(), promotion)
	if err != nil {
		return nil, err
	}
	return promotion, nil
}
