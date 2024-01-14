package server

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"niceSite/views"
)

func (s *ApiDbEndpoints) GetItems(w http.ResponseWriter, r *http.Request) {
	connection := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	items, err := connection.Find(Background, bson.M{})
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 400, "me stupido\n")
	}
	fmt.Printf("%v\n", items)
	views.ResponseWithJSON(w, 200, items)
}
