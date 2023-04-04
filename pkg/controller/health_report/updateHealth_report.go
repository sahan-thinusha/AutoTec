package health_report

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateHealthReport(healthReport *entity.HealthReport) (*entity.HealthReport, error) {
	currentTime := time.Now()
	healthReport.CreatedAt = &currentTime
	healthReport.UpdatedAt = &currentTime
	db := env.MongoDBConnection
	_, err := db.Collection("HealthReport").UpdateOne(context.Background(), bson.M{"_id": healthReport.Id}, healthReport)
	if err != nil {
		return nil, err
	}
	return healthReport, nil
}
