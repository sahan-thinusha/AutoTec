package job

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddrNewJob(job *entity.Job) (*entity.Job, error) {
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	job.Id = primitive.NewObjectID().Hex()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("Job").InsertOne(context.Background(), job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
