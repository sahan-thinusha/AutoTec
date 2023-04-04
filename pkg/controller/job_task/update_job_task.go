package job_task

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateJobTask(job *entity.JobTask) (*entity.JobTask, error) {
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("JobTask").UpdateOne(context.Background(), bson.M{"_id": job.Id}, job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
