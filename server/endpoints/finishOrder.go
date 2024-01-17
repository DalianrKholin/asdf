package server

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	. "niceSite/model/dataBaseModel"
	"niceSite/model/serverResponseModel"
	. "niceSite/views"
)

func (s *ApiDbEndpoints) FinishOrder(w http.ResponseWriter, r *http.Request) {
	userConnect := s.DB.Database(DataBaseName).Collection(UsersCollection)
	orderConnect := s.DB.Database(DataBaseName).Collection(OrderCollection)
	session, err := s.DB.StartSession()
	var finishOrder requestForFinishOrder
	bodyReader, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(bodyReader, &finishOrder)
	if err != nil {
		ResponseWithError(w, 400, "bad request")
		return
	}
	defer session.EndSession(Background)

	err = session.StartTransaction()
	if err != nil {
		ResponseWithError(w, 500, "server error")
		return
	}
	token, err := primitive.ObjectIDFromHex(r.Header.Get("token"))
	if err != nil {
		ResponseWithError(w, 400, "where token?")
		return
	}
	res := userConnect.FindOne(Background, bson.M{"token": token})
	var user User
	err = res.Decode(&user)
	if err != nil {
		ResponseWithError(w, 418, err.Error())
		return
	}
	err = mongo.WithSession(context.Background(), session, func(sc mongo.SessionContext) error {
		res := orderConnect.FindOne(Background, bson.M{"_id": finishOrder.OrderId})

		var order Order
		err := res.Decode(&order)
		if err != nil {
			return err
		}
		if order.UserMail != user.Mail {
			ResponseWithError(w, 401, "bad user")
		}

		_, err = orderConnect.UpdateOne(Background, bson.M{"_id": finishOrder.OrderId}, bson.M{"$set": bson.M{"finished": true}})
		if err != nil {
			return err
		}
		_, err = userConnect.UpdateOne(Background, bson.M{"_id": user.Id}, bson.M{"$pull": bson.M{"orders": order.Id}})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		ResponseWithError(w, 400, err.Error())
		return
	}
	if err := session.CommitTransaction(Background); err != nil {
		ResponseWithError(w, 500, "server cant make finishOrder")
		return
	}

	ServerResponse := serverResponseModel.FinishOrderServerResult{
		OrderId:     finishOrder.OrderId.Hex(),
		UserUpdated: user.Id.Hex(),
		OrderStatus: true,
	}

	ResponseWithJSON(w, 200, ServerResponse)
}

type requestForFinishOrder struct {
	OrderId primitive.ObjectID `json:"orderId"`
}
