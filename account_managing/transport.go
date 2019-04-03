package account_managing

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/walletManager/internal/database"
	"github.com/voltento/walletManager/internal/walletErrors"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request
	if er := json.NewDecoder(r.Body).Decode(&request); er != nil {
		return nil, er
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	er := json.NewEncoder(w).Encode(response)
	if er != nil {
		er = walletErrors.BuildProcessingError(er.Error())
	}
	return er
}

type request = database.Account

type response struct {
	Response string `json:"response"`
	Err      string `json:"err,omitempty"`
}

func MakeHandler(bs Service) http.Handler {
	return kithttp.NewServer(
		makeAddEndpoint(bs),
		DecodeRequest,
		EncodeResponse,
	)
}
