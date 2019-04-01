package brawsing

import (
	"github.com/voltento/pursesManager/database"
)

type Service interface {
	getUsers() ([]*accResponse, error)
	getPayments() ([]*paymentResponse, error)
}

func CreateService(m database.WalletManager) Service {
	return serviceImplementation{m}
}

type serviceImplementation struct {
	m database.WalletManager
}

func (s serviceImplementation) getUsers() ([]*accResponse, error) {
	return s.m.GetAllAccounts()
}

func (s serviceImplementation) getPayments() ([]*paymentResponse, error) {
	return s.m.GetPayments()
}
