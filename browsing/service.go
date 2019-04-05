package browsing

import (
	"github.com/voltento/walletManager/internal/database/ctrl"
	"github.com/voltento/walletManager/internal/httpmodel"
)

type Payment = httpModel.Payment

type Service interface {
	// Get all users accounts
	getUsers() ([]httpModel.Account, error)

	// Get all payments
	getPayments() ([]Payment, error)
}

func CreateService(c ctrl.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c ctrl.WalletMgrCluster
}

func (s serviceImplementation) getUsers() ([]httpModel.Account, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetAllAccounts()
}

func (s serviceImplementation) getPayments() ([]Payment, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetPayments()
}
