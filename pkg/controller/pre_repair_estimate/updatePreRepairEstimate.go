package pre_repair_estimate

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdatePreRepairEstimate(preRepairEstimate *entity.PreRepairEstimate) (*entity.PreRepairEstimate, error) {
	id := xid.New()
	currentTime := time.Now()
	preRepairEstimate.CreatedAt = &currentTime
	preRepairEstimate.UpdatedAt = &currentTime
	preRepairEstimate.Id = id.String()
	db := env.MongoDBConnection
	_, err := db.Collection("PreRepairEstimate").UpdateOne(context.Background(), bson.M{"_id": preRepairEstimate.Id}, preRepairEstimate)
	if err != nil {
		return nil, err
	}
	return preRepairEstimate, nil
}
