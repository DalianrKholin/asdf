package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	. "niceSite/model"
	. "niceSite/views"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *ApiDbEndpoints) EditProduct(w http.ResponseWriter, r *http.Request) {
    connect := s.DB.Database(DataBaseName).Collection(ProductsCollection)
    id, err := primitive.ObjectIDFromHex(r.URL.Query().Get("id"))
    	if err != nil {
    		ResponseWithError(w, 400, "bad id")
    		return
    	}
    filter := bson.M{"_id": id}
    	var prod Product
    	bodyReader, _ := io.ReadAll(r.Body)
    	err = json.Unmarshal(bodyReader, &prod)
    	if err != nil {
    		fmt.Printf("%v\n", prod)
    		ResponseWithError(w, 400, "bad request")
    		return
    	}
    prod.Id,_  = primitive.ObjectIDFromHex(r.URL.Query().Get("id"))

    _, err = connect.ReplaceOne(Background, filter, prod)
    if err != nil{
        		ResponseWithError(w, 400, "err")
        		return
    }
    ResponseWithJSON(w, 200, "updated")
}