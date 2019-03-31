package brawsing

import "github.com/go-kit/kit/endpoint"
import "context"

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func MakeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		v, err := svc.getUsers()
		if err != nil {
			return nil, nil
		}
		return v, nil
	}
}
