package server

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	. "niceSite/views"
)

func (s *ApiDbEndpoints) AddUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/addUser.html")
}

func mapStringToInt(tab string) int {
	result := int32(0)
	for _, x := range tab {
		result += x - '0'
		result *= 10
	}
	return int(result / 10)
}

func HashSHA256(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func mapQueryParamsToBSONM(queryParams map[string][]string, pass string) bson.M {
	filter := bson.M{}
	for key, values := range queryParams {
		if len(values) == 1 {
			if key == "age" {
				filter[key] = mapStringToInt(values[0])
				continue
			}
			filter[key] = values[0]
		} else {
			fmt.Printf("slice : %v\n", values)
			filter[key] = values
		}
	}
	filter["password"] = pass
	if filter["admin"] == nil {
		filter["admin"] = false
	}
	return filter
}

func (s *ApiDbEndpoints) AddUserApi(w http.ResponseWriter, r *http.Request) {
	connect := s.DB.Database(DataBaseName).Collection(UsersCollection)
	count, err := connect.CountDocuments(context.Background(), bson.D{{"name", r.URL.Query().Get("name")}}) // users already exists?
	if count > 0 {
		ResponseWithError(w, 406, "user already exists")
		return
	}
	pass := HashSHA256(r.Header.Get("pass"))
	data, err := connect.InsertOne(context.Background(), mapQueryParamsToBSONM(r.URL.Query(), pass)) //dodanie usera
	_ = data
	if err != nil {
		ResponseWithError(w, 500, "cant inset user")
		return
	}
	ResponseWithJSON(w, 200, struct {
		ID interface{}
	}{
		ID: data.InsertedID,
	})
}
