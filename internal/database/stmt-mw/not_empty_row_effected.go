package stmt_mw

import (
	"github.com/go-pg/pg"
	"github.com/voltento/walletManager/internal/utils"
)

func NotEmptyRowEffected(s Decorator, missiedField string) Decorator {
	return notEmptyRawEffected{stm: s, missiedField: missiedField}
}

type notEmptyRawEffected struct {
	stm          Decorator
	missiedField string
}

func (s notEmptyRawEffected) Exec(params ...interface{}) (pg.Result, error) {
	return s.stm.Exec(params...)
}

func (s notEmptyRawEffected) Query(model interface{}, params ...interface{}) (pg.Result, error) {
	r, er := s.stm.Query(model, params...)
	if er != nil {
		if r.RowsAffected() == 0 {
			er = utils.BuildNoDataError(s.missiedField)
		}
	}

	return r, er
}
