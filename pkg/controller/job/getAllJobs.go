package job

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllJob(index, limit int) ([]*entity.Job, error) {
	var job []*entity.Job
	job = []*entity.Job{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("Job").Find(context.Background(), bson.M{}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	} else {
		cursor, err := db.Collection("Job").Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	}
}

func GetAllCustomerJob(index, limit int, cid string) ([]*entity.Job, error) {
	var job []*entity.Job
	job = []*entity.Job{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("Job").Find(context.Background(), bson.M{"customerId": cid}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	} else {
		cursor, err := db.Collection("Job").Find(context.Background(), bson.M{"customerId": cid})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	}
}
