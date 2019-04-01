package browsing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/database"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, er := svc.getUsers()
		if er != nil {
			return nil, er
		}
		return v, nil
	}
}

func makeGetPaymentsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, er := svc.getPayments()
		if er != nil {
			return database.Error{Msg: "Field", Error: er.Error()}, er
		}
		return v, nil
	}
}
