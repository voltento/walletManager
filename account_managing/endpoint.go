package account_managing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/database"
)
import "context"

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request)
		v, err := svc.createUser(req.Id, req.Currency, req.Amount)
		if err != nil {
			return database.Error{Msg: "Field", Error: err.Error()}, nil
		}
		return response{Response: v}, nil
	}
}
