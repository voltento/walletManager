package accmamaging

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/wallet_manager/internal/httpmodel"
	"github.com/voltento/wallet_manager/internal/utils"
)
import "context"

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(httpModel.Account)
		v, er := svc.CreateUser(req.Id, req.Currency, req.Amount)
		if er != nil {
			if _, ok := er.(utils.HttpError); ok {
				return nil, er
			}
			return nil, utils.BuildDecodeError(er.Error())
		}
		return httpModel.GeneralResponse{Response: v}, nil
	}
}
