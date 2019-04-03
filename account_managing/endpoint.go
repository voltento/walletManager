package account_managing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/walletManager/internal/walletErrors"
)
import "context"

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request)
		v, er := svc.createUser(req.Id, req.Currency, req.Amount)
		if er != nil {
			if _, ok := er.(walletErrors.HttpError); ok {
				return nil, er
			}
			return nil, walletErrors.BuildDecodeError(er.Error())
		}
		return response{Response: v}, nil
	}
}
