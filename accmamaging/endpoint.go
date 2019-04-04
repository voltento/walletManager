package accmamaging

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/walletManager/internal/httpQueryModels"
	"github.com/voltento/walletManager/internal/utils"
)
import "context"

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(httpQueryModels.Account)
		v, er := svc.createUser(req.Id, req.Currency, req.Amount)
		if er != nil {
			if _, ok := er.(utils.HttpError); ok {
				return nil, er
			}
			return nil, utils.BuildDecodeError(er.Error())
		}
		return httpQueryModels.GeneralResponse{Response: v}, nil
	}
}
