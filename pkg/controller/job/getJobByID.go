package job

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetJobByID(index, limit int, id string) ([]*entity.JobDetails, error) {
	var job []*entity.JobDetails
	job = []*entity.JobDetails{}
	ctx := context.Background()
	db := env.MongoDBConnection

	matchStage := bson.D{{"$match", bson.D{{"Id", id}}}}
	lookupStage := bson.D{{"$lookup", bson.M{"from": "JobTask", "localField": "_id", "foreignField": "JobID", "as": "jobTask"}}}

	if index >= 0 && limit >= 0 {
		skipStage := bson.D{{"$skip", index}}
		limitStage := bson.D{{"$limit", limit}}
		pipeLine := mongo.Pipeline{matchStage, lookupStage, skipStage, limitStage}
		cursor, err := db.Collection("Job").Aggregate(context.Background(), pipeLine)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	} else {
		pipeLine := mongo.Pipeline{matchStage, lookupStage}
		cursor, err := db.Collection("Job").Aggregate(context.Background(), pipeLine)
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
