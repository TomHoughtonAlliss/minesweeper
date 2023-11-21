package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Record struct {
		ID 		primitive.ObjectID	`json:"id" bson:"_id,omitempty"`
		Name	string				`json:"name" bson:"name"`
		Score	string				`json:"score" bson:"score"`	
	}
)
