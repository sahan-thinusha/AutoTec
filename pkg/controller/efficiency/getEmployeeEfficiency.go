package efficiency

import (
	"autotec/pkg/entity"
	"autotec/pkg/env"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

func GetTechnicianEfficiency() ([]*entity.LabourEfficiency, error) {
	var labourEfficiency []*entity.LabourEfficiency
	labourEfficiency = []*entity.LabourEfficiency{}
	ctx := context.Background()
	db := env.MongoDBConnection

	now := time.Now()
	_, month, _ := now.Date()
	startOfMonth := time.Date(now.Year(), month, 1, 0, 0, 0, 0, time.Local)
	daysInMonth := now.Sub(startOfMonth).Hours() / 24.0
	totalShiftHours := daysInMonth * 8
	endOfMonth := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC)

	matchStageEmp := bson.D{{"$match", bson.D{{"role", "Technician"}}}}
	pipeLine1 := mongo.Pipeline{matchStageEmp}

	cursor1, err := db.Collection("Users").Aggregate(context.Background(), pipeLine1)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}
	defer cursor1.Close(ctx)

	for cursor1.Next(ctx) {
		user := entity.User{}
		err := cursor1.Decode(&user)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		matchStage := bson.D{{"$match", bson.D{{"status", env.Finished}}}}

		matchStageDate := bson.D{{"$match", bson.D{
			{"date", bson.M{
				"$gte": startOfMonth,
				"$lte": endOfMonth,
			},
			},
		},
		}}

		inlinePipeJobTask := bson.A{
			bson.M{
				"$match": bson.D{{"$expr", bson.D{{"$and", bson.A{
					bson.M{"$eq": bson.A{"$jobID", "$$jobId"}}, bson.M{"$eq": bson.A{"$$labourID", "$labourID"}},
				}}}}},
			},
		}
		lookupStage := bson.D{{"$lookup", bson.M{
			"from":     "JobTask",
			"let":      bson.M{"jobId": "$_id", "labourID": user.Id},
			"pipeline": inlinePipeJobTask,
			"as":       "jobTask",
		}}}

		unwindStage1 := bson.D{{"$unwind", bson.M{"path": "$jobTask", "preserveNullAndEmptyArrays": true}}}
		setStage := bson.D{{
			"$set", bson.M{
				"hoursSold": bson.M{
					"$multiply": []string{"$jobTask.labourRate", "$jobTask.labour_time"},
				},
			},
		}}
		groupStage := bson.D{{"$group", bson.D{
			{"_id", "$jobTask.labourID"},
			{"totalHoursEstimated", bson.M{"$sum": "$jobTask.estimatedTime"}},
			{"totalHoursWorked", bson.M{"$sum": "$jobTask.labour_time"}},
			{"totalHoursSold", bson.M{"$sum": "$hoursSold"}},
		}}}

		pipeLine := mongo.Pipeline{matchStage, matchStageDate, lookupStage, unwindStage1, setStage, groupStage}
		cursor, err := db.Collection("Job").Aggregate(context.Background(), pipeLine)
		if err != nil {
			fmt.Println("s", err)
			return nil, err
		}
		defer cursor.Close(ctx)
		eff := entity.LabourEfficiency{}
		eff.FirstName = user.FirstName
		eff.LastName = user.LastName
		eff.EmployeeId = user.Id

		empEff := entity.EmployeeEfficiencyFetch{}

		for cursor.Next(ctx) {
			err = cursor.Decode(&empEff)
			if err != nil {
				return nil, err
			}
			lUtilization := (empEff.TotalHoursWorked / totalShiftHours) * 100
			s1 := strconv.FormatFloat(lUtilization, 'f', 2, 64)
			res, err := strconv.ParseFloat(s1, 64)
			if err != nil {
				fmt.Println(err)
			}
			eff.LaborUtilization = res
			lProductivity := (empEff.TotalHoursSold / empEff.TotalHoursWorked) * 100
			s2 := strconv.FormatFloat(lProductivity, 'f', 2, 64)
			res2, err2 := strconv.ParseFloat(s2, 64)
			if err2 != nil {
				fmt.Println(err2)
			}
			eff.LaborProductivity = res2
			eff.LaborEfficiency = eff.LaborUtilization * eff.LaborProductivity

		}
		labourEfficiency = append(labourEfficiency, &eff)

	}
	return labourEfficiency, nil
}
