package dataBaseModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	UserMail    string      `bson:"userMail" json:"userMail,omitempty"`
	ProductInfo []OrderSpan `bson:"productInfo" json:"productInfo"`
	TotalPrice  float64     `bson:"totalPrice" json:"totalPrice,omitempty"`
	Finished    bool        `bson:"finished" json:"finished"`
}

type OrderSpan struct {
	ProductID primitive.ObjectID `bson:"productID" json:"productID,omitempty"`
	Amount    int                `bson:"amount" json:"amount,omitempty"`
}
