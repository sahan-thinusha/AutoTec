package health_report

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllHealthReports(index, limit int) ([]*entity.HealthReport, error) {
	var healthReport []*entity.HealthReport
	healthReport = []*entity.HealthReport{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("HealthReport").Find(context.Background(), bson.M{}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &healthReport); err != nil {
			return nil, err
		}

		return healthReport, nil
	} else {
		cursor, err := db.Collection("HealthReport").Find(context.Background(), bson.M{})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &healthReport); err != nil {
			return nil, err
		}

		return healthReport, nil
	}
}

func GetAllHealthReportsForCustomer(index, limit int, uid string) ([]*entity.HealthReport, error) {
	var healthReport []*entity.HealthReport
	healthReport = []*entity.HealthReport{}
	ctx := context.Background()
	db := env.MongoDBConnection

	if index >= 0 && limit >= 0 {
		offset := index * limit
		opts := options.Find()
		opts = opts.SetLimit(int64(limit))
		opts = opts.SetSkip(int64(offset))
		cursor, err := db.Collection("HealthReport").Find(context.Background(), bson.M{"customerId": uid, "status": env.CONFIRMED}, opts)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &healthReport); err != nil {
			return nil, err
		}

		return healthReport, nil
	} else {
		cursor, err := db.Collection("HealthReport").Find(context.Background(), bson.M{"customerId": uid, "status": env.CONFIRMED})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &healthReport); err != nil {
			return nil, err
		}

		return healthReport, nil
	}
}
