package test

import (
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
	er = json.NewDecoder(strings.NewReader(resp.data)).Decode(accs)
	if er != nil {
		return nil, er
	}
	return accs, nil
}
