package model

type Product struct {
	Name        string  `bson:"name"`
	Price       float64 `bson:"price"`
	Description string  `bson:"description"`
	InStack     int     `bson:"inStack"`
}
