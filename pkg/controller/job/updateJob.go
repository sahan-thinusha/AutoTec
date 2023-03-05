package job

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateJob(job *entity.Job) (*entity.Job, error) {
	id := xid.New()
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	job.Id = id.String()
	db := env.MongoDBConnection
	_, err := db.Collection("Job").UpdateOne(context.Background(), bson.M{"_id": job.Id}, job)
	if err != nil {
		return nil, err
	}
	return job, nil
}
