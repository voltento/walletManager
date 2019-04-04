package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/voltento/walletManager/accmamaging"
	"reflect"
	"strings"
)

type CheckResp = func(resp httpResp) error

const (
	walltMgrAddr   = "http://:8080"
	addAccountUrl  = walltMgrAddr + "/accmamaging/add/"
	getAccountsUrl = walltMgrAddr + "/browsing/accounts"
	sendMoneyUrl   = walltMgrAddr + "/payment/send_money"
)

func assertEqHttpResp(resp httpResp, expected httpResp) error {
	if !reflect.DeepEqual(resp, expected) {
		return errors.New(fmt.Sprintf("Got unexpected message. Expected: '%v', got '%v'", expected, resp))
	}
	return nil
}

func getAccounts() ([]accmamaging.Account, error) {
	resp, er := sendRequest(getAccountsUrl, "GET", "")
	if er != nil {
		return nil, er
	}

	var accs []accmamaging.Account
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

func getAccount(id string) (accmamaging.Account, error) {
	accs, er := getAccounts()
	if er != nil {
		return accmamaging.Account{}, nil
	}

	for _, ac := range accs {
		if ac.Id == id {
			return ac, nil
		}
	}

	return accmamaging.Account{}, errors.New(fmt.Sprintf("Can't find account with id='%v'", id))
}

func addAccount(ac accmamaging.Account) error {
	var b bytes.Buffer
	er := json.NewEncoder(&b).Encode(ac)
	if er != nil {
		return er
	}
	_, er = sendRequest(addAccountUrl, "PUT", b.String())
	return er
}
