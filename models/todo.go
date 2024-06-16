package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ToDo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task        string             `json:"task,omitempty" bson:"task,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Created_At  string             `json:"created-at,omitempty" bson:"created-at,omitempty"`
}
