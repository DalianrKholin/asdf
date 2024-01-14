package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"niceSite/model"
	"niceSite/views"
)

func (s *ApiDbEndpoints) AddProduct(w http.ResponseWriter, r *http.Request) {
	connect := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	var prod model.Product
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &prod)
	if err != nil {
		fmt.Printf("%v\n", prod)
	}
	index, err := connect.InsertOne(Background, prod)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	views.ResponseWithJSON(w, 200, index)
}
