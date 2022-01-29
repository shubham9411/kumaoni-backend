package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shubham9411/kumaoni-backend/routes"
)

func main() {
	fmt.Println("Welcome to Kumaoni API")
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":3000", r))
}
