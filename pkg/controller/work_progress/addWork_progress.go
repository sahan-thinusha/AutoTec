package work_progress

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddrNewWorkProgress(workProgress *entity.WorkProgress) (*entity.WorkProgress, error) {
	currentTime := time.Now()
	workProgress.CreatedAt = &currentTime
	workProgress.UpdatedAt = &currentTime
	workProgress.Id = primitive.NewObjectID().Hex()
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
