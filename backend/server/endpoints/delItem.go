package server

import (
	"fmt"
	"net/http"
	. "niceSite/backend/views"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ApiDbEndpoints) DelItems(w http.ResponseWriter, r *http.Request) {
	connection := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
	if err != nil {
		ResponseWithError(w, 400, "bad id")
		return
	}
	index, err := connection.DeleteOne(Background, bson.M{"_id": id})

	if err != nil {
		fmt.Printf("%v\n", err)
		ResponseWithError(w, 400, "bad request")
		return
	}
	ResponseWithJSON(w, 200, index)
}
