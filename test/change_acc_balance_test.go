package test

import (
	"errors"
	"fmt"
	"github.com/voltento/walletManager/internal/httpmodel"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestChangeAccBalance(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	acc := httpModel.Account{Id: "test_" + strconv.Itoa(rand.Intn(100000)), Currency: "USD", Amount: 10}
	er := addAccount(acc)
	if er != nil {
		t.Error(er.Error())
		return
	}

	type args struct {
		body string
	}
	tests := []struct {
		name      string
		args      args
		want      CheckResp
		postCheck func() error
	}{
		{
			name: "Change balance positive: ok",
			args: args{fmt.Sprintf("{\"id\":\"%v\", \"change_amount\": %v}", acc.Id, 1)},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{"{\"response\":\"Success\"}", 200})
			},
			postCheck: func() error {
				newAcc, er := getAccount(acc.Id)
				if er != nil {
					return er
				}
				expAmount := acc.Amount + 1
				if newAcc.Amount != expAmount {
					return errors.New(fmt.Sprintf("Get unexpected account balance. Expected: '%v' got: '%v'", expAmount, newAcc.Amount))
				}
				return nil
			},
		},
		{
			name: "Change balance positive: ok",
			args: args{fmt.Sprintf("{\"id\":\"%v\", \"change_amount\": %v}", acc.Id, -1)},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{"{\"response\":\"Success\"}", 200})
			},
			postCheck: func() error {
				newAcc, er := getAccount(acc.Id)
				if er != nil {
					return er
				}
				expAmount := acc.Amount
				if newAcc.Amount != expAmount {
					return errors.New(fmt.Sprintf("Get unexpected account balance. Expected: '%v' got: '%v'", expAmount, newAcc.Amount))
				}
				return nil
			},
		},
		{
			name: "Change balance positive: ok",
			args: args{fmt.Sprintf("{\"id\":\"%v\", \"change_amount\": %v}", acc.Id, -100)},
			want: func(r httpResp) error {
				return assertEqHttpResp(r,
					httpResp{
						data: fmt.Sprintf("{\"error\": \"Few balance for the operation. Account id: `%v`\"}", acc.Id),
						code: 400,
					})
			},
			postCheck: func() error {
				newAcc, er := getAccount(acc.Id)
				if er != nil {
					return er
				}
				expAmount := acc.Amount
				if newAcc.Amount != expAmount {
					return errors.New(fmt.Sprintf("Get unexpected account balance. Expected: '%v' got: '%v'", expAmount, newAcc.Amount))
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(changeBalanceUrl, "PUT", tt.args.body)
			if err = tt.want(got); err != nil {
				t.Error(err.Error())
			}
			if tt.postCheck != nil {
				err = tt.postCheck()
				if err != nil {
					t.Error(err.Error())
				}
			}
		})
	}
}
