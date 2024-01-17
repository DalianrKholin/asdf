package dataBaseModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name" json:"name,omitempty"`
	Price      float64            `bson:"price" json:"price,omitempty"`
	Properties string             `bson:"properties" json:"properties,omitempty"`
	InStack    int                `bson:"inStack" json:"inStack,omitempty"`
}
