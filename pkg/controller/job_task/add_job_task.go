package job_task

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddNewJobTask(job *entity.JobTask) (*entity.JobTask, error) {
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	job.Id = primitive.NewObjectID().Hex()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("JobTask").InsertOne(context.Background(), job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
