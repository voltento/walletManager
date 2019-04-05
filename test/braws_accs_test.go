package test

import (
	"errors"
	"fmt"
	"github.com/voltento/walletManager/internal/httpmodel"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestBrawsAccount(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	ac := httpModel.Account{
		Id:       "test_" + strconv.Itoa(rand.Intn(100000)),
		Currency: "USD",
		Amount:   10,
	}

	er := addAccount(ac)
	if er != nil {
		t.Error(er.Error())
		return
	}

	tests := []struct {
		name string
		want CheckResp
	}{
		{
			name: "Browsing account: ok",
			want: func(r httpResp) error {
				expectedCode := 200
				if r.code != expectedCode {
					return errors.New(fmt.Sprintf("Got the unexpected return code. Expected: `%v` got: `%v`", expectedCode, r.code))
				}

				gotAcc, er := getAccount(ac.Id)
				if er != nil {
					return er
				}

				if !reflect.DeepEqual(ac, gotAcc) {
					return errors.New(fmt.Sprintf("Got the unexpected account. Expected: `%v` got: `%v`", ac, gotAcc))
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(getAccountsUrl, "GET", "")

			if err = tt.want(got); err != nil {
				t.Error(err.Error())
			}
		})
	}
}
