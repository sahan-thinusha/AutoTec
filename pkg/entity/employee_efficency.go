package entity

type EmployeeEfficiencyFetch struct {
	EmployeeId          string  `json:"employeeId" bson:"_id"`
	TotalHoursEstimated float64 `json:"totalHoursEstimated" bson:"totalHoursEstimated"`
	TotalHoursWorked    float64 `json:"totalHoursWorked" bson:"totalHoursWorked"`
	TotalHoursSold      float64 `json:"totalHoursSold" bson:"totalHoursSold"`
}
