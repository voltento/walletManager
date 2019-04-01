package payment

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/httpErrors"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(changeBalanceRequest)
		v, er := svc.changeBalance(req)
		if er != nil {
			return nil, httpErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}

func makeSendMoneyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(sendMoneyRequest)
		v, er := svc.sendMoney(req)
		if er != nil {
			return nil, httpErrors.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}
