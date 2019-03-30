package account_managing

import "github.com/go-kit/kit/endpoint"
import "context"

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func MakeEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request)
		v, err := svc.createUser(req.Id, req.Currency, req.Amount)
		if err != nil {
			return response{v, err.Error()}, nil
		}
		return response{v, ""}, nil
	}
}
