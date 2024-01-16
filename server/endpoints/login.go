package server

import (
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	. "niceSite/model"
    . "niceSite/views"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
"encoding/json"
"io"
)




func (s *ApiDbEndpoints) LoginIntoApp(w http.ResponseWriter, r *http.Request) {
    connection := s.DB.Database(DataBaseName).Collection(UsersCollection)
    var creds Creds
    bodyReader, _ := io.ReadAll(r.Body)
    err := json.Unmarshal(bodyReader, &creds)
    if err!= nil{
        fmt.Printf("%v\n",err )
        ResponseWithError(w, 400, "bad request")
        return
    }
    res := connection.FindOne(context.Background(), bson.D{
        {"mail", creds.Email},
    })
    user := User{}
    err = res.Decode(&user)
    if err != nil {
        	ResponseWithError(w, 401, "err")
        	return
       }
    if user.Password !=  creds.Password{
        ResponseWithError(w, 418, "bad password")
        return
    }

    authToken := primitive.NewObjectID()
    filter := bson.M{"id": user.Id}
    update := bson.M{
        "$set": bson.M{"token": authToken},
    }

    // Aktualizacja dokumentu
    _, err = connection.UpdateOne(Background, filter, update)
    if err != nil {
        ResponseWithError(w,500, "server stupido\n")
        return
    }
    w.Header().Set("authToken", authToken.Hex())
    ResponseWithJSON(w, 200, "logged in")

}


