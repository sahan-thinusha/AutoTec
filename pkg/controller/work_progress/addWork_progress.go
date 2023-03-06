package work_progress

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"time"
)

func AddrNewWorkProgress(workProgress *entity.WorkProgress) (*entity.WorkProgress, error) {
	id := xid.New()
	currentTime := time.Now()
	workProgress.CreatedAt = &currentTime
	workProgress.UpdatedAt = &currentTime
	workProgress.Id = id.String()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("WorkProgress").InsertOne(context.Background(), workProgress)
	if err != nil {
		return nil, err
	}
	return workProgress, nil
}
