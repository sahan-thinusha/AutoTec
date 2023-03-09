package entity

type WorkProgress struct {
	Id         string `json:"id" bson:"_id"`
	Base       `bson:",inline"`
	JobId      string `json:"jobId" bson:"jobId"`
	Job        Job    `json:"job" bson:"job"`
	Status     string `json:"status" bson:"status"`
	Time       int64  `json:"time" bson:"time"`
	CustomerId string `json:"customerId" bson:"customerId"`
	Tasks      []Task `json:"tasks" bson:"tasks"`
}

type Task struct {
	TaskNo      int64  `json:"task_no" bson:"task_no"`
	Name        string `json:"name" bson:"Name"`
	IsCompleted bool   `json:"isCompleted" bson:"isCompleted"`
}
