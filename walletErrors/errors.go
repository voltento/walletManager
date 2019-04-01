package walletErrors

import (
	"fmt"
	"github.com/go-kit/kit/transport/http"
)

type HttpError interface {
	error
	http.StatusCoder
}

func BuildProcessingError(p string) HttpError {
	return processingError{p: p}
}

type processingError struct {
	p string
}

func (e processingError) StatusCode() int {
	return 400
}

func (e processingError) Error() string {
	return fmt.Sprintf("The error occured during processin the query. Error: `%v`", e.p)
}

func BuildDecodeError(p string) HttpError {
	return decodeError{p: p}
}

type decodeError struct {
	p string
}

func (e decodeError) StatusCode() int {
	return 400
}

func (e decodeError) Error() string {
	return fmt.Sprintf("The error occured during decode the query. Error: `%v`", e.p)
}

func BuildFindAccountError(acId string) HttpError {
	return findAccountError{p: acId}
}

type findAccountError struct {
	p string
}

func (e findAccountError) StatusCode() int {
	return 400
}

func (e findAccountError) Error() string {
	return fmt.Sprintf("Can't find record an account with id `%v`", e.p)
}
