package model

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Record struct {
		ID 		bson.ObjectId	`json:"id" bson:"_id,omitempty"`
		Name	string			`json:"name" bson:"name"`
		Score	string			`json:"score" bson:"score"`	
	}
)