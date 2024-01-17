package dataBaseModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Price      float64            `bson:"price" json:"price"`
	Properties string             `bson:"properties" json:"properties"`
	InStack    int                `bson:"inStack" json:"inStack"`
}
