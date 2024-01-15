package server

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"niceSite/model"
    . "niceSite/views"
	"context"
	"fmt"
"encoding/json"
"io"
)




func (s *ApiDbEndpoints) LoginIntoApp(w http.ResponseWriter, r *http.Request) {
    connection := s.DB.Database(DataBaseName).Collection(UsersCollection)
    var prod Creds
    bodyReader, _ := io.ReadAll(r.Body)
    err := json.Unmarshal(bodyReader, &prod)
    if err!= nil{
        fmt.Printf("%v\n",err )
        ResponseWithError(w, 400, "bad request")
        return
    }
    res := connection.FindOne(context.Background(), bson.D{
        {"mail", prod.Email},
    })
    	user := model.User{}
    	err = res.Decode(&user)

    fmt.Printf("%v %v \n",user, prod)
if err != nil {
    		ResponseWithError(w, 401, "cant find user")
    		return
    	}

}


type Creds struct{
    Password string `bson:"password"`
    Email string `bson:"mail"`
}