package accmamaging

import (
	"fmt"
	"github.com/voltento/wallet_manager/internal/database/ctrl"
	"github.com/voltento/wallet_manager/internal/database/model"
	"github.com/voltento/wallet_manager/internal/utils"
)

type Service interface {
	// Create user account
	CreateUser(id string, currency string, balance float64) (string, error)
}

func CreateService(c ctrl.WalletMgrCluster) Service {
	return serviceImplementation{c}
}

type serviceImplementation struct {
	c ctrl.WalletMgrCluster
}

// Implementation of Service interface
func (s serviceImplementation) CreateUser(id string, currency string, balance float64) (string, error) {
	m, closer := s.c.GetWalletMgr()
	defer closer()
	if id == "" {
		return "", utils.BuildEmptyFieldError("id")
	}

	if currency == "" {
		return "", utils.BuildEmptyFieldError("currency")
	}

	if balance < 0 {
		return "", utils.BuildGeneralQueryError(fmt.Sprintf("got unexpected balue for field `%v` expected non negotive value.", balance))
	}

	er := m.AddAccount(model.Account{Id: id, Currency: currency, Amount: balance})
	if er != nil {
		return "", er
	}

	return "Success", nil
}
