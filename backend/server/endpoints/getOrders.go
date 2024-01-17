package server

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"niceSite/backend/model/dataBaseModel"
	"niceSite/backend/views"
)

func (s *ApiDbEndpoints) GetOrders(w http.ResponseWriter, r *http.Request) {
	connection := s.DB.Database(DataBaseName).Collection(OrderCollection)
	cursor, err := connection.Find(Background, bson.M{})
	defer cursor.Close(Background)

	var results []dataBaseModel.Order

	if err = cursor.All(context.TODO(), &results); err != nil {
		views.ResponseWithError(w, 500, "server stupido\n")
		return
	}
	views.ResponseWithJSON(w, 200, results)
}
