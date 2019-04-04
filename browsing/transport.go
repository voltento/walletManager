package browsing

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/walletManager/internal/utils"
	"net/http"
)

func DecodeRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	er := json.NewEncoder(w).Encode(response)
	if er != nil {
		er = utils.BuildProcessingError(er.Error())
	}
	return er
}

func MakeGetAccountsHandler(s Service) http.Handler {
	return kithttp.NewServer(
		makeGetAccountsEndpoint(s),
		DecodeRequest,
		EncodeResponse,
	)
}

func MakeGetPaymentsHandler(s Service) http.Handler {
	return kithttp.NewServer(
		makeGetPaymentsEndpoint(s),
		DecodeRequest,
		EncodeResponse,
	)
}
