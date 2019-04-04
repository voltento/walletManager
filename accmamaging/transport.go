package accmamaging

import (
	"context"
	"encoding/json"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/voltento/walletManager/internal/httpQueryModels"
	"github.com/voltento/walletManager/internal/utils"
	"net/http"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request httpQueryModels.Account
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

func MakeHandler(bs Service) http.Handler {
	return kithttp.NewServer(
		makeAddEndpoint(bs),
		DecodeRequest,
		EncodeResponse,
	)
}
