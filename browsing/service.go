package browsing

import (
	"github.com/voltento/walletManager/internal/database"
	"github.com/voltento/walletManager/internal/httpQueryModels"
)

type Service interface {
	// Get all users accounts
	getUsers() ([]httpQueryModels.Account, error)

	// Get all payments
	getPayments() ([]paymentResponse, error)
}

func CreateService(c database.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c database.WalletMgrCluster
}

func (s serviceImplementation) getUsers() ([]httpQueryModels.Account, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetAllAccounts()
}

func (s serviceImplementation) getPayments() ([]paymentResponse, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetPayments()
}
