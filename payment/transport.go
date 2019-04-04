package payment

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/walletManager/internal/httpQueryModels"
	"github.com/voltento/walletManager/internal/utils"
	"net/http"
)

type Account = httpQueryModels.Account

type ChangeBalanceRequest struct {
	Id     string  `json:"id"`
	Amount float64 `json:"change_amount"`
}

type SendMoneyRequest struct {
	FromAccId string  `json:"from_account"`
	ToAccId   string  `json:"to_account"`
	Amount    float64 `json:"change_amount"`
}

func DecodeChangeBalanceRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ChangeBalanceRequest
	if er := json.NewDecoder(r.Body).Decode(&request); er != nil {
		return nil, er
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	er := json.NewEncoder(w).Encode(response)
	if er != nil {
		er = utils.BuildProcessingError(er.Error())
	}
	return er
}

func MakeChangeBalanceHandler(s Service) http.Handler {
	return kithttp.NewServer(
		makeGetAccountsEndpoint(s),
		DecodeChangeBalanceRequest,
		EncodeResponse,
	)
}

func DecodeSendMoneyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SendMoneyRequest
	if er := json.NewDecoder(r.Body).Decode(&request); er != nil {
		return nil, er
	}
	return request, nil
}

func MakeSendMoneyHandler(s Service) http.Handler {
	return kithttp.NewServer(
		makeSendMoneyEndpoint(s),
		DecodeSendMoneyRequest,
		EncodeResponse,
	)
}
