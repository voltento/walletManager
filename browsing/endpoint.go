package browsing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/httpErrors"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, er := svc.getUsers()
		if er != nil {
			return nil, httpErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}

func makeGetPaymentsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, er := svc.getPayments()
		if er != nil {
			return nil, httpErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}
