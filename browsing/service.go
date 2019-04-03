package browsing

import (
	"github.com/voltento/walletManager/internal/database"
)

type Service interface {
	getUsers() ([]accResponse, error)
	getPayments() ([]paymentResponse, error)
}

func CreateService(c database.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c database.WalletMgrCluster
}

func (s serviceImplementation) getUsers() ([]accResponse, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetAllAccounts()
}

func (s serviceImplementation) getPayments() ([]paymentResponse, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()

	return m.GetPayments()
}
