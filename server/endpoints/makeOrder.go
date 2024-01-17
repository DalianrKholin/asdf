package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"niceSite/model/dataBaseModel"
	. "niceSite/model/serverResponseModel"
	. "niceSite/views"
)

func makeOrderTransactionContent(s *mongo.Client, order dataBaseModel.Order) error {
	connectProduct := s.Database(DataBaseName).Collection(ProductsCollection)
	//connectOrder := s.Database(DataBaseName).Collection(OrderCollection)
	connectUser := s.Database(DataBaseName).Collection(UsersCollection)

	ids := make([]primitive.ObjectID, len(order.ProductInfo))
	for i, v := range order.ProductInfo {
		ids[i] = v.ProductID
	}

	cursor, err := connectProduct.Find(Background, bson.M{"_id": bson.M{"$in": ids}})
	defer cursor.Close(Background)
	if err != nil {
		return err
	}

	var products []dataBaseModel.Product

	if err = cursor.All(context.TODO(), &products); err != nil {
		return err
	}

	price := float64(0)
	for i, v := range products {
		amount := order.ProductInfo[i].Amount
		if v.InStack < amount {
			return errors.New("not enough prods in stack")
		}
		price += v.Price * float64(amount)
	}
	fmt.Printf("%v\n", price)

	//var serverResult MakeOrderResult

	userDataRes := connectUser.FindOne(Background, bson.M{"_id": order.UserMail})
	//insert, err := connectOrder.InsertOne(Background, order)
	if err != nil {
		return err
	}

	var user dataBaseModel.User
	err = userDataRes.Decode(&user)
	if err != nil {
		return err
	}

	return nil
}

func (s *ApiDbEndpoints) MakeOrder(w http.ResponseWriter, r *http.Request) {
	var err error
	if err != nil {
		ResponseWithError(w, 500, "server error")
		return
	}

	var order dataBaseModel.Order
	bodyReader, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(bodyReader, &order)
	if err != nil {
		fmt.Printf("%v\n", err)
		ResponseWithError(w, 400, "bad request")
		return
	}
	session, err := s.DB.StartSession()
	defer session.EndSession(Background)
	err = session.StartTransaction()
	if err != nil {
		ResponseWithError(w, 500, "server error")
		return
	}
	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		res := makeOrderTransactionContent(s.DB, order)

		if res != nil {
			return res
		}

		return nil
	})
	if err != nil {
		ResponseWithError(w, 400, err.Error())
		return
	}
	if err := session.CommitTransaction(Background); err != nil {
		ResponseWithError(w, 500, "server cant make order")
		return
	}
	ResponseWithJSON(w, 200, order)
}
