package model

type Product struct {
	Name       string  `bson:"name"`
	Price      float64 `bson:"price"`
	Properties string  `bson:"properties"`
	InStack    int     `bson:"inStack"`
}
