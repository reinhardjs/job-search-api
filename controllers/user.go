package controllers

import (
	"context"
	"encoding/json"
	"job-search-api/configs"
	"job-search-api/models"
	"job-search-api/responses"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Login() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		user := &models.User{}

		rw.Header().Add("Content-Type", "application/json")

		var err = json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
		if err != nil {
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		filter := bson.M{"email": user.Email}

		var result models.User
		var usersCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
		err = usersCollection.FindOne(ctx, filter).Decode(&result)

		if err != nil && err == mongo.ErrNoDocuments {
			response := responses.BaseResponse{Status: http.StatusNotFound, Message: "User not found", Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		// create JWT token
		threeMinutes := (time.Hour / 60) * 3
		tk := &models.Token{Email: result.Email, RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(threeMinutes)),
		}}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("token_secret_key")))

		response := responses.BaseResponse{Status: http.StatusOK, Message: "token", Data: map[string]interface{}{"email": user.Email, "token": tokenString}}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return
	}
}
