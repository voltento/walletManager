package browsing

import (
	"github.com/voltento/wallet_manager/internal/database/ctrl"
	"github.com/voltento/wallet_manager/internal/httpmodel"
)

type Payment = httpModel.Payment

type Service interface {
	// Get all users accounts
	GetUsers() ([]httpModel.Account, error)

	// Get all payments
	GetPayments() ([]Payment, error)
}

func CreateService(c ctrl.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c ctrl.WalletMgrCluster
}

// Implementation of Service interface
func (s serviceImplementation) GetUsers() ([]httpModel.Account, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetAllAccounts()
}

// Implementation of Service interface
func (s serviceImplementation) GetPayments() ([]Payment, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetPayments()
}
