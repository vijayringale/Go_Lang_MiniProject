package entiti

import "gopkg.in/mgo.v2/bson"

type Product struct {
	Id bson.ObjectId `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age int `json:"age" bson`

}