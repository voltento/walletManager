package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/voltento/walletManager/internal/database/model"
	"strings"
	"testing"
)

func TestBrawsPayments(t *testing.T) {
	accs, er := addAccountsWithCurrency("USD", 2)
	if er != nil {
		t.Error(er.Error())
		return
	}

	sendMoneyQuery := fmt.Sprintf("{\"from_account\":\"%v\", \"to_account\": \"%v\", \"change_amount\": 1}",
		accs[1].Id,
		accs[0].Id)
	_, er = sendRequest(sendMoneyUrl, "PUT", sendMoneyQuery)
	if er != nil {
		t.Error(er.Error())
		return
	}

	tests := []struct {
		name string
		want CheckResp
	}{
		{
			name: "Browsing payments: ok",
			want: func(r httpResp) error {
				var payments []model.Payment
				er := json.NewDecoder(strings.NewReader(r.data)).Decode(&payments)
				if er != nil {
					return er
				}
				if len(payments) < 2 {
					return errors.New(fmt.Sprintf("Didn't get expected paymetns. Got: %v", payments))
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(getPaymentsUrl, "GET", "")

			if err = tt.want(got); err != nil {
				t.Error(err.Error())
			}
		})
	}
}
