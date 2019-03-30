package account_managing

import (
	"errors"
	"fmt"
)

type Service interface {
	createUser(id string, currency string, balance float64) (string, error)
}

func CreateService() Service {
	return serviceImplementation{}
}

type serviceImplementation struct{}

func (s serviceImplementation) createUser(id string, currency string, balance float64) (string, error) {
	if id == "" {
		return "", buildEmptyFieldError("id")
	}

	if currency == "" {
		return "", buildEmptyFieldError("currency")
	}

	if balance < 0 {
		return "", errors.New(fmt.Sprintf("got unexpected balue for field `%v` expected non negotive value.", balance))
	}

	// add logic here
	return "Success", nil
}

func buildEmptyFieldError(fieldName string) error {
	return errors.New(fmt.Sprintf("got empty value for mandatory field `%v`", fieldName))
}
