package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/voltento/walletManager/internal/httpQueryModels"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type CheckResp = func(resp httpResp) error

const (
	walltMgrAddr   = "http://:8080"
	addAccountUrl  = walltMgrAddr + "/accmamaging/add/"
	getAccountsUrl = walltMgrAddr + "/browsing/accounts"
	sendMoneyUrl   = walltMgrAddr + "/payment/send_money"
	getPaymentsUrl = walltMgrAddr + "/browsing/payments"
)

func assertEqHttpResp(resp httpResp, expected httpResp) error {
	if !reflect.DeepEqual(resp, expected) {
		return errors.New(fmt.Sprintf("Got unexpected message. Expected: '%v', got '%v'", expected, resp))
	}
	return nil
}

func getAccounts() ([]httpQueryModels.Account, error) {
	resp, er := sendRequest(getAccountsUrl, "GET", "")
	if er != nil {
		return nil, er
	}

	var accs []httpQueryModels.Account
	er = json.NewDecoder(strings.NewReader(resp.data)).Decode(&accs)
	if er != nil {
		return nil, er
	}
	return accs, nil
}

func assertAccExists(id string) error {
	_, er := getAccount(id)
	return er
}

func getAccount(id string) (httpQueryModels.Account, error) {
	accs, er := getAccounts()
	if er != nil {
		return httpQueryModels.Account{}, nil
	}

	for _, ac := range accs {
		if ac.Id == id {
			return ac, nil
		}
	}

	return httpQueryModels.Account{}, errors.New(fmt.Sprintf("Can't find account with id='%v'", id))
}

func addAccount(ac httpQueryModels.Account) error {
	var b bytes.Buffer
	er := json.NewEncoder(&b).Encode(ac)
	if er != nil {
		return er
	}
	_, er = sendRequest(addAccountUrl, "PUT", b.String())
	return er
}

func addAccountsWithCurrency(currency string, accCount int) ([]httpQueryModels.Account, error) {
	var er error
	rand.Seed(time.Now().UTC().UnixNano())

	accs := make([]httpQueryModels.Account, 0, accCount)
	for i := 0; i < accCount; i += 1 {
		ac := httpQueryModels.Account{Id: "test_" + strconv.Itoa(rand.Intn(10000000)), Currency: currency, Amount: 10}
		if er = addAccount(ac); er != nil {
			return nil, er
		}
		accs = append(accs, ac)
	}

	return accs, nil
}
