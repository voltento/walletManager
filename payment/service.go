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

	tr, er := s.m.StartTransaction()
	if er != nil {
		return nil, er
	}

	var fromAcc *Account
	fromAcc, er = s.m.GetAccount(r.FromAccId)
	if er != nil {
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	if fromAcc.Amount < r.Amount {
		er = errors.New(fmt.Sprintf("Not enough money for transfering"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	var toAcc *Account
	toAcc, er = s.m.GetAccount(r.ToAccId)
	if er != nil {
		return nil, er
	}

	if fromAcc.Currency != toAcc.Currency {
		er = errors.New(fmt.Sprintf("Can't transfer between account with different currency"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	fromAcc.Amount -= r.Amount
	er = s.m.UpdateAccount(fromAcc.Id, fromAcc)
	if er != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	toAcc.Amount += r.Amount
	er = s.m.UpdateAccount(toAcc.Id, toAcc)
	if er != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	er = s.m.AddPayment(&database.Payment{From_account: r.FromAccId, To_account: r.ToAccId, Amount: r.Amount})
	if er != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	tr.Commit()
	return &sendMoneyResponse{Response: "Success"}, nil
}
