package middlewares

import (
	"context"
	"encoding/json"
	"job-search-api/models"
	"job-search-api/responses"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		notAuth := []map[string]string{
			{
				"path":   "/login",
				"method": "POST",
			},
		} //List of endpoints that doesn't require jwt auth
		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, mapItem := range notAuth {
			if mapItem["path"] == requestPath && mapItem["method"] == r.Method {
				next.ServeHTTP(rw, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusForbidden, Message: "Missing auth token", Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusForbidden, Message: "Invalid token format", Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_secret_key")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusForbidden, Message: err.Error(), Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusForbidden)
			response := responses.BaseResponse{Status: http.StatusForbidden, Message: "Token is not valid", Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		ctx := context.WithValue(r.Context(), "user-email", tk.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r) //proceed in the middleware chain!
	})
}
