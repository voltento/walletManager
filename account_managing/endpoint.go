package account_managing

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/voltento/pursesManager/database"
)
import "context"

func makeAddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request)
		v, er := svc.createUser(req.Id, req.Currency, req.Amount)
		if er != nil {
			return database.Error{Msg: "Field", Error: er.Error()}, er
		}
		return response{Response: v}, nil
	}
}
