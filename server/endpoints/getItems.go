package server

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"niceSite/model"
	"niceSite/views"
)

func (s *ApiDbEndpoints) GetItems(w http.ResponseWriter, r *http.Request) {
	connection := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	number, err := connection.CountDocuments(Background, bson.M{})
	currsor, err := connection.Find(Background, bson.M{})
	defer currsor.Close(Background)
	results := make([]model.Product, number)
	i := 0
	for currsor.Next(Background) {
		var result model.Product
		err := currsor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results[i] = result
		i += 1
	}
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 400, "me stupido\n")
	}
	views.ResponseWithJSON(w, 200, results)
}
