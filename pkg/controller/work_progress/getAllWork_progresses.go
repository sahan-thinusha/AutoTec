package work_progress

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllWorkProgress(index, limit int) ([]*entity.WorkProgress, error) {
	var workProgress []*entity.WorkProgress
	workProgress = []*entity.WorkProgress{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("WorkProgress").Find(context.Background(), bson.M{}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &workProgress); err != nil {
			return nil, err
		}

		return workProgress, nil
	} else {
		cursor, err := db.Collection("Users").Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &workProgress); err != nil {
			return nil, err
		}

		return workProgress, nil
	}
}
