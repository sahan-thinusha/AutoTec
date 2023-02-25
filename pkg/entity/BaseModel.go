package entity

import "time"

type Base struct {
	CreatedAt *time.Time `bson:"created_at" swaggerignore:"true"`
	UpdatedAt *time.Time `bson:"updated_at" swaggerignore:"true"`
	DeletedAt *time.Time `bson:"deleted_at" swaggerignore:"true"`
}
