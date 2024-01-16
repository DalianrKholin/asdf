package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	s "niceSite/server/endpoints"
	"niceSite/server/middleWear"
	"os"
)

func main() {

	err := godotenv.Load(".env")        // loading .env to let us read it
	conString := os.Getenv("conString") //geting connection string from .env
	port := os.Getenv("PORT")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(conString).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	defer func() {
		err = client.Disconnect(context.Background()) //connecting to mongo
		if err != nil {
			panic(err)
		}
	}()
	//now we are connected to mongo
	mainRouter := chi.NewRouter()
	if err != nil {
		panic(err)
	}
	serv := &http.Server{
		Handler: mainRouter,
		Addr:    ":" + port,
	}

	apiDbEndpoints := s.ApiDbEndpoints{
		DB: client,
	}
	apiMiddleWear := middleWear.ApiDbMiddleWear{
		DB: client,
	}
	mainRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	admin := chi.NewRouter()
	nUser := chi.NewRouter()


	mainRouter.Mount("/api", nUser)
	nUser.Mount("/admin", admin)
    nUser.Post("/login",apiMiddleWear.SaveData(apiDbEndpoints.LoginIntoApp))

	admin.Post("/user", apiMiddleWear.SaveData(apiMiddleWear.Auth(apiDbEndpoints.AddUserApi)))//http://localhost:8080/api/admin/user

	admin.Post("/item", apiMiddleWear.SaveData(apiDbEndpoints.AddProduct))//http://localhost:8080/api/admin/item

    admin.Post("/item/edit", apiMiddleWear.SaveData(apiMiddleWear.EnableCors(apiDbEndpoints.EditProduct)))// http://localhost:8080/api/admin/item/edit
	admin.Delete("/item", apiMiddleWear.SaveData(apiDbEndpoints.DelItems))//http://localhost:8080/api/admin/item

	nUser.Get("/item", apiMiddleWear.SaveData(apiDbEndpoints.GetItems))//http://localhost:8080/api/item

	err = serv.ListenAndServe()
	if err != nil {
		fmt.Printf("server is dead by %v\n", err)
	}
}
