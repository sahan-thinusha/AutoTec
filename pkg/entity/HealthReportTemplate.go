package entity

type HealthReportTemplate struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"Name"`
}
