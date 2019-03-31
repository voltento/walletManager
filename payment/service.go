package payment

import (
	"github.com/voltento/pursesManager/database"
)

type Service interface {
	changeBalance(r changeBalance) (*changeBalanceResponse, error)
}

func CreateService(m database.WalletManager) Service {
	return serviceImplementation{m}
}

type serviceImplementation struct {
	m database.WalletManager
}

func (s serviceImplementation) changeBalance(r changeBalance) (*changeBalanceResponse, error) {
	tr, err := s.m.StartTransaction()
	if err != nil {
		return nil, err
	}

	acc, er := s.m.GetAccount(r.Id)

	if er != nil {
		return &changeBalanceResponse{Response: "Field", Err: er.Error()}, nil
	}

	newAmount := acc.Amount + r.Amount
	if newAmount < 0 {
		return &changeBalanceResponse{Response: "Not enough balance", Acc: acc}, nil
	} else {
		acc.Amount = newAmount
		er = s.m.UpdateAccount(acc.Id, acc)
		if er != nil {
			return &changeBalanceResponse{Response: "Field", Err: er.Error()}, nil
		}
		er = tr.Commit()
		if er != nil {
			return &changeBalanceResponse{Response: "Field", Err: er.Error()}, nil
		}
	}

	return &changeBalanceResponse{Response: "Success", Acc: acc}, nil
}
