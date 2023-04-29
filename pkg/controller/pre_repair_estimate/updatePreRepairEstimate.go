package pre_repair_estimate

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdatePreRepairEstimate(preRepairEstimate *entity.PreRepairEstimateUpdate) (*entity.PreRepairEstimateUpdate, error) {
	currentTime := time.Now()
	preRepairEstimate.CreatedAt = &currentTime
	preRepairEstimate.UpdatedAt = &currentTime
	db := env.MongoDBConnection
	_, err := db.Collection("PreRepairEstimate").UpdateOne(context.Background(), bson.M{"_id": preRepairEstimate.Id}, bson.M{"$set": preRepairEstimate})
	if err != nil {
		return nil, err
	}
	return preRepairEstimate, nil
}
