package job_task

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
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

	if job.LabourTime > 0 {
		_, err = db.Collection("Job").UpdateOne(
			context.Background(),
			bson.M{"_id": job.JobID, "status": "TODO"},
			bson.M{"$set": bson.M{"status": "In Progress", "updated_at": time.Now()}},
		)

		if err != nil {
			return nil, err
		}

	}
	ctx := context.Background()
	curser, err := db.Collection("JobTask").Find(ctx, bson.M{"jobID": job.JobID})
	if err != nil {
		return nil, err
	}

	isJobDone := true
	for curser.Next(ctx) {
		job := entity.JobTask{}
		curser.Decode(&job)
		if !strings.EqualFold(job.Status, "Finished") {
			isJobDone = false
			break
		}
	}

	if isJobDone {
		_, err := db.Collection("Job").UpdateOne(context.Background(), bson.M{"_id": job.JobID}, bson.M{"$set": bson.M{"status": "Finished", "updated_at": time.Now()}})
		if err != nil {
			return nil, err
		}
	}

	return job, nil
}
