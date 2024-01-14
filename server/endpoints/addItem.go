package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"niceSite/model"
)

func (s *ApiDbEndpoints) AddProduct(w http.ResponseWriter, r *http.Request) {
	connect := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	connect.Name()
	var mes model.Product
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &mes)
	if err != nil {
		fmt.Printf("%v\n", mes)
	}
}
