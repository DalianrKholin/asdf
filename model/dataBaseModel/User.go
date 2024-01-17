package dataBaseModel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nick           string             `bson:"nick" json:"nick,omitempty"`
	Mail           string             `bson:"mail" json:"mail,omitempty"`
	Password       string             `bson:"password" json:"password,omitempty"`
	Admin          bool               `bson:"admin" json:"admin,omitempty"`
	Token          primitive.ObjectID `bson:"token" json:"token,omitempty"`
	OrderInExecute []Order            `bson:"orderInExecute" json:"orderInExecute,omitempty"`
}
