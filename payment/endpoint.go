package payment

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/walletErrors"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(changeBalanceRequest)
		v, er := svc.changeBalance(req)
		if er != nil {
			if _, ok := er.(walletErrors.HttpError); ok {
				return nil, er
			}
			return nil, walletErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}

func makeSendMoneyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(sendMoneyRequest)
		v, er := svc.sendMoney(req)
		if er != nil {
			if _, ok := er.(walletErrors.HttpError); ok {
				return nil, er
			}
			return nil, walletErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}
