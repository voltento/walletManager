package stmt_middleware

import "github.com/go-pg/pg"

type Decorator interface {
	Exec(params ...interface{}) (pg.Result, error)
	Query(model interface{}, params ...interface{}) (pg.Result, error)
}
