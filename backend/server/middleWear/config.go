package middleWear

import ("go.mongodb.org/mongo-driver/mongo"
    "context")

type ApiDbMiddleWear struct {
	DB *mongo.Client
}


var DataBaseName string = "site"
var UsersCollection string = "users"
var ProductsCollection string = "products"
var Background context.Context = context.Background()