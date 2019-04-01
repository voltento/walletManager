package payment

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/database"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(changeBalanceRequest)
		v, err := svc.changeBalance(req)
		if err != nil {
			return database.Error{Msg: "Field", Error: err.Error()}, nil
		}
		return v, nil
	}
}

func makeSendMoneyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(sendMoneyRequest)
		v, err := svc.sendMoney(req)
		if err != nil {
			return database.Error{Msg: "Field", Error: err.Error()}, nil
		}
		return v, nil
	}
}
