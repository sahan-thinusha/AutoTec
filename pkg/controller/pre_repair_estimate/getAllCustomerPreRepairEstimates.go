package pre_repair_estimate

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllPreRepairEstimateForCustomer(index, limit int, uid string) ([]*entity.PreRepairEstimate, error) {
	var preRepairEstimate []*entity.PreRepairEstimate
	preRepairEstimate = []*entity.PreRepairEstimate{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("PreRepairEstimate").Find(context.Background(), bson.M{"customerId": uid, "status": bson.M{"$in": []string{env.CONFIRMED, env.REJECTED, env.APPROVED}}}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &preRepairEstimate); err != nil {
			return nil, err
		}

		return preRepairEstimate, nil
	} else {
		cursor, err := db.Collection("PreRepairEstimate").Find(context.Background(), bson.M{"customerId": uid, "status": bson.M{"$in": []string{env.CONFIRMED, env.REJECTED, env.APPROVED}}})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &preRepairEstimate); err != nil {
			return nil, err
		}

		return preRepairEstimate, nil
	}
}
