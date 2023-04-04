package job

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateJob(job *entity.Job) (*entity.Job, error) {
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	db := env.MongoDBConnection
	_, err := db.Collection("Job").UpdateOne(context.Background(), bson.M{"_id": job.Id}, job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
