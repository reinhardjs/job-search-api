package controllers

import (
	"context"
	"encoding/json"
	"job-search-api/models"
	"job-search-api/responses"
	"net/http"
	"time"
)

var externalAPIEndpoint = "http://dev3.dansmultipro.co.id/api/recruitment"

func GetPositions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		var err error
		var client = &http.Client{}
		var data []models.Position

		request, err := http.NewRequest("GET", externalAPIEndpoint+"/positions.json", nil)
		if err != nil {
			response := responses.BaseResponse{Status: http.StatusOK, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		res, err := client.Do(request)
		if err != nil {
			response := responses.BaseResponse{Status: http.StatusOK, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}
		defer res.Body.Close()

		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			response := responses.BaseResponse{Status: http.StatusOK, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "", Data: map[string]interface{}{"results": data}}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return
	}
}
