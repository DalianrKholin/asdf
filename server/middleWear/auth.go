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
        token,err := primitive.ObjectIDFromHex(r.Header.Get("token"))
        if err!=nil{
            ResponseWithError(w, 400, "where token?")
            return
        }
        res := connection.FindOne(Background, bson.M{"token": token})
        if res==nil{
            ResponseWithError(w, 418, "have to login in")
            return
        }
        var user User
        err = res.Decode(&user)
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
