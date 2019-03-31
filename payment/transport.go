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
	var request changeBalance
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type changeBalance struct {
	Id     string  `json:"id"`
	Amount float64 `json:"change_amount"`
}

type changeBalanceResponse struct {
	Response string   `json:"changeBalanceResponse"`
	Acc      *Account `json:"account,omitempty"`
	Err      string   `json:"err,omitempty"`
}

func MakeChangeBalanceHandler(s Service) http.Handler {
	return kithttp.NewServer(
		MakeGetAccountsEndpoint(s),
		DecodeChangeBalanceRequest,
		EncodeResponse,
	)
}
