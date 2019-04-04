package database

import (
	"github.com/go-pg/pg"
)

type psqlErrorType struct {
	ind byte
	msg string
}

var (
	constraintViolation = psqlErrorType{ind: byte(82), msg: "ExecConstraints"}
	duplicateAccountId  = psqlErrorType{ind: byte(67), msg: "23505"}
)

// Check if psql returned constraint violation error
func IsConstraintViolationError(er error) bool {
	return checkPgErrorType(er, constraintViolation)
}

// Check if psql returned duplicate key error
func IsAccIdDuplicate(er error) bool {
	return checkPgErrorType(er, duplicateAccountId)
}

// General method for matching psql error type
func checkPgErrorType(er error, expected psqlErrorType) bool {
	if pgEr, ok := er.(pg.Error); ok {
		return pgEr.Field(expected.ind) == expected.msg
	}
	return false
}
