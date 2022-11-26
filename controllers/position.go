package controllers

import (
	"context"
	"encoding/json"
	"job-search-api/models"
	"job-search-api/responses"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var externalAPIEndpoint = "http://dev3.dansmultipro.co.id/api/recruitment"

func GetPositions() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		description := r.URL.Query().Get("description")
		location := r.URL.Query().Get("location")
		page := r.URL.Query().Get("page")

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		var err error
		var client = &http.Client{}
		var data []models.Position

		request, err := http.NewRequest("GET", externalAPIEndpoint+"/positions.json?description="+description+"&location="+location+"&page="+page, nil)
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
			data = []models.Position{}
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "", Data: map[string]interface{}{"results": data}}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return
	}
}

func GetPosition() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		positionId := params["positionId"]

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		var err error
		var client = &http.Client{}
		var data models.Position

		request, err := http.NewRequest("GET", externalAPIEndpoint+"/positions/"+positionId, nil)
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
			data = models.Position{}
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "", Data: data}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return
	}
}
