package stmt_mw

import (
	"github.com/go-pg/pg"
	"github.com/voltento/wallet_manager/internal/utils"
)

func NotEmptyResp(s Decorator, violationField string) Decorator {
	return notEmptyResp{stm: s, violationField: violationField}
}

type notEmptyResp struct {
	stm            Decorator
	violationField string
}

func (s notEmptyResp) Exec(params ...interface{}) (pg.Result, error) {
	return s.stm.Exec(params...)
}

func (s notEmptyResp) Query(model interface{}, params ...interface{}) (pg.Result, error) {
	r, er := s.stm.Query(model, params...)
	if er != nil {
		if r.RowsReturned() == 0 {
			er = utils.BuildNoDataError(s.violationField)
		}
	}

	return r, er
}
