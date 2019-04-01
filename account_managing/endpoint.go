package account_managing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/httpErrors"
)
import "context"

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request)
		v, er := svc.createUser(req.Id, req.Currency, req.Amount)
		if er != nil {
			return nil, httpErrors.BuildDecodeError(er.Error())
		}
		return response{Response: v}, nil
	}
}
