package job_task

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllJobTasksByStatusAndEmployee(index, limit int, uid, status string) ([]*entity.JobTask, error) {
	var tasks []*entity.JobTask
	tasks = []*entity.JobTask{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("JobTask").Find(context.Background(), bson.M{"labourID": uid, "status": status}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &tasks); err != nil {
			return nil, err
		}

		return tasks, nil
	} else {
		cursor, err := db.Collection("JobTask").Find(context.Background(), bson.M{"labourID": uid, "status": status})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &tasks); err != nil {
			return nil, err
		}

		return tasks, nil
	}
}

func GetAllJobTasksByJob(jobId string) ([]*entity.JobTask, error) {
	var tasks []*entity.JobTask
	tasks = []*entity.JobTask{}
	ctx := context.Background()
	db := env.MongoDBConnection

	cursor, err := db.Collection("JobTask").Find(context.Background(), bson.M{"jobID": jobId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil

}
