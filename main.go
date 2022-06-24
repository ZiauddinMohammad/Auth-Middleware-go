package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/Auth-Middleware-go/authservice"
)

func main() {
	router := mux.NewRouter()

	//create a auth subrouter for authentication
	auth_router := router.PathPrefix("/auth").Subrouter()

	auth_router.HandleFunc("/signup", authservice.Signuphandler).Methods("POST")
	auth_router.HandleFunc("/signin", authservice.Signinhandler).Methods("GET")

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server started and listening at port 8080")
	log.Fatal(server.ListenAndServe())

}
