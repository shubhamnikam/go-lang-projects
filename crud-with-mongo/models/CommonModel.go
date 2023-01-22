package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Movie struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Watched bool `json:"watched,omitempty"`
}