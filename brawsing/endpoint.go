package brawsing

import "github.com/go-kit/kit/endpoint"
import "context"

func MakeGetAccountsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, err := svc.getUsers()
		if err != nil {
			return nil, nil
		}
		return v, nil
	}
}
