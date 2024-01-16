package middleWear

import (
	"net/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
		. "niceSite/model"
        . "niceSite/views"
)

func (s *ApiDbMiddleWear) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        connection := s.DB.Database(DataBaseName).Collection(UsersCollection)
        token,_ :=primitive.ObjectIDFromHex(r.Header.Get("token"))
        res :=connection.FindOne(Background, bson.D{
            {"token", token},})
        if res!=nil{
            ResponseWithError(w, 418, "have to login in")
            return
        }
        user := User{}
        err := res.Decode(&user)
        if err!=nil{
            ResponseWithError(w, 500, "me stupido")
            return
        }
        if user.Admin!=true{
            ResponseWithError(w, 401, "need authorization")
            return
        }
		next(w, r)
	}
}
