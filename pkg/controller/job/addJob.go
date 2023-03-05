package job

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"time"
)

func AddrNewJob(job *entity.Job) (*entity.Job, error) {
	id := xid.New()
	currentTime := time.Now()
	job.CreatedAt = &currentTime
	job.UpdatedAt = &currentTime
	job.Id = id.String()
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
