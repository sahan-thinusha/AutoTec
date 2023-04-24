package entity

type LabourEfficiency struct {
	EmployeeId        string  `json:"employeeId" bson:"employeeId"`
	FirstName         string  `json:"firstName" bson:"firstName"`
	LastName          string  `json:"lastName" bson:"lastName"`
	LaborUtilization  float64 `json:"laborUtilization" bson:"laborUtilization"`
	LaborProductivity float64 `json:"laborProductivity" bson:"laborProductivity"`
	LaborEfficiency   float64 `json:"laborEfficiency" bson:"laborEfficiency"`
}
