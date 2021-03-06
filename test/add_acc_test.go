package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestAddAccount(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	accId := "test_" + strconv.Itoa(rand.Intn(100000))
	addUserArgs := fmt.Sprintf("{\"id\":\"%v\", \"currency\": \"USD\", \"amount\": 100}", accId)

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
			name: "Add account: ok",
			args: args{addUserArgs},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{"{\"response\":\"Success\"}", 200})
			},
			postCheck: func() error { return assertAccExists(accId) },
		},
		{
			name: "Add account: again",
			args: args{addUserArgs},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{"{\"error\": \"Uniq violation for the field 'account id'\"}", 400})
			},
		},
		{
			name: "Add account: miss id",
			args: args{fmt.Sprintf("{\"currency\": \"USD\", \"amount\": 100}")},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{"{\"error\": \"got empty value for mandatory field `id`\"}", 400})
			},
		},
		{
			name: "Add account: miss id",
			args: args{fmt.Sprintf("{\"id\":\"test_%v\", \"amount\": 100}", rand.Intn(100000))},
			want: func(r httpResp) error {
				return assertEqHttpResp(r, httpResp{"{\"error\": \"got empty value for mandatory field `currency`\"}", 400})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(addAccountUrl, "PUT", tt.args.body)
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
