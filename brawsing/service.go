package brawsing

import (
	"github.com/voltento/pursesManager/database"
)

type Service interface {
	getUsers() ([]*response, error)
}

func CreateService(m database.WalletManager) Service {
	return serviceImplementation{m}
}

type serviceImplementation struct {
	m database.WalletManager
}

func (s serviceImplementation) getUsers() ([]*response, error) {
	return s.m.GetAllAccounts()
}
