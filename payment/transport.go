package payment

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/pursesManager/database"
	"net/http"
)

type Account = database.Account

func DecodeChangeBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request changeBalanceRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type changeBalanceRequest struct {
	Id     string  `json:"id"`
	Amount float64 `json:"change_amount"`
}

type changeBalanceResponse struct {
	Response string   `json:"Response"`
	Acc      *Account `json:"account,omitempty"`
	Err      string   `json:"err,omitempty"`
}

func MakeChangeBalanceHandler(s Service) http.Handler {
	return kithttp.NewServer(
		makeGetAccountsEndpoint(s),
		DecodeChangeBalanceRequest,
		EncodeResponse,
	)
}

func DecodeSendMoneyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request sendMoneyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

type sendMoneyRequest struct {
	FromAccId string  `json:"from_account"`
	ToAccId   string  `json:"to_account"`
	Amount    float64 `json:"change_amount"`
}

type sendMoneyResponse struct {
	Response string `json:"Response"`
	Err      string `json:"err,omitempty"`
}

func MakeSendMoneyHandler(s Service) http.Handler {
	return kithttp.NewServer(
		makeSendMoneyEndpoint(s),
		DecodeSendMoneyRequest,
		EncodeResponse,
	)
}
