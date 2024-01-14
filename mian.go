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

	mainRouter.Get("/", apiMiddleWear.SaveData(apiMiddleWear.Auth(apiDbEndpoints.MainSite)))

	mainRouter.Mount("/admin", admin)
	mainRouter.Mount("/user", nUser)

	admin.Get("/addUser", apiMiddleWear.SaveData(apiMiddleWear.Auth(apiDbEndpoints.AddUser)))
	admin.Post("/addUser", apiMiddleWear.SaveData(apiMiddleWear.Auth(apiDbEndpoints.AddUserApi)))

	admin.Post("/addItem", apiMiddleWear.SaveData(apiDbEndpoints.AddProduct))

	admin.Patch("/updateItem", apiMiddleWear.SaveData(apiDbEndpoints.AddProduct))

	nUser.Get("/itemList", apiMiddleWear.SaveData(apiMiddleWear.Auth(apiDbEndpoints.GetItems)))

	err = serv.ListenAndServe()
	if err != nil {
		fmt.Printf("server is dead by %v\n", err)
	}
}
