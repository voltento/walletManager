package errorcheck

import (
	"github.com/go-pg/pg"
)

type psqlErrorType struct {
	ind byte
	msg string
}

var (
	constraintViolation = psqlErrorType{ind: byte(82), msg: "ExecConstraints"}
	uniqVialation       = psqlErrorType{ind: byte(67), msg: "23505"}
)

// Check if psql returned constraint violation error
func IsConstraintViolationError(er error) bool {
	return checkPgErrorType(er, constraintViolation)
}

// Check if psql returned duplicate key error
func IsUniqVialation(er error) bool {
	return checkPgErrorType(er, uniqVialation)
}

// General method for matching psql error type
func checkPgErrorType(er error, expected psqlErrorType) bool {
	if er == nil {
		return false
	}

	if pgEr, ok := er.(pg.Error); ok {
		return pgEr.Field(expected.ind) == expected.msg
	}
	return false
}

// Check if service lost connection with database
func IsLoseConnection(er error) bool {
	return er != nil && er.Error() == "EOF"
}
