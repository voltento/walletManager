package main

import (
	"github.com/gorilla/mux"
	"github.com/voltento/pursesManager/account_managing"
	"log"
	"net/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	r := mux.NewRouter()
	s := account_managing.CreateService()
	r.Handle("/account_managing/add/", account_managing.MakeHandler(s)).Methods("POST")

	http.Handle("/", r)
	address := ":8080"
	log.Printf("Start listen: %v", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
