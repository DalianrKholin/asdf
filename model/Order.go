package model

type Order struct {
	UserMail   string    `bson:"user_mail"`
	Product    []Product `bson:"product"`
	Amounts    []int     `bson:"amounts"`
	TotalProce float64   `bson:"totalProce"`
}
