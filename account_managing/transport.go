package account_managing

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/pursesManager/database"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type request = database.Account

type response struct {
	Response string `json:"response"`
	Err      string `json:"err,omitempty"`
}

func MakeHandler(bs Service) http.Handler {
	return kithttp.NewServer(
		MakeAddEndpoint(bs),
		DecodeRequest,
		EncodeResponse,
	)
}
