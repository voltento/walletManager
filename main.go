package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/voltento/pursesManager/account_managing"
	"github.com/voltento/pursesManager/browsing"
	"github.com/voltento/pursesManager/database"
	"github.com/voltento/pursesManager/payment"
	"log"
	"net/http"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	r := mux.NewRouter()
	dbCtrl, err := database.CreateWalletMgrCluster(10)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %v", err.Error()))
	}

	s := account_managing.CreateService(dbCtrl)
	r.Handle("/account_managing/add/", account_managing.MakeHandler(s)).Methods("PUT")

	b := browsing.CreateService(dbCtrl)
	r.Handle("/browsing/accounts", browsing.MakeGetAccountsHandler(b)).Methods("GET")
	r.Handle("/browsing/payments", browsing.MakeGetPaymentsHandler(b)).Methods("GET")

	p := payment.CreateService(dbCtrl)
	r.Handle("/payment/change_balance", payment.MakeChangeBalanceHandler(p)).Methods("PUT")
	r.Handle("/payment/send_money", payment.MakeSendMoneyHandler(p)).Methods("PUT")

	http.Handle("/", r)
	address := ":8080"
	log.Printf("Start listen: %v", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
