package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"niceSite/model"
	. "niceSite/views"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *ApiDbEndpoints) AddProduct(w http.ResponseWriter, r *http.Request) {
	connect := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	var prod model.Product
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &prod)
	if err != nil {
		fmt.Printf("%v\n", prod)
		ResponseWithError(w, 400, "bad request")
		return
	}
	prod.Id =primitive.NewObjectID()
	_, err = connect.InsertOne(Background, prod)
	if err != nil {
		fmt.Printf("%v\n", prod)
		ResponseWithError(w, 400, "cant add")
		return
	}
	ResponseWithJSON(w, 200, "added")
}
