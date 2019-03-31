package payment

import (
	"errors"
	"fmt"
	"github.com/voltento/pursesManager/database"
)

type Service interface {
	changeBalance(r changeBalanceRequest) (*changeBalanceResponse, error)
	sendMoney(r sendMoneyRequest) (*sendMoneyResponse, error)
}

func CreateService(m database.WalletManager) Service {
	return serviceImplementation{m}
}

type serviceImplementation struct {
	m database.WalletManager
}

func (s serviceImplementation) changeBalance(r changeBalanceRequest) (*changeBalanceResponse, error) {
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

func (s serviceImplementation) sendMoney(r sendMoneyRequest) (*sendMoneyResponse, error) {
	if r.FromAccId == "" {
		err := errors.New(fmt.Sprintf("From account param isn't provided"))
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	if r.ToAccId == "" {
		err := errors.New(fmt.Sprintf("To account param isn't provided"))
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	if r.FromAccId == r.ToAccId {
		err := errors.New(fmt.Sprintf("Can't transfer from the same account"))
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	tr, err := s.m.StartTransaction()
	if err != nil {
		return nil, err
	}

	var fromAcc *Account
	fromAcc, err = s.m.GetAccount(r.FromAccId)
	if err != nil {
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	if fromAcc.Amount < r.Amount {
		err = errors.New(fmt.Sprintf("Not enough money for transfering"))
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	var toAcc *Account
	toAcc, err = s.m.GetAccount(r.ToAccId)
	if err != nil {
		return nil, err
	}

	if fromAcc.Currency != toAcc.Currency {
		err = errors.New(fmt.Sprintf("Can't transfer between account with different currency"))
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	fromAcc.Amount -= r.Amount
	err = s.m.UpdateAccount(fromAcc.Id, fromAcc)
	if err != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	toAcc.Amount += r.Amount
	err = s.m.UpdateAccount(toAcc.Id, toAcc)
	if err != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: err.Error()}, err
	}

	tr.Commit()

	return &sendMoneyResponse{Response: "Success"}, nil
}
