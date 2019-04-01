package account_managing

import (
	"errors"
	"fmt"
	"github.com/voltento/walletManager/database"
)

type Service interface {
	createUser(id string, currency string, balance float64) (string, error)
}

func CreateService(c database.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c database.WalletMgrCluster
}

func (s serviceImplementation) createUser(id string, currency string, balance float64) (string, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()
	if id == "" {
		return "", buildEmptyFieldError("id")
	}

	if currency == "" {
		return "", buildEmptyFieldError("currency")
	}

	if balance < 0 {
		return "", errors.New(fmt.Sprintf("got unexpected balue for field `%v` expected non negotive value.", balance))
	}

	er := m.AddAccount(&database.Account{Id: id, Currency: currency, Amount: balance})
	if er != nil {
		return "", er
	}

	return "Success", nil
}

func buildEmptyFieldError(fieldName string) error {
	return errors.New(fmt.Sprintf("got empty value for mandatory field `%v`", fieldName))
}
