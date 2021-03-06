package stmt_mw

import (
	"fmt"
	"github.com/go-pg/pg"
	error_check "github.com/voltento/wallet_manager/internal/database/errorcheck"
	"github.com/voltento/wallet_manager/internal/utils"
)

func UniqViolation(s Decorator, violationField string) Decorator {
	return uniqViolation{stm: s, violationField: violationField}
}

type uniqViolation struct {
	stm            Decorator
	violationField string
}

func (s uniqViolation) Exec(params ...interface{}) (pg.Result, error) {
	r, er := s.stm.Exec(params...)
	er = s.handleError(er)
	return r, er
}

func (s uniqViolation) Query(model interface{}, params ...interface{}) (pg.Result, error) {
	r, er := s.stm.Query(model, params...)
	er = s.handleError(er)
	return r, er
}

func (s uniqViolation) handleError(er error) error {
	if error_check.IsUniqVialation(er) {
		er = utils.BuildGeneralQueryError(fmt.Sprintf("Uniq violation for the field '%v'", s.violationField))
	}
	return er
}
