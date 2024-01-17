package dataBaseModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	Id       primitive.ObjectID `bson:"_id"`
	Products []Product          `bson:"products"`
	Amounts  int                `bson:"amounts"`
}
