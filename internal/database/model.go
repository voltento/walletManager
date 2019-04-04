package database

import "github.com/voltento/walletManager/internal/httpQueryModels"

type Account = httpQueryModels.Account

type Payment struct {
	Id           int
	From_account string  `json:"from_account_id"`
	Amount       float64 `json:"amount"`
	To_account   string  `json:"to_account_id"`
}

type Error struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
