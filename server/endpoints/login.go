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
    "go.mongodb.org/mongo-driver/mongo"
    "time"
)

func TokenCleaner(id primitive.ObjectID, db *mongo.Client){
    <-time.NewTicker(10 * time.Minute).C
    connection := db.Database(DataBaseName).Collection(UsersCollection)
    filter := bson.M{"_id": id}
    val,_ := primitive.ObjectIDFromHex("000000000000")
        update := bson.M{
            "$set": bson.M{"token": val},
        }
    connection.UpdateOne(Background,filter,update)

}


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
        	ResponseWithError(w, 400, "bad password")
        	return
       }
    if user.Password !=  creds.Password{
        ResponseWithError(w, 418, "bad password")
        return
    }

    authToken := primitive.NewObjectID()
    filter := bson.M{"_id": user.Id}
    update := bson.M{
        "$set": bson.M{"token": authToken},
    }

    // Aktualizacja dokumentu
    _, err = connection.UpdateOne(Background, filter, update)
    go TokenCleaner(user.Id, s.DB)
    if err != nil {
        ResponseWithError(w,500, "server stupido\n")
        return
    }

    var token Token
    token.Token= authToken.Hex()
    w.Header().Set("Content-Type", "application/json")
    ResponseWithJSON(w, 200, token)
}

type Token struct{
    Token string `json:"token"`
}
