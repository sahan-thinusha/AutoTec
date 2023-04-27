package efficiency

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTechnicianEfficiency(index, limit int) ([]*entity.LabourEfficiency, error) {
	var job []*entity.LabourEfficiency
	job = []*entity.LabourEfficiency{}
	ctx := context.Background()
	db := env.MongoDBConnection

	matchStage := bson.D{{"$match", bson.D{{"role", "TECHNICIAN"}}}}
	lookupStage := bson.D{{"$lookup", bson.M{"from": "JobTask", "localField": "_id", "foreignField": "labourID", "as": "jobTask"}}}
	unwindStage1 := bson.D{{"$unwind", bson.M{"path": "$jobTask", "preserveNullAndEmptyArrays": true}}}
	setStage := bson.D{{
		"$set", bson.M{
			"HoursSold": bson.M{
				"$multiply": []string{"$jobTask.labourRate", "$jobTask.labour_time"},
			},
		},
	}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"totalHoursEstimated", bson.M{"$sum": "$jobTask.estimatedTime"}},
		{"totalHoursWorked", bson.M{"$sum": "$jobTask.labour_time"}},
		{"totalHoursSold", bson.M{"$sum": "$hoursSold"}},
		{"firstName", bson.D{{"$first", "$firstName"}}},
		{"lastName", bson.D{{"$first", "$lastName"}}},
	}}}

	projectStage := bson.D{{
		"$project", bson.D{
			{"employeeId", "$_id"},
			{"firstName", "$firstName"},
			{"lastName", "$lastName"},
			{"laborUtilization", bson.D{
				{"$multiply", bson.A{
					bson.D{{"$cond", bson.A{
						bson.D{{"$eq", bson.A{"$hoursattended", 0}}},
						0,
						bson.D{{"$round", bson.A{bson.D{{"$multiply", bson.A{bson.D{{"$divide", bson.A{"$totalHoursWorked", "$c"}}}, 100}}}, 2}}}}}},
					100,
				}},
			}},
		}}}
	if index >= 0 && limit >= 0 {
		skipStage := bson.D{{"$skip", index}}
		limitStage := bson.D{{"$limit", limit}}
		pipeLine := mongo.Pipeline{matchStage, lookupStage, skipStage, limitStage, unwindStage1, setStage, groupStage, projectStage}
		cursor, err := db.Collection("Job").Aggregate(context.Background(), pipeLine)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	} else {
		pipeLine := mongo.Pipeline{matchStage, lookupStage, unwindStage1, setStage, groupStage, projectStage}
		cursor, err := db.Collection("Job").Aggregate(context.Background(), pipeLine)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		if err = cursor.All(context.Background(), &job); err != nil {
			return nil, err
		}

		return job, nil
	}
}
