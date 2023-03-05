package health_report

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"github.com/rs/xid"
	"time"
)

func AddrNewHealthReport(healthReport *entity.HealthReport) (*entity.HealthReport, error) {
	id := xid.New()
	currentTime := time.Now()
	healthReport.CreatedAt = &currentTime
	healthReport.UpdatedAt = &currentTime
	healthReport.Id = id.String()
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
