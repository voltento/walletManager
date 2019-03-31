package payment

import "github.com/go-kit/kit/endpoint"
import "context"

func MakeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(changeBalance)
		v, err := svc.changeBalance(req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
