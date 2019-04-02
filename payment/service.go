package payment

import (
	"errors"
	"fmt"
	"github.com/voltento/walletManager/database"
)

type Service interface {
	changeBalance(r changeBalanceRequest) (*changeBalanceResponse, error)
	sendMoney(r sendMoneyRequest) (*sendMoneyResponse, error)
}

func CreateService(c database.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c database.WalletMgrCluster
}

func (s serviceImplementation) changeBalance(r changeBalanceRequest) (*changeBalanceResponse, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()
	er := m.IncAccountBalance(r.Id, r.Amount)
	if er != nil {
		return nil, er
	}

	return &changeBalanceResponse{Response: "Succeed"}, nil
}

func (s serviceImplementation) sendMoney(r sendMoneyRequest) (*sendMoneyResponse, error) {
	if r.FromAccId == "" {
		er := errors.New(fmt.Sprintf("From account param isn't provided"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	if r.ToAccId == "" {
		er := errors.New(fmt.Sprintf("To account param isn't provided"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	if r.FromAccId == r.ToAccId {
		er := errors.New(fmt.Sprintf("Can't transfer from the same account"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	m, closer := s.c.GetWalletMgr()
	defer closer()

	tr, er := m.StartTransaction()
	if er != nil {
		return nil, er
	}

	var fromAcc *Account
	fromAcc, er = m.GetAccount(r.FromAccId)
	if er != nil {
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	if fromAcc.Amount < r.Amount {
		er = errors.New(fmt.Sprintf("Not enough money for transfering"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	var toAcc *Account
	toAcc, er = m.GetAccount(r.ToAccId)
	if er != nil {
		return nil, er
	}

	if fromAcc.Currency != toAcc.Currency {
		er = errors.New(fmt.Sprintf("Can't transfer between account with different currency"))
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	fromAcc.Amount -= r.Amount
	er = m.UpdateAccount(fromAcc.Id, fromAcc)
	if er != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	toAcc.Amount += r.Amount
	er = m.UpdateAccount(toAcc.Id, toAcc)
	if er != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	er = m.AddPayment(&database.Payment{From_account: r.FromAccId, To_account: r.ToAccId, Amount: r.Amount})
	if er != nil {
		tr.Rollback()
		return &sendMoneyResponse{Err: er.Error()}, er
	}

	tr.Commit()
	return &sendMoneyResponse{Response: "Success"}, nil
}
