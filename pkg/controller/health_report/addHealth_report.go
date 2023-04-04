package health_report

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddrNewHealthReport(healthReport *entity.HealthReport) (*entity.HealthReport, error) {
	currentTime := time.Now()
	healthReport.CreatedAt = &currentTime
	healthReport.UpdatedAt = &currentTime
	healthReport.Id = primitive.NewObjectID().Hex()
	var e error
	if e != nil {
		return nil, e
	}
	db := env.MongoDBConnection
	_, err := db.Collection("HealthReport").InsertOne(context.Background(), healthReport)
	if err != nil {
		return nil, err
	}
	return healthReport, nil
}
