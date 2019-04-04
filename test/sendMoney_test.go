package test

import (
	"fmt"
	"github.com/voltento/walletManager/accmamaging"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func addUsers(currency string, count int) ([]accmamaging.Account, error) {
	var er error
	rand.Seed(time.Now().UTC().UnixNano())

	accs := make([]accmamaging.Account, 0, count)
	for i := 0; i < count; i += 1 {
		ac := accmamaging.Account{Id: "test_" + strconv.Itoa(rand.Intn(10000000)), Currency: currency, Amount: 10}
		if er = addAccount(ac); er != nil {
			return nil, er
		}
		accs = append(accs, ac)
	}

	return accs, nil
}

func TestSendPayment(t *testing.T) {
	var (
		er     error
		accUsd []accmamaging.Account
		accEur []accmamaging.Account
	)

	accUsd, er = addUsers("USD", 2)
	if er != nil {
		t.Error(er.Error())
	}

	accEur, er = addUsers("EUR", 1)
	if er != nil {
		t.Error(er.Error())
		print(accEur)
	}

	type args struct {
		body string
	}
	tests := []struct {
		name      string
		args      args
		want      CheckResp
		wantErr   bool
		postCheck func() error
	}{
		{
			name: "Send money: few balance",
			args: args{fmt.Sprintf("{\"from_account\":\"%v\", \"to_account\": \"%v\", \"change_amount\": %v}",
				accUsd[0].Id,
				accUsd[1].Id,
				accUsd[0].Amount+1),
			},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{
					data: fmt.Sprintf("{\"error\": \"Few balance for the operation. Account id: `%v`\"}", accUsd[0].Id),
					code: 400,
				})
			},
			wantErr: false,
		},
		{
			name: "Send money: ok",
			args: args{fmt.Sprintf("{\"from_account\":\"%v\", \"to_account\": \"%v\", \"change_amount\": %v}",
				accUsd[0].Id,
				accUsd[1].Id,
				accUsd[0].Amount),
			},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{
					data: "{\"Response\":\"Success\"}",
					code: 200,
				})
			},
			wantErr: false,
		},
		{
			name: "Send money: diff currencies",
			args: args{fmt.Sprintf("{\"from_account\":\"%v\", \"to_account\": \"%v\", \"change_amount\": %v}",
				accUsd[1].Id,
				accEur[0].Id,
				1),
			},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{
					data: "{\"error\": \"Can't transfer between account with different currency\"}",
					code: 400,
				})
			},
			wantErr: false,
		},
		{
			name: "Send money: diff currencies",
			args: args{fmt.Sprintf("{\"from_account\":\"%v\", \"to_account\": \"%v\", \"change_amount\": 0}",
				accUsd[1].Id,
				accUsd[0].Id),
			},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{
					data: "{\"error\": \"Can't send 0 amount\"}",
					code: 400,
				})
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(sendMoneyUrl, "PUT", tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
