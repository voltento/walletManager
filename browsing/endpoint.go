package browsing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/walletManager/internal/walletErrors"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, er := svc.getUsers()
		if er != nil {
			if _, ok := er.(walletErrors.HttpError); ok {
				return nil, er
			}
			return nil, walletErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}

func makeGetPaymentsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, er := svc.getPayments()
		if er != nil {
			if _, ok := er.(walletErrors.HttpError); ok {
				return nil, er
			}
			return nil, walletErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}
