package work_progress

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateWorkProgress(workProgress *entity.WorkProgress) (*entity.WorkProgress, error) {
	id := xid.New()
	currentTime := time.Now()
	workProgress.CreatedAt = &currentTime
	workProgress.UpdatedAt = &currentTime
	workProgress.Id = id.String()
	db := env.MongoDBConnection
	_, err := db.Collection("WorkProgress").UpdateOne(context.Background(), bson.M{"_id": workProgress.Id}, workProgress)
	if err != nil {
		return nil, err
	}
	return workProgress, nil
}
