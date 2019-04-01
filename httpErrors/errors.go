package httpErrors

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
