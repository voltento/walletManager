package payment

import (
	"fmt"
	"github.com/voltento/walletManager/internal/database/ctrl"
	"github.com/voltento/walletManager/internal/database/model"
	"github.com/voltento/walletManager/internal/httpmodel"
	"github.com/voltento/walletManager/internal/utils"
)

type Service interface {
	// Change account balance for amount
	changeBalance(r httpModel.ChangeBalanceRequest) (*httpModel.GeneralResponse, error)

	// Send money from one account to another
	sendMoney(r httpModel.SendMoneyRequest) (*httpModel.GeneralResponse, error)
}

func CreateService(c ctrl.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c ctrl.WalletMgrCluster
}

func (s serviceImplementation) changeBalance(r httpModel.ChangeBalanceRequest) (*httpModel.GeneralResponse, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()
	er := m.ChangeAccountBalance(r.Id, r.Amount)
	if er != nil {
		return nil, er
	}

	return &httpModel.GeneralResponse{Response: "Success"}, nil
}

func (s serviceImplementation) transferMoney(m ctrl.WalletManager, fromId string, toId string, amount float64) error {
	var er error

	_, er = s.changeBalance(httpModel.ChangeBalanceRequest{Id: fromId, Amount: -amount})
	if er != nil {
		return er
	}

	_, er = s.changeBalance(httpModel.ChangeBalanceRequest{Id: toId, Amount: amount})
	if er != nil {
		return er
	}

	er = m.AddPayment(model.Payment{From_account: fromId, To_account: toId, Amount: amount})

	return er
}

func (s serviceImplementation) assertEqualCurrency(m ctrl.WalletManager, acc1 string, acc2 string) error {
	var er error

	var fromAcc *Account
	fromAcc, er = m.GetAccount(acc1)
	if er != nil {
		return er
	}

	var toAcc *Account
	toAcc, er = m.GetAccount(acc2)
	if er != nil {
		return er
	}

	if fromAcc.Currency != toAcc.Currency {
		return utils.BuildGeneralQueryError(fmt.Sprintf("Can't transfer between account with different currency"))
	}

	return nil
}

func (s serviceImplementation) sendMoney(r httpModel.SendMoneyRequest) (*httpModel.GeneralResponse, error) {
	var er error

	if r.Amount == 0 {
		er = utils.BuildGeneralQueryError(fmt.Sprintf("Can't send 0 amount"))
		return nil, er
	}

	if r.FromAccId == "" {
		er = utils.BuildGeneralQueryError(fmt.Sprintf("From account param can't be empty"))
		return nil, er
	}

	if r.ToAccId == "" {
		er := utils.BuildGeneralQueryError(fmt.Sprintf("To account param can't be empty"))
		return nil, er
	}

	if r.FromAccId == r.ToAccId {
		er = utils.BuildGeneralQueryError(fmt.Sprintf("Can't transfer from the same account"))
		return nil, er
	}

	m, closer := s.c.GetWalletMgr()
	defer closer()

	er = s.assertEqualCurrency(m, r.FromAccId, r.ToAccId)
	if er != nil {
		return nil, er
	}

	er = m.RunInTransaction(func() error { return s.transferMoney(m, r.FromAccId, r.ToAccId, r.Amount) })
	if er != nil {
		return nil, er
	}

	return &httpModel.GeneralResponse{Response: "Success"}, nil
}
