package pre_repair_estimate

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddrNewPreRepairEstimate(preRepairEstimate *entity.PreRepairEstimate) (*entity.PreRepairEstimate, error) {
	currentTime := time.Now()
	preRepairEstimate.CreatedAt = &currentTime
	preRepairEstimate.UpdatedAt = &currentTime
	preRepairEstimate.Id = primitive.NewObjectID().Hex()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("PreRepairEstimate").InsertOne(context.Background(), preRepairEstimate)
	if err != nil {
		return nil, err
	}
	return preRepairEstimate, nil
}
