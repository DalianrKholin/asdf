package server

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"niceSite/backend/model/dataBaseModel"
	"niceSite/backend/views"
)

func (s *ApiDbEndpoints) GetItems(w http.ResponseWriter, r *http.Request) {
	connection := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	cursor, err := connection.Find(Background, bson.M{})
	defer cursor.Close(Background)

	var results []dataBaseModel.Product

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Printf("%v\n", err)
		views.ResponseWithError(w, 400, "me stupido\n")
		return
	}
	views.ResponseWithJSON(w, 200, results)
}
