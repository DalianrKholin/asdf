package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"niceSite/model/dataBaseModel"
	"niceSite/model/serverResponseModel"
	. "niceSite/views"
)

func (s *ApiDbEndpoints) AddProduct(w http.ResponseWriter, r *http.Request) {
	connect := s.DB.Database(DataBaseName).Collection(ProductsCollection)
	var prod dataBaseModel.Product
	bodyReader, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(bodyReader, &prod)
	if err != nil {
		fmt.Printf("%v\n", prod)
		ResponseWithError(w, 400, "bad request")
		return
	}
	_, err = connect.InsertOne(Background, prod)

	serverResponse := serverResponseModel.AddItemResponse{
		Name:        prod.Name,
		Amount:      prod.InStack,
		Description: prod.Properties,
	}

	if err != nil {
		fmt.Printf("%v\n", prod)
		ResponseWithError(w, 400, "cant add")
		return
	}
	ResponseWithJSON(w, 200, serverResponse)
}
