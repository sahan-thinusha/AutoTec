package health_report

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllHealthReportTemplates() ([]*entity.HealthReportTemplate, error) {
	var healthReport []*entity.HealthReportTemplate
	healthReport = []*entity.HealthReportTemplate{}
	ctx := context.Background()
	db := env.MongoDBConnection

	cursor, err := db.Collection("HealthReportTemplate").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	if err = cursor.All(context.Background(), &healthReport); err != nil {
		return nil, err
	}

	return healthReport, nil

}

