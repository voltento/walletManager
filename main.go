package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/voltento/pursesManager/account_managing"
	"github.com/voltento/pursesManager/browsing"
	"github.com/voltento/pursesManager/config"
	"github.com/voltento/pursesManager/database"
	"github.com/voltento/pursesManager/payment"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg := loadConfig()
	r := mux.NewRouter()
	dbCtrl, er := database.CreateWalletMgrCluster(cfg.Db.User, cfg.Db.Password, cfg.Db.Name, cfg.Db.Addr, cfg.Db.DbPoolSize)
	if er != nil {
		log.Fatal(fmt.Sprintf("Error: %v", er.Error()))
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
	address := cfg.Addr
	log.Printf("Start listen: %v", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func loadConfig() config.Config {
	if len(os.Args) < 2 {
		fmt.Printf("Expected params: <path to config file>\n")
		os.Exit(1)
	}
	path := os.Args[1]
	dat, er := ioutil.ReadFile(path)
	if er != nil {
		fmt.Printf("Error ocured during load config. Error: %v\n", er.Error())
	}

	cfg := config.Config{}
	er = json.Unmarshal(dat, &cfg)
	if er != nil {
		fmt.Printf("Error ocured during load config. Error: %v\n", er.Error())
		os.Exit(1)
	}

	return cfg
}
