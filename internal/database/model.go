package database

import "github.com/voltento/walletManager/internal/httpQueryModels"

type Account = httpQueryModels.Account

type Error struct {
	Msg   string `json:"msg"`
	Error string `json:"error"`
}
