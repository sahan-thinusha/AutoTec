package user

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetCustomerDetails(uid string) (*entity.CustomerDetails, error) {
	customerDetails := entity.CustomerDetails{}
	ctx := context.Background()
	db := env.MongoDBConnection
	fetchedUser := entity.User{}
	coll := db.Collection("Users").FindOne(context.Background(), bson.M{"id": uid})
	err := coll.Decode(&fetchedUser)
	customerDetails.CustomerId = fetchedUser.Id
	customerDetails.CustomerName = fetchedUser.FirstName + " " + fetchedUser.LastName
	if err != nil {
		return nil, err
	}

	now := time.Now()
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	endOfYear := time.Date(now.Year(), 12, 31, 23, 59, 59, 999999999, time.Local)
	matchStage := bson.D{{"$match", bson.M{"date": bson.M{"$gte": startOfYear, "$lte": endOfYear}, "customerId": uid}}}
	groupStage := bson.D{{"$group", bson.M{"_id": bson.M{"year": bson.M{"$year": "$date"}}, "count": bson.M{"$sum": 1}, "last_job_date": bson.D{{"$max", "$date"}}}}}
	pipeline := mongo.Pipeline{matchStage, groupStage}
	cursor, err := db.Collection("Job").Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result struct {
			Count       int64     `bson:"count"`
			LastJobDate time.Time `bson:"last_job_date"`
		}
		if err := cursor.Decode(&result); err != nil {
			return &customerDetails, err
		}

		customerDetails.LastSaleDate = result.LastJobDate
		customerDetails.TotalJobs = result.Count

	}

	matchStage1 := bson.D{{"$match", bson.M{"status": "Finished"}}}
	lookupStage := bson.D{{"$lookup", bson.M{"from": "JobTask", "localField": "_id", "foreignField": "jobID", "as": "jobTask"}}}
	unwindStage1 := bson.D{{"$unwind", bson.M{"path": "$jobTask", "preserveNullAndEmptyArrays": true}}}

	setStage := bson.D{{
		"$set", bson.M{
			"hoursSold": bson.M{
				"$multiply": []string{"$jobTask.labourRate", "$jobTask.labour_time"},
			},
		},
	}}
	groupStage1 := bson.D{{"$group", bson.D{
		{"_id", nil},
		{"totalHoursSold", bson.M{"$sum": "$hoursSold"}},
	}}}
	pipeline1 := mongo.Pipeline{matchStage, matchStage1, lookupStage, unwindStage1, setStage, groupStage1}
	cursor1, err := db.Collection("Job").Aggregate(context.Background(), pipeline1)
	if err != nil {
		return nil, err
	}
	defer cursor1.Close(ctx)

	if err != nil {
		return nil, err
	}
	defer cursor1.Close(ctx)

	if cursor1.Next(ctx) {
		data := entity.EmployeeEfficiencyFetch{}
		cursor.Decode(&data)
		customerDetails.TotalSpends = data.TotalHoursSold
	}

	if cursor1.Next(ctx) {
		var result struct {
			Count       int64     `bson:"count"`
			LastJobDate time.Time `bson:"last_job_date"`
		}
		if err := cursor.Decode(&result); err != nil {
			return &customerDetails, err
		}

		customerDetails.LastSaleDate = result.LastJobDate
		customerDetails.TotalJobs = result.Count

	}

	durationSinceLastSale := time.Since(customerDetails.LastSaleDate)

	type rfmst struct {
		R int
		F int
		M int
	}
	rfm := rfmst{}
	if durationSinceLastSale < time.Hour*24*7 {
		rfm.R = 5
	} else if durationSinceLastSale < time.Hour*24*30 {
		rfm.R = 4
	} else if durationSinceLastSale < time.Hour*24*90 {
		rfm.R = 3
	} else if durationSinceLastSale < time.Hour*24*180 {
		rfm.R = 2
	} else {
		rfm.R = 1
	}

	if customerDetails.TotalJobs >= 100 {
		rfm.F = 5
	} else if customerDetails.TotalJobs >= 75 {
		rfm.F = 4
	} else if customerDetails.TotalJobs >= 50 {
		rfm.F = 3
	} else if customerDetails.TotalJobs >= 25 {
		rfm.F = 2
	} else {
		rfm.F = 1
	}

	if customerDetails.TotalSpends >= 1000000.0 {
		rfm.M = 5
	} else if customerDetails.TotalSpends >= 500000.0 {
		rfm.M = 4
	} else if customerDetails.TotalSpends >= 100000.0 {
		rfm.M = 3
	} else if customerDetails.TotalSpends >= 50000.0 {
		rfm.M = 2
	} else {
		rfm.M = 1
	}

	customerDetails.RFMScore = float64((rfm.F + rfm.R + rfm.M) / 3.0)

	return &customerDetails, nil
}
