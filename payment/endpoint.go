package payment

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/wallet_manager/internal/httpmodel"
	"github.com/voltento/wallet_manager/internal/utils"
)
import "context"

func makeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(httpModel.ChangeBalanceRequest)
		v, er := svc.ChangeBalance(req)
		if er != nil {
			if _, ok := er.(utils.HttpError); ok {
				return nil, er
			}
			return nil, utils.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}

func makeSendMoneyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(httpModel.SendMoneyRequest)
		v, er := svc.SendMoney(req)
		if er != nil {
			if _, ok := er.(utils.HttpError); ok {
				return nil, er
			}
			return nil, utils.BuildDecodeError(er.Error())
		}
		return v, nil
	}
}
