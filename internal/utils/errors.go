package utils

import (
	"fmt"
	"github.com/go-kit/kit/transport/http"
)

// Http error interface for sending error message and http return code
type HttpError interface {
	error
	http.StatusCoder
}

// HttpError implementation
type httpError struct {
	code int
	msg  string
}

// http.StatusCoder implementation
func (e httpError) StatusCode() int {
	return e.code
}

// Build json response from the error
func (e httpError) Error() string {
	return fmt.Sprintf("{\"error\": \"%v\"}", e.msg)
}

// General purpose error
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

func BuildNoDataError(acId string) HttpError {
	return httpError{code: 200, msg: fmt.Sprintf("No data for `%v`", acId)}
}

func BuildEmptyFieldError(fieldName string) error {
	return BuildGeneralQueryError(fmt.Sprintf("got empty value for mandatory field `%v`", fieldName))
}
