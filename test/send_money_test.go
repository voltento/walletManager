package test

import (
	"fmt"
	"github.com/voltento/walletManager/internal/httpmodel"
	"testing"
)

func TestSendPayment(t *testing.T) {
	var (
		er     error
		accUsd []httpModel.Account
		accEur []httpModel.Account
	)

	accUsd, er = addAccountsWithCurrency("USD", 2)
	if er != nil {
		t.Error(er.Error())
		return
	}

	accEur, er = addAccountsWithCurrency("EUR", 1)
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
					data: "{\"response\":\"Success\"}",
					code: 200,
				})
			},
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(sendMoneyUrl, "PUT", tt.args.body)
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
