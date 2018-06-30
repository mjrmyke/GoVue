package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initEndPoints(apiRouter *mux.Router) {
	apiRouter.HandleFunc("/user", retrieveNewUser).Methods("GET")
	apiRouter.HandleFunc("/user", createNewUser).Methods("GET")
}

func retrieveNewUser(response http.ResponseWriter, request *http.Request) {

}

func createNewUser(response http.ResponseWriter, request *http.Request) {

}
