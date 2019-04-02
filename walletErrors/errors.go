package walletErrors

import (
	"fmt"
	"github.com/go-kit/kit/transport/http"
)

type HttpError interface {
	error
	http.StatusCoder
}

type httpError struct {
	code int
	msg  string
}

func (e httpError) StatusCode() int {
	return e.code
}

func (e httpError) Error() string {
	return e.msg
}

func BuildGeneralQueryError(msg string) HttpError {
	return httpError{code: 400, msg: msg}
}

func BuildProcessingError(er string) HttpError {
	return httpError{code: 400, msg: fmt.Sprintf("The error occured during processin the query. Error: `%v`", er)}
}

func BuildDecodeError(er string) HttpError {
	return httpError{code: 400, msg: fmt.Sprintf("The error occured during processin the query. Error: `%v`", er)}
}

func BuildFindAccountError(acId string) HttpError {
	return httpError{code: 400, msg: fmt.Sprintf("Can't find an account with id `%v`", acId)}
}

func BuildFewBalanceError(acId string) HttpError {
	return httpError{code: 400, msg: fmt.Sprintf("Few balance for the operation. Account id: `%v`", acId)}
}
