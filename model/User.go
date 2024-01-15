package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	id       primitive.ObjectID `bson:"_id"`
	Name    string `bson:"nick"`
	Mail     string             `bson:"mail"`
	Password string             `bson:"password"`
	Admin    bool               `bson:"admin"`
	Token string                `bson:"token"`
}
