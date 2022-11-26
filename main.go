package main

import (
	"fmt"
	"job-search-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login()).Methods("POST")

	fmt.Println("Starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
