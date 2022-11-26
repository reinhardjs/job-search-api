package main

import (
	"fmt"
	"job-search-api/controllers"
	"job-search-api/middlewares"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// add JWTAuth middleware
	router.Use(middlewares.JwtAuthentication)

	router.HandleFunc("/positions", controllers.GetPositions()).Methods("GET")
	router.HandleFunc("/login", controllers.Login()).Methods("POST")

	fmt.Println("Starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
