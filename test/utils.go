package test

import (
	"errors"
	"fmt"
	"reflect"
)

type CheckResp = func(resp httpResp) error

func assertEqHttpResp(resp httpResp, expected httpResp) error {
	if !reflect.DeepEqual(resp, expected) {
		return errors.New(fmt.Sprintf("Got unexpected message. Expected: '%v', got '%v'", expected, resp))
	}
	return nil
}

const (
	walltMgrAddr  = "http://:8080"
	addAccountUrl = walltMgrAddr + "/account_managing/add/"
)
