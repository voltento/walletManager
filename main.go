package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/voltento/pursesManager/service"
	"github.com/voltento/pursesManager/transport"
	"log"
	"net/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	svc := service.CreateService()

	uppercaseHandler := httptransport.NewServer(
		transport.MakeUppercaseEndpoint(svc),
		transport.DecodeUppercaseRequest,
		transport.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)

	address := ":8080"
	log.Printf("Start listen: %v", address)
	log.Fatal(http.ListenAndServe(address, nil))

}
