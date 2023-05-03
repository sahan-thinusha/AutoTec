package job_task

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateJobTask(job *entity.JobTaskUpdate) (*entity.JobTaskUpdate, error) {
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("JobTask").UpdateOne(context.Background(), bson.M{"_id": job.Id}, bson.M{"$set": job})
	if err != nil {
		return nil, err
	}
	return job, nil
}
