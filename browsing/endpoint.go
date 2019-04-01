package browsing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/database"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, err := svc.getUsers()
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeGetPaymentsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, err := svc.getPayments()
		if err != nil {
			return database.Error{Msg: "Field", Error: err.Error()}, nil
		}
		return v, nil
	}
}
