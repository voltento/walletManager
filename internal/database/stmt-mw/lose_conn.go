package stmt_mw

import (
	"fmt"
	"github.com/go-pg/pg"
	error_check "github.com/voltento/wallet_manager/internal/database/errorcheck"
	"os"
)

func LoseConWithDb(s Decorator) Decorator {
	return loseConnection{s}
}

type loseConnection struct {
	stm Decorator
}

func (s loseConnection) Exec(params ...interface{}) (pg.Result, error) {
	r, er := s.stm.Exec(params...)
	existOnLostConnection(er)
	return r, er
}

func (s loseConnection) Query(model interface{}, params ...interface{}) (pg.Result, error) {
	r, er := s.stm.Query(model, params...)
	existOnLostConnection(er)
	return r, er
}

func existOnLostConnection(er error) {
	if error_check.IsLoseConnection(er) {
		fmt.Print("Lost connection with the database. Exit.")
		os.Exit(1)
	}
}
