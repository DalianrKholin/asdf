package server

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var DataBaseName string = "site"
var UsersCollection string = "users"
var ProductsCollection string = "products"
var Background context.Context = context.Background()

type ApiDbEndpoints struct {
	DB *mongo.Client
}
