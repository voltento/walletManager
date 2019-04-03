package account_managing

import (
	"fmt"
	"github.com/voltento/walletManager/database"
	"github.com/voltento/walletManager/walletErrors"
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
		return "", walletErrors.BuildGeneralQueryError(fmt.Sprintf("got unexpected balue for field `%v` expected non negotive value.", balance))
	}

	er := m.AddAccount(&database.Account{Id: id, Currency: currency, Amount: balance})
	if er != nil {
		return "", er
	}

	return "Success", nil
}

func buildEmptyFieldError(fieldName string) error {
	return walletErrors.BuildGeneralQueryError(fmt.Sprintf("got empty value for mandatory field `%v`", fieldName))
}
