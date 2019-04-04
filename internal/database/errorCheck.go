package database

import (
	"github.com/go-pg/pg"
)

type errorType struct {
	ind byte
	msg string
}

var (
	constraintVialation = errorType{ind: byte(82), msg: "ExecConstraints"}
	duplicateAccountId  = errorType{ind: byte(67), msg: "23505"}
)

func IsConstraintVialationError(er error) bool {
	return checkPgErrorType(er, constraintVialation)
}

func IsAccIdDuplicate(er error) bool {
	return checkPgErrorType(er, duplicateAccountId)
}

func checkPgErrorType(er error, expected errorType) bool {
	if pgEr, ok := er.(pg.Error); ok {
		return pgEr.Field(expected.ind) == expected.msg
	}
	return false
}
