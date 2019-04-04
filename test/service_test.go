package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreateService(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	addUserArgs := fmt.Sprintf("{\"id\":\"test_%v\", \"currency\": \"USD\", \"amount\": 100}", rand.Intn(100000))

	type args struct {
		body string
	}
	tests := []struct {
		name    string
		args    args
		want    httpResp
		wantErr bool
	}{
		{
			name:    "Add account: ok",
			args:    args{addUserArgs},
			want:    httpResp{"{\"response\":\"Success\"}", 200},
			wantErr: false,
		},
		{
			name:    "Add account: again",
			args:    args{addUserArgs},
			want:    httpResp{"{\"error\": \"Account id already exists\"}", 400},
			wantErr: false,
		},
		{
			name:    "Add account: miss id",
			args:    args{fmt.Sprintf("{\"currency\": \"USD\", \"amount\": 100}")},
			want:    httpResp{"{\"error\": \"got empty value for mandatory field `id`\"}", 400},
			wantErr: false,
		},
		{
			name:    "Add account: miss id",
			args:    args{fmt.Sprintf("{\"id\":\"test_%v\", \"amount\": 100}", rand.Intn(100000))},
			want:    httpResp{"{\"error\": \"got empty value for mandatory field `currency`\"}", 400},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sendRequest(addAccountUrl, "PUT", tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sendRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
