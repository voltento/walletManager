package payment

import "github.com/go-kit/kit/endpoint"
import "context"

func MakeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(changeBalanceRequest)
		v, err := svc.changeBalance(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeSendMoneyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(sendMoneyRequest)
		v, err := svc.sendMoney(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
