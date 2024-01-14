package model

type Cart struct {
	Products []Product `bson:"products"`
	Amounts  int       `bson:"amounts"`
}
