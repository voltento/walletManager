package service

import (
	"errors"
	"strings"
)

// StringService provides operations on strings.
type Service interface {
	Uppercase(string) (string, error)
	Count(string) int
}

func CreateService() Service {
	return serviceImplementation{}
}

// stringService is a concrete implementation of StringService
type serviceImplementation struct{}

func (serviceImplementation) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (serviceImplementation) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
