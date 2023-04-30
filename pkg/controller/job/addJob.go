package job

import (
	"autotec/pkg/controller/job_task"
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
	job.Date = currentTime
	job.Id = primitive.NewObjectID().Hex()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("Job").InsertOne(context.Background(), job)

	for _, v := range job.JobTask {
		v.JobID = job.Id
		job_task.AddNewJobTask(&v)
	}

	if err != nil {
		return nil, err
	}
	return job, nil
}
